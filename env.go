package main

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func envSet(l *lua.LState) int {
	name := l.CheckString(1)
	value := l.CheckString(2)
	l.Pop(2)
	err := os.Setenv(name, value)
	if err != nil {
		l.RaiseError(fmt.Sprintf("%v", err))
	}
	return 0
}

func envGet(l *lua.LState) int {
	name := l.CheckString(1)
	value := os.Getenv(name)
	l.Push(lua.LString(value))
	return 1
}
