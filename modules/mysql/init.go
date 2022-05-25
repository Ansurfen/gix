package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	lua "github.com/yuin/gopher-lua"
)

func MySQLLoader(L *lua.LState) int {
	mysql := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"connect": connect,
		"ping":    ping,
		"insert":  insert,
		"delete":  delete,
	})
	L.Push(mysql)
	return 1
}

func connect(L *lua.LState) int {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5)))
	if err != nil {
		panic(err)
	}
	L.Push(&lua.LUserData{Value: &MySQL{DB: db}})
	return 1
}

func ping(L *lua.LState) int {
	db := L.CheckUserData(1).Value.(*MySQL)
	if err := db.Ping(); err != nil {
		panic("fail to ping")
	}
	return 0
}

func insert(L *lua.LState) int {
	db := L.CheckUserData(1).Value.(*MySQL)
	db.insertRow(L.CheckString(2))
	return 0
}

func delete(L *lua.LState) int {
	db := L.CheckUserData(1).Value.(*MySQL)
	db.deleteRow(L.CheckString(2))
	return 0
}
