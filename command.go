package zurvan

const CommandBufferSize int = 1024

type command interface {
	execute(w *World)
}

type commands struct {
	commands chan command
}

func newCommands() *commands {
	return &commands{
		commands: make(chan command, CommandBufferSize),
	}
}

func (c *commands) addCommand(command command) {
	c.commands <- command
}

func (c *commands) apply(w *World) {
	for len(c.commands) > 0 {
		command := <-c.commands
		command.execute(w)
	}
}

type setComponentsCommand struct {
	entity     Entity
	components []any
}

func NewSetComponentsCommand(entity Entity, components ...any) *setComponentsCommand {
	return &setComponentsCommand{
		entity:     entity,
		components: components,
	}
}

func (s *setComponentsCommand) execute(w *World) {
	w.archetypeAllocator.addComponents(s.entity, s.components...)
}

type deleteComponentsCommand struct {
	entity     Entity
	components []any
}

func NewDeleteComponentsCommand(entity Entity, components ...any) *deleteComponentsCommand {
	return &deleteComponentsCommand{
		entity:     entity,
		components: components,
	}
}

func (d *deleteComponentsCommand) execute(w *World) {
	w.archetypeAllocator.deleteComponents(d.entity, d.components...)
}

type addResourceCommand struct {
	resource any
}

func NewAddResourceCommand(resource any) *addResourceCommand {
	return &addResourceCommand{
		resource: resource,
	}
}

func (a *addResourceCommand) execute(w *World) {
	w.resources.addResource(a.resource)
}

type despawnCommand struct {
	entity Entity
}

func NewDespawnCommand(entity Entity) *despawnCommand {
	return &despawnCommand{
		entity: entity,
	}
}

func (d *despawnCommand) execute(w *World) {
	w.archetypeAllocator.removeEntity(d.entity)
}
