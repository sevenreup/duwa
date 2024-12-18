package object

import (
	"log/slog"

	"github.com/sevenreup/duwa/src/library/runtime"
	"github.com/sevenreup/duwa/src/runtime/native"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := Default()
	env.outer = outer
	return env
}

func Default() *Environment {
	logger := slog.Default()
	s := make(map[string]Object)
	console := native.NewConsole()
	return &Environment{store: s, outer: nil, Logger: logger, Console: console}
}

func New(logger *slog.Logger, console runtime.Console) *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, Logger: logger, Console: console}
}

type Environment struct {
	store   map[string]Object
	outer   *Environment
	Logger  *slog.Logger
	Console runtime.Console
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

func (environment *Environment) Call(function string, args []Object) Object {
	if object, ok := environment.Get(function); ok {
		if function, ok := object.(*Function); ok {
			return function.Evaluate(environment, args)
		}
	}

	return NewError("function not found: %s", function)
}
