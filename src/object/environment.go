package object

import "log/slog"

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := Default()
	env.outer = outer
	return env
}

func Default() *Environment {
	logger := slog.Default()
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, Logger: logger}
}

func New(logger *slog.Logger) *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, Logger: logger}
}

type Environment struct {
	store  map[string]Object
	outer  *Environment
	Logger *slog.Logger
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name string, val Object) Object {
	// TODO: Make sure we dont accidentally mutate data that is not in the current scope
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		e.outer.Set(name, val)
		return val
	}
	e.store[name] = val
	return val
}

func (e *Environment) Has(name string) bool {
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Has(name)
	}
	return ok
}

func (e *Environment) Delete(name string) {
	delete(e.store, name)
}
