# Zurvan

[![Go Reference](https://pkg.go.dev/badge/github.com/rouzbehsbz/zurvan.svg)](https://pkg.go.dev/github.com/rouzbehsbz/zurvan)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rouzbehsbz/zurvan)
[![Build Status](https://github.com/rouzbehsbz/zurvan/actions/workflows/build.yml/badge.svg)](https://github.com/rouzbehsbz/zurvan/actions/workflows/build.yml)

> Zurvan (زُروان) — pronounced ZOOR-vahn — a Persian name of Avestan origin meaning “Time,” often associated with eternity and the primordial source of existence.

This library implements the Entity Component System (ECS) architecture. ECS is based on data-oriented design principles. Instead of organizing code around objects, it focuses on organizing data in a cache-friendly memory layout. This approach improves CPU cache utilization and can significantly enhance performance for both read and write operations. ECS is especially useful for performance-critical applications such as game engines or large-scale simulations, where a high number of entities with diverse behaviors must be processed efficiently.

## Overview
In ECS, entities are simply unique identifiers that group components together. They do not contain any logic themselves. Systems are pure functions that operate on components. They query for entities that match a specific set of components and then execute logic on them, potentially mutating their state. For example, you might have entities with `Position` and `Velocity` components. A `MovementSystem` would query all entities that contain both components, apply movement logic, and update their positions accordingly.

Here’s a simple example:
```go
func (m *MovementSystem) Update(w *zurvan.World, dt time.Duration) {
	dt := dt.Seconds()

	zurvan.QueryMany2(w, func(e []zurvan.Entity, p []Position, v []Velocity) {
		for i, e := range entities {
			p[i].X += v[i].X * sec
			p[i].Y += v[i].Y * sec
		}
	})
}
```

## Concepts

### Entity
It is simply a unique identifier that represents an object in the world. you can define a player or a bullet or any object that can have some property and behavior as an entity.

```go
player := world.Spawn()
bullet := world.Spawn()
```

### Component
Components contain data only, no behavior or logic associated to them. you can set, or delete components for spawned entities in your world.

```go
type Position struct {
	X, Y float64
}

type Health struct {
	Value int
}
```

### System
A system contains the logic that operates on components. you can query entities that match a specific set of components or query an exact entity for retrieving its components and read or update them. Systems execute in the order and stage they were added during each loop.

```go
type MovementSystem struct {}

func (m *MovementSystem) Update(w *zurvan.World, dt time.Duration) {
	dt := dt.Seconds()

	zurvan.QueryMany2(w, func(e []zurvan.Entity, p []Position, v []Velocity) {
		for i, e := range entities {
			p[i].X += v[i].X * sec
			p[i].Y += v[i].Y * sec
		}
	})
}
```

### Command
Despawning an entity, adding or removing components, and managing resources can be executed as commands. For system reliability in every tick or frame, these operations are applied at the end of each system. This ensures they are atomic and that the next stage operates on a consistent state.

```go
w.PushCommands(
	zurvan.NewSetComponentsCommand(player,
		Position{X: 10, Y: 10},
		Velocity{X: 2, Y: 2},
	),
)
```

### Resource
A resource is a globally shared piece of state that can be accessed from anywhere in the world, such as configuration, time or global game state.

```go
// Register a new resource

type Players struct {
	Map map[int]zurvan.Entity
}

players := &Players{}

world.PushCommands(
	zurvan.NewAddResourceCommand(players),
)

// Getting a registered resource by its type

zurvan.Resource[*Players](world)
```

### Event
An event is a message that represents something that happened during a single tick. It can be emitted and should be consumed within the same tick by upstream systems. At the end of the tick, all events are cleared. Events are used for decoupling systems and a way of communication between them.

```go
// Pushing an event

if health[i].Value <= 0 {
	world.EmitEvents(
		DeathEvent{
			Entity: e,
		},
	)
}

// Retrieving events

events := zurvan.OnEvent[DeathEvent](world)
```

You can find a complete example demonstrating all of these concepts in the `examples/basic` directory.