package main

import (
	"io"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func ioutilWriteText(l *lua.LState) int {
	dest := l.CheckString(1)
	input := l.CheckString(2)
	f, err := os.OpenFile(dest, os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		l.RaiseError("unable to open file: %v", err)
	}
	defer f.Close()
	bytes, err := io.WriteString(f, input)
	if err != nil {
		l.RaiseError("unable to write file: %v", err)
	}
	err = f.Sync()
	if err != nil {
		l.RaiseError("unable to sync file: %v", err)
	}
	l.Push(lua.LNumber(float64(bytes)))
	return 1
}

func ioutilAppendText(l *lua.LState) int {
	dest := l.CheckString(1)
	input := l.CheckString(2)
	f, err := os.OpenFile(dest, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		l.RaiseError("unable to open file: %v", err)
	}
	defer f.Close()
	bytes, err := io.WriteString(f, input)
	if err != nil {
		l.RaiseError("unable to write file: %v", err)
	}
	err = f.Sync()
	if err != nil {
		l.RaiseError("unable to sync file: %v", err)
	}
	l.Push(lua.LNumber(float64(bytes)))
	return 1
}
