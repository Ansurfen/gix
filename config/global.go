package config

import (
	"github.com/spf13/viper"
	lua "gix/lvm"
	"gix/modules/mysql"
	"gix/mq"
)

var Lvm *lua.LuaVM
var MQ *mq.Broker
var Conf *viper.Viper

func GetLvm() *lua.LuaVM {
	if Lvm == nil {
		GetLog().Info("Prepare to create Lvm")
		return lua.NewLuaVM()
	}
	return Lvm
}

func GetMQ() *mq.Broker {
	if MQ == nil {
		GetLog().Info("Prepare to create MQ")
		cap := GetConf().GetInt("mq.cap")
		if cap <= 0 {
			cap = 10
			GetLog().Warn("Reset MQ cap")
		}
		return mq.NewBroker().SetCap(cap)
	}
	return MQ
}

func GetConf() *viper.Viper {
	if Conf == nil {
		return viper.New()
	}
	return Conf
}

func InitGixConf() *viper.Viper {
	Conf = GetConf()
	Conf.AddConfigPath(".")
	Conf.SetConfigFile("gix.yml")
	if err := Conf.ReadInConfig(); err != nil {
		panic("Fail to open config")
	}
	return Conf
}

func InitGixMQ() *mq.Broker {
	MQ = GetMQ()
	MQ.MultiSub(GetConf().GetStringSlice("mq.sub"))
	return MQ
}

func InitGixLvm() *lua.LuaVM {
	Lvm = GetLvm()
	Lvm.PreloadModule("gix-mq", MQLoader)
	Lvm.PreloadModule("lua-mysql", mysql.MySQLLoader)
	return Lvm
}
