package lvm

import (
	lua "github.com/yuin/gopher-lua"
)

type LuaVM struct {
	*lua.LState
}

func NewLuaVM() *LuaVM {
	return &LuaVM{
		LState: lua.NewState(),
	}
}

func (Lvm *LuaVM) PreHandler(context string) (string, string) {
	Lvm.ReadFile("./hooks/prehandler.lua")
	if err := Lvm.CallByParam(lua.P{
		Fn:      Lvm.GetGlobal("PreHandler"),
		NRet:    2,
		Protect: true,
	}, lua.LString(context)); err != nil {
		panic("Fail to exec prehandler")
	}
	return Lvm.Get(-2).String(), Lvm.Get(-1).String()
}

func (Lvm *LuaVM) PastHandler() {
	Lvm.ReadFile("./hooks/pasthandler.lua")
	if err := Lvm.CallByParam(lua.P{
		Fn:      Lvm.GetGlobal("PastHandler"),
		NRet:    0,
		Protect: true,
	}); err != nil {
		panic("Fail to exec pasthandler")
	}
}

func (Lvm *LuaVM) ReadFile(filename string) {
	if err := Lvm.DoFile(filename); err != nil {
		panic(err.Error())
	}
}
