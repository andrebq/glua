# glua - go + lua

This is a very basic shell to serve as a cross platform automation helper.

## How it works?

Most of the content from `main.go` are a simple copy/paste of
[yuin/gopher-lua/cmd/glua](https://github.com/yuin/gopher-lua/blob/master/cmd/glua/glua.go).

The main difference is the list of available libraries, since this project
is targeted towards making automation easy to use, some libs are pre-loaded:

- **filepath** to provide filesystem manipulation
- **path** to provide path manipulation
