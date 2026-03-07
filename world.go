package zurvan

import (
	"time"
)

type World struct {
	entityAllocator    *entityAllocator
	archetypeAllocator *archetypeAllocator

	scheduler *scheduler
	commands  *commands
	events    *events
	resources *resources

	componentRegistry *registry
}

func NewWorld(tickRate time.Duration) *World {
	componentRegistry := newRegistry()
	eventRegistry := newRegistry()

	commands := newCommands()
	events := newEvents(eventRegistry)

	return &World{
		entityAllocator:    newEntityAllocator(),
		archetypeAllocator: newArchetypeAllocator(componentRegistry),
		scheduler:          newScheduler(commands, events, tickRate),
		commands:           commands,
		events:             events,
		resources:          newResources(),
		componentRegistry:  componentRegistry,
	}
}

func (w *World) AddSystems(stage Stage, systems []System) {
	for _, system := range systems {
		w.scheduler.addSystem(stage, system)
	}
}

func (w *World) Spawn() Entity {
	return w.entityAllocator.create()
}

func (w *World) PushCommands(commands ...command) {
	for _, command := range commands {
		w.commands.addCommand(command)
	}
}

func (w *World) EmitEvents(events ...any) {
	for _, event := range events {
		w.events.emit(event)
	}
}

func (w *World) Run() {
	w.scheduler.run(w)
}
