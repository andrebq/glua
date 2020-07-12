package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

func b64EncodeFile(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	file := l.CheckString(1)
	enc := "std"
	if t == 2 {
		enc = l.CheckString(2)
	}
	l.Pop(t)
	buf, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	ud := l.NewUserData()
	ud.Value = buf
	l.Push(ud)
	l.Push(lua.LString(enc))
	return b64Encode(l)
}

func b64Encode(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString(""))
		return 1
	}
	buf, ok := l.CheckUserData(1).Value.([]byte)
	if !ok {
		l.Error(lua.LString("value is not a byte slice"), 1)
		return 0
	}
	encoding := base64.StdEncoding
	if t == 2 {
		switch v := l.CheckString(t); v {
		case "url":
			encoding = base64.URLEncoding
		case "std":
			encoding = base64.StdEncoding
		default:
			l.Error(lua.LString(fmt.Sprintf("%v is not a valid encoding", v)), 1)
			return 0
		}
	}
	l.Pop(t)
	l.Push(lua.LString(encoding.EncodeToString(buf)))
	return 1
}

func b64Decode(l *lua.LState) int {
	t := l.GetTop()
	if t == 0 {
		l.Push(lua.LString("Missing encoded string"))
		return 1
	}
	encoded := l.CheckString(1)
	encoding := base64.StdEncoding
	if t == 2 {
		switch v := l.CheckString(t); v {
		case "url":
			encoding = base64.URLEncoding
		case "std":
			encoding = base64.StdEncoding
		default:
			l.Error(lua.LString(fmt.Sprintf("%v is not a valid encoding", v)), 1)
			return 0
		}
	}
	buf, err := encoding.DecodeString(encoded)
	l.Pop(t)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	ud := l.NewUserData()
	ud.Value = buf
	l.Push(ud)
	return 1
}
