package ttlua

import (
	"github.com/yuin/gopher-lua"
	"sync"
	"fmt"
)

//测试GameUser绑定lua
//@lua
type GameUser struct {
	name       string
	id         uint64
	code, CODE string `xml:"Code_Xml" lua:"CODE"`
}

//@lua
func (u *GameUser) TestArray(i int, n ...string) {
	u.name = n[i]
}

//@lua
func (u *GameUser) AnyArray(n ...string) {
	u.name = n[2]
}

//@lua
func (u *GameUser) CloneOne(o *GameUser) {
	u.name = o.name
	u.id = o.id
}

//@lua
func (u *GameUser) SetName(name string) {
	u.name = name
}

//@lua
func (u *GameUser) GetName() string {
	return u.name
}

//@lua
func (u *GameUser) FFFFOOO(name string, id uint64) string {
	u.name = name
	u.id = id
	return u.name
}

//@lua:GameUser
func (u *GameUser) NoNoName() {
}

func (u *GameUser) foo() {
}

//@lua:GameUser
func NewUser(name string) *GameUser {
	/*lll*/
	return &GameUser{name: name}
}

//@lua:GameUser
func NameArray(n ...string) {
	fmt.Println(n)
}

// ssdsd// sss//
//@bindlua:fib(int64)int64
//@bindlua:SetUserName(*GameUser, string)
//@bindlua:InitUser(*GameUser)
//@bindlua:DumpUser(*GameUser)
//===================================================================================

var __lua_once sync.Once
var __L *lua.LState

func GetLuaState() *lua.LState {
	__lua_once.Do(func() {
		__L = lua.NewState()
	})
	return __L
}

