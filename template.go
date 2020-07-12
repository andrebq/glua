package main

import (
	"bytes"
	"html/template"

	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

func tmplRender(l *lua.LState) int {
	text := l.CheckString(1)
	input := l.CheckAny(2)
	val := gluamapper.ToGoValue(input, gluamapper.Option{
		NameFunc: func(s string) string { return s },
	})
	tmpl, err := template.New("root").Parse(text)
	l.Pop(l.GetTop())
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, val)
	if err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(lua.LString(buf.String()))
	return 1
}
