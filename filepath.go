package main

import (
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

func fpJoin(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	parts := make([]string, t)
	for i := 1; i <= t; i++ {
		parts[i-1] = l.CheckString(i)
	}
	l.Pop(t)
	l.Push(lua.LString(filepath.Join(parts...)))
	return 1
}

func fpClean(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	v := l.CheckString(t)
	l.Pop(t)
	l.Push(lua.LString(filepath.Clean(v)))
	return 1
}

func fpFromSlash(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	v := l.CheckString(t)
	l.Pop(t)
	l.Push(lua.LString(filepath.FromSlash(v)))
	return 1
}

func fpSplit(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	v := l.CheckString(t)
	l.Pop(t)
	dir, file := filepath.Split(v)
	l.Push(lua.LString(dir))
	l.Push(lua.LString(file))
	return 1
}

func fpSplitList(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	v := l.CheckString(t)
	l.Pop(t)
	items := filepath.SplitList(v)
	tbl := l.CreateTable(len(items), 0)
	for _, v := range items {
		tbl.Append(lua.LString(v))
	}
	l.Push(tbl)
	return 1
}
