package zurvan

import (
	"time"
)

type Stage = uint8

const (
	StartupStage Stage = iota
	PreUpdateStage
	FixedUpdateStage
	UpdateStage
	PostUpdateStage
	EndStage
)

type System interface {
	Update(w *World, dt time.Duration)
}

func BuildStageSystems(stage Stage, systems ...System) (Stage, []System) {
	return stage, systems
}

type scheduler struct {
	isRunning bool

	systems  map[Stage][]System
	commands *commands
	events   *events

	tickRate    time.Duration
	accumulator time.Duration
}

func NewScheduler(commands *commands, events *events, tickRate time.Duration) *scheduler {
	return &scheduler{
		isRunning:   true,
		systems:     make(map[Stage][]System),
		commands:    commands,
		events:      events,
		tickRate:    tickRate,
		accumulator: 0,
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

func (s *scheduler) runStage(world *World, stage Stage, dt time.Duration) {
	systems := s.stage(stage)

	for _, system := range systems {
		system.Update(world, dt)
	}

	s.commands.apply(world)
}

func (s *scheduler) run(world *World) {
	last := time.Now()

	s.runStage(world, StartupStage, 0)

	for s.isRunning {
		now := time.Now()
		frameTime := now.Sub(last)
		last = now

		s.accumulator += frameTime

		s.runStage(world, PreUpdateStage, frameTime)

		for s.accumulator >= s.tickRate {
			s.runStage(world, FixedUpdateStage, s.tickRate)
			s.accumulator -= s.tickRate
		}

		s.runStage(world, UpdateStage, frameTime)
		s.runStage(world, PostUpdateStage, frameTime)

		s.events.Clear()

		sleepTime := s.tickRate - time.Since(now)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}
	}

	s.runStage(world, EndStage, 0)
}
