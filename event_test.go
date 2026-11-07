package zurvan

import (
	"context"
	"testing"
	"time"
)

type eventTestSystem func(*World, time.Duration)

func (f eventTestSystem) Update(w *World, dt time.Duration) { f(w, dt) }

type frameTestEvent struct{ value int }

func TestExternalEventSurvivesFrameEnd(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	w := NewWorld(time.Millisecond)
	frame, firstSeen, secondSeen := 0, 0, 0
	w.AddSystems(UpdateStage, []System{
		eventTestSystem(func(w *World, _ time.Duration) {
			frame++
			events := OnEvent[frameTestEvent](w)
			if frame == 2 {
				if len(events) != 1 || events[0].value != 42 {
					t.Error("first subscriber did not receive the external event next frame")
				}
			} else if len(events) != 0 {
				t.Error("event appeared outside its delivery frame")
			}
			firstSeen += len(events)
			if frame == 1 {
				// An external producer emits after this subscriber has read events.
				done := make(chan struct{})
				go func() {
					w.EmitEvents(frameTestEvent{42})
					close(done)
				}()
				<-done
			}
			if frame == 3 {
				cancel()
			}
		}),
		eventTestSystem(func(w *World, _ time.Duration) {
			events := OnEvent[frameTestEvent](w)
			if frame == 1 && len(events) != 0 {
				t.Error("later subscriber saw an event published midway through the frame")
			}
			secondSeen += len(events)
		}),
	})
	w.Run(ctx)
	if frame != 3 || firstSeen != 1 || secondSeen != 1 {
		t.Errorf("frames=%d, subscriber deliveries=%d/%d; want 3, 1/1", frame, firstSeen, secondSeen)
	}
}

func TestSystemEventsUseNextFrameSnapshot(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	w := NewWorld(time.Millisecond)
	frame, fixedSteps := 0, 0
	w.AddSystems(PreUpdateStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		frame++
		if frame == 1 {
			// Run multiple fixed steps against the same event snapshot.
			w.scheduler.accumulator = 2 * time.Millisecond
			w.EmitEvents(frameTestEvent{1})
		} else if frame == 2 {
			w.scheduler.accumulator = 2 * time.Millisecond
		}
	})})
	checkSnapshot := func(w *World) {
		t.Helper()
		want := 0
		if frame == 2 {
			want = 1
		}
		if events := OnEvent[frameTestEvent](w); len(events) != want {
			t.Errorf("frame %d: got %d events, want %d", frame, len(events), want)
		}
	}
	w.AddSystems(FixedUpdateStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		if frame <= 2 {
			fixedSteps++
		}
		checkSnapshot(w)
	})})
	w.AddSystems(UpdateStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		checkSnapshot(w)
		if frame == 3 {
			cancel()
		}
	})})
	w.Run(ctx)
	if frame != 3 || fixedSteps != 4 {
		t.Errorf("frames=%d, fixed steps=%d; want 3, 4", frame, fixedSteps)
	}
}

func TestStartupEventsReachFirstFrame(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	w := NewWorld(time.Millisecond)
	w.EmitEvents(frameTestEvent{1})
	w.AddSystems(StartupStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		if events := OnEvent[frameTestEvent](w); len(events) != 1 || events[0].value != 1 {
			t.Error("pre-run event was not readable during startup")
		}
		w.EmitEvents(frameTestEvent{2})
	})})
	frames := 0
	w.AddSystems(UpdateStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		frames++
		events := OnEvent[frameTestEvent](w)
		if len(events) != 2 || events[0].value != 1 || events[1].value != 2 {
			t.Error("first frame did not receive pre-run and startup events in order")
		}
		cancel()
	})})
	w.Run(ctx)
	if frames != 1 {
		t.Errorf("ran %d frames, want 1", frames)
	}
}

func TestShutdownPublishesPendingEvents(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	w := NewWorld(time.Millisecond)
	w.AddSystems(UpdateStage, []System{eventTestSystem(func(w *World, _ time.Duration) {
		w.EmitEvents(frameTestEvent{1})
		cancel()
	})})
	endCalls := 0
	w.AddSystems(EndStage, []System{
		eventTestSystem(func(w *World, _ time.Duration) {
			endCalls++
			if events := OnEvent[frameTestEvent](w); len(events) != 1 || events[0].value != 1 {
				t.Error("pending event was not published before shutdown")
			}
			w.EmitEvents(frameTestEvent{2})
		}),
		eventTestSystem(func(w *World, _ time.Duration) {
			endCalls++
			if events := OnEvent[frameTestEvent](w); len(events) != 2 || events[1].value != 2 {
				t.Error("shutdown event was not visible to the next shutdown system")
			}
		}),
	})
	w.Run(ctx)
	if endCalls != 2 {
		t.Errorf("ran %d shutdown systems, want 2", endCalls)
	}
	if len(OnEvent[frameTestEvent](w)) != 0 || len(w.events.events.drain()) != 0 {
		t.Error("events remain after shutdown")
	}
}
