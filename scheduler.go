package zurvan

import (
	"context"
	"time"
)

// Stage represents a phase in the execution of systems.
// Systems can be added to different stages to control their execution order.
type Stage = uint8

const (
	// Runs once at the start of the world
	StartupStage Stage = iota

	// Runs every frame, before FixedUpdateStage
	PreUpdateStage

	// Runs at a fixed tick rate, independent of frame rate
	FixedUpdateStage

	// Runs every frame, after FixedUpdateStage
	UpdateStage

	// Runs once at the end of the world
	EndStage
)

// System represents a unit of logic that operates on the world.
// Systems are executed in stages, and can be added to multiple stages if needed.
type System interface {
	Update(w *World, dt time.Duration)
}

// Helper function to build a stage with its associated systems.
func BuildStageSystems(stage Stage, systems ...System) (Stage, []System) {
	return stage, systems
}

type scheduler struct {
	systems  map[Stage][]System
	commands *commands
	events   *events

	tickRate    time.Duration
	accumulator time.Duration
	tick        uint64
}

func newScheduler(commands *commands, events *events, tickRate time.Duration) *scheduler {
	return &scheduler{
		systems:     make(map[Stage][]System),
		commands:    commands,
		events:      events,
		tickRate:    tickRate,
		accumulator: 0,
		tick:        0,
	}
}

func (s *scheduler) stage(stage Stage) []System {
	systems, ok := s.systems[stage]
	if !ok {
		systems = []System{}
		s.systems[stage] = systems
	}

	return systems
}

func (s *scheduler) addSystem(stage Stage, system System) {
	systems := s.stage(stage)
	systems = append(systems, system)
	s.systems[stage] = systems
}

func (s *scheduler) runStage(ctx context.Context, world *World, stage Stage, dt time.Duration) bool {
	systems := s.stage(stage)

	for _, system := range systems {
		if ctx.Err() != nil {
			return false
		}

		system.Update(world, dt)

		s.commands.apply(world)
		s.events.apply()
	}

	return ctx.Err() == nil
}

func (s *scheduler) run(ctx context.Context, world *World) {
	defer func() {
		s.commands.apply(world)
		s.events.apply()

		s.runStage(context.Background(), world, EndStage, 0)

		s.commands.apply(world)
		s.events.apply()

		s.events.clear()
	}()

	s.commands.apply(world)
	s.events.apply()
	if !s.runStage(ctx, world, StartupStage, 0) {
		return
	}

	last := time.Now()

	timer := time.NewTimer(s.tickRate)
	defer timer.Stop()

	for ctx.Err() == nil {
		now := time.Now()
		frameTime := now.Sub(last)
		last = now

		s.accumulator += frameTime

		if !s.runStage(ctx, world, PreUpdateStage, frameTime) {
			return
		}

		for s.accumulator >= s.tickRate {
			if !s.runStage(ctx, world, FixedUpdateStage, s.tickRate) {
				return
			}
			s.accumulator -= s.tickRate
			s.tick += 1
		}

		if !s.runStage(ctx, world, UpdateStage, frameTime) {
			return
		}

		s.events.clear()

		sleepTime := s.tickRate - time.Since(now)
		if sleepTime > 0 {
			timer.Reset(sleepTime)

			select {
			case <-ctx.Done():
				return
			case <-timer.C:
			}
		}
	}
}
