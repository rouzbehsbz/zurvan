package zurvan

import (
	"time"
)

// World represents the main context of the ECS framework.
// It manages entities, components, systems, resources, commands, and events.
type World struct {
	entityAllocator    *entityAllocator
	archetypeAllocator *archetypeAllocator

	scheduler *scheduler
	commands  *commands
	events    *events
	resources *resources

	componentRegistry *registry
}

// Creates a new World instance with the specified tick rate.
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

// Adds systems to the specified stage. Systems will be executed in the order they were added.
func (w *World) AddSystems(stage Stage, systems []System) {
	for _, system := range systems {
		w.scheduler.addSystem(stage, system)
	}
}

// Spawns a new entity and returns it.
// The entity will be created with no components.
//
// Note: This function is completely thread safe
func (w *World) Spawn() Entity {
	return w.entityAllocator.create()
}

// Pushes commands to be executed at the end of the current system execution.
// Commands will be executed in the order they were added.
//
// Note: This function is completely thread safe
func (w *World) PushCommands(commands ...command) {
	for _, command := range commands {
		w.commands.addCommand(command)
	}
}

// Emits events that will be processed after the current system execution.
// Events will be processed in the order they were emitted.
//
// Note: This function is completely thread safe
func (w *World) EmitEvents(events ...any) {
	for _, event := range events {
		w.events.emit(event)
	}
}

// Runs the world, executing systems and processing commands and events.
func (w *World) Run() {
	w.scheduler.run(w)
}
