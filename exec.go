package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

var (
	dirStack = []string{}
)

func exReadout(l *lua.LState) int {
	s := l.CheckString(1)
	args := make([]string, l.GetTop()-1)
	for i := 2; i <= l.GetTop(); i++ {
		args[i-2] = l.CheckString(i)
	}
	l.Pop(l.GetTop())
	cmd := exec.Command(s, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	buf := &bytes.Buffer{}
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LFalse)
		l.Push(lua.LNumber(float64(cmd.ProcessState.ExitCode())))
		l.Push(lua.LString(err.Error()))
		return 4
	}
	l.Push(lua.LString(buf.String()))
	l.Push(lua.LTrue)
	l.Push(lua.LNumber(float64(cmd.ProcessState.ExitCode())))
	return 3
}

func exExecute(l *lua.LState) int {
	s := l.CheckString(1)
	args := make([]string, l.GetTop()-1)
	for i := 2; i <= l.GetTop(); i++ {
		args[i-2] = l.CheckString(i)
	}
	l.Pop(l.GetTop())
	cmd := exec.Command(s, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		l.Push(lua.LFalse)
		l.Push(lua.LNumber(float64(cmd.ProcessState.ExitCode())))
		l.Push(lua.LString(err.Error()))
		return 3
	}
	l.Push(lua.LTrue)
	l.Push(lua.LNumber(float64(cmd.ProcessState.ExitCode())))
	return 2
}

func exPushd(l *lua.LState) int {
	target, err := filepath.Abs(l.CheckString(1))
	l.Pop(l.GetTop())
	if err != nil {
		l.Error(lua.LString(err.Error()), 1)
		return 0
	}
	cur, err := os.Getwd()
	if err != nil {
		l.Error(lua.LString(err.Error()), 1)
		return 0
	}
	dirStack = append(dirStack, cur)
	os.Chdir(target)
	l.Push(toLuaTable(l, dirStack...))
	return 1
}

func exPopd(l *lua.LState) int {
	l.Pop(l.GetTop())
	if len(dirStack) == 0 {
		l.Push(lua.LNil)
		return 1
	}
	pop := dirStack[len(dirStack)-1]
	dirStack[len(dirStack)-1] = ""
	dirStack = dirStack[:len(dirStack)-1]
	l.Push(lua.LString(1))
	os.Chdir(pop)
	return 1
}

func toLuaTable(l *lua.LState, args ...string) *lua.LTable {
	t := l.CreateTable(len(args), 0)
	for _, v := range args {
		t.Append(lua.LString(v))
	}
	return t
}
