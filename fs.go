package main

import (
	"io"
	"os"

	copydir "github.com/otiai10/copy"
	lua "github.com/yuin/gopher-lua"
)

func fsCopydir(l *lua.LState) int {
	dest := l.CheckString(1)
	src := l.CheckString(2)
	l.Pop(2)

	err := copydir.Copy(src, dest)
	l.Push(lua.LBool(err == nil))
	if err != nil {
		l.RaiseError("unable to copy directory: %v", err)
		return 0
	}
	return 1
}

func fsCopyfile(l *lua.LState) int {
	dest := l.CheckString(1)
	src := l.CheckString(2)
	l.Pop(2)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		l.RaiseError("unable to stat file %v: %v", src, err)
	}

	if !sourceFileStat.Mode().IsRegular() {
		l.RaiseError("file %v is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		l.RaiseError("unable to open %v for reading: %v", src, err)
	}
	defer source.Close()

	destination, err := os.Create(dest)
	if err != nil {
		l.RaiseError("unable to open %v for writing: %v", dest, err)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		l.RaiseError("unable to copy content: %v", err)
	}
	err = destination.Sync()
	if err != nil {
		l.RaiseError("unable to sync destination to disk: %v", err)
	}
	l.Push(lua.LNumber(float64(nBytes)))
	return 1
}

func fsRemoveAll(l *lua.LState) int {
	target := l.CheckString(1)
	l.Pop(1)

	err := os.RemoveAll(target)
	l.Push(lua.LBool(err == nil))
	if err != nil {
		l.RaiseError("unable to remove directory: %v", err)
		return 0
	}
	return 1
}

func fsMkdir(l *lua.LState) int {
	target := l.CheckString(1)
	perm := l.CheckInt(2)
	l.Pop(2)
	err := os.MkdirAll(target, os.FileMode(perm))
	l.Push(lua.LBool(err == nil))
	if err != nil {
		l.RaiseError("unable to create directory: %v", err)
		return 0
	}
	return 1
}
