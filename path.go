package main

import (
	"path"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func pJoin(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	parts := make([]string, t)
	for i := 1; i <= t; i++ {
		parts[i-1] = strings.TrimSpace(l.CheckString(i))
	}
	l.Pop(t)
	l.Push(lua.LString(path.Join(parts...)))
	return 1
}

func pClean(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	v := l.CheckString(t)
	l.Pop(t)
	l.Push(lua.LString(path.Clean(v)))
	return 1
}
