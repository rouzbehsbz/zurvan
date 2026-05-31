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
		p[i].X += v[i].X * dt
		p[i].Y += v[i].Y * dt
	})
}
```