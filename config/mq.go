package config

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func MQLoader(L *lua.LState) int {
	mq := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"pub":        pub,
		"getPayLoad": getPayLoad,
	})
	L.Push(mq)
	GetLog().Info("Load MQ module to lua")
	return 1
}

func getPayLoad(L *lua.LState) int {
	val := GetMQ().GetPayLoad(L.CheckString(1))
	if val == nil {
		L.Push(lua.LNil)
		return 1
	}
	L.Push(lua.LString(val.(string)))
	return 1
}

func pub(L *lua.LState) int {
	GetMQ().Pub(L.CheckString(1), L.CheckString(2))
	return 0
}

func Consume() {
	timer := time.NewTicker(time.Duration(GetConf().GetInt("mq.consume.time")) * time.Second)
	defer timer.Stop()
	for range timer.C {
		GetLvm().PastHandler()
	}
}
