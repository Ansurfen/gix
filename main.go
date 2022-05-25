package main

import (
	. "gix/api"
	. "gix/config"
)

func main() {
	Conf, Logger, MQ, Lvm = InitGixConf(), GetLog(), InitGixMQ(), InitGixLvm()
	defer Lvm.Close()
	go Consume()
	gix := NewApp().Defalut().UseRouter()
	panic(gix.Run())
}
