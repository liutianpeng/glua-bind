/*
自动生成的Lua绑定, 对应GO文件ttlua.go
*/
package ttlua

import "github.com/yuin/gopher-lua"
import "errors"

func init() {
	Luabind_GameUser(GetLuaState())
}
func check_GameUser(L *lua.LState, n int) *GameUser {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*GameUser); ok {
		return v
	}
	L.ArgError(1, "GameUser Expected")
	return nil
}
func luaUserData_GameUser(_L *lua.LState, d *GameUser) *lua.LUserData {
	ld := _L.NewUserData()
	ld.Value = d
	_L.SetMetatable(ld, _L.GetTypeMetatable("GameUser"))
	return ld
}

func Luabind_GameUser(L *lua.LState) {
	mt := L.NewTypeMetatable("GameUser")
	L.SetGlobal("GameUser", mt)

	L.SetField(mt, "NewUser",
		L.NewFunction(func(L *lua.LState) int {
			if L.GetTop() != 1 {
				L.RaiseError("参数数量不对,期望%d个,实际%d个", 1-0, L.GetTop()-0)
				return 0
			}
			r := NewUser(L.CheckString(1))

			ud := L.NewUserData()
			ud.Value = r
			L.SetMetatable(ud, L.GetTypeMetatable("GameUser"))
			L.Push(ud)
			return 1
		}))
	L.SetField(mt, "NameArray",
		L.NewFunction(func(L *lua.LState) int {
			if L.GetTop() == 2 {
				NameArray(L.CheckString(1))
			} else if L.GetTop() == 3 {
				NameArray(L.CheckString(1), L.CheckString(2))
			} else if L.GetTop() == 4 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3))
			} else if L.GetTop() == 5 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4))
			} else if L.GetTop() == 6 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5))
			} else if L.GetTop() == 7 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6))
			} else if L.GetTop() == 8 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7))
			} else if L.GetTop() == 9 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8))
			} else if L.GetTop() == 10 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9))
			} else if L.GetTop() == 11 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10))
			} else if L.GetTop() == 12 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11))
			} else if L.GetTop() == 13 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12))
			} else if L.GetTop() == 14 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13))
			} else if L.GetTop() == 15 {
				NameArray(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14))
			} else {
				L.RaiseError("参数数量不对...")
				return 0
			}

			return 0
		}))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(),
		map[string]lua.LGFunction{

			"CODE": func(L *lua.LState) int {
				p := check_GameUser(L, 1)

				if L.GetTop() > 2 {
					L.RaiseError("参数数量不对...")
					return 0
				}
				if L.GetTop() == 2 {
					p.CODE = L.CheckString(2)
					return 0
				}
				r := p.CODE

				L.Push(lua.LString(r))
				return 1
			},

			"TestArray": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() == 3 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3))
				} else if L.GetTop() == 4 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4))
				} else if L.GetTop() == 5 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5))
				} else if L.GetTop() == 6 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6))
				} else if L.GetTop() == 7 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7))
				} else if L.GetTop() == 8 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8))
				} else if L.GetTop() == 9 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9))
				} else if L.GetTop() == 10 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10))
				} else if L.GetTop() == 11 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11))
				} else if L.GetTop() == 12 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12))
				} else if L.GetTop() == 13 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13))
				} else if L.GetTop() == 14 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14))
				} else if L.GetTop() == 15 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14), L.CheckString(15))
				} else if L.GetTop() == 16 {
					p.TestArray(int(L.CheckInt(2)), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14), L.CheckString(15), L.CheckString(16))
				} else {
					L.RaiseError("参数数量不对...")
					return 0
				}

				return 0
			},

			"AnyArray": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() == 2 {
					p.AnyArray(L.CheckString(2))
				} else if L.GetTop() == 3 {
					p.AnyArray(L.CheckString(2), L.CheckString(3))
				} else if L.GetTop() == 4 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4))
				} else if L.GetTop() == 5 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5))
				} else if L.GetTop() == 6 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6))
				} else if L.GetTop() == 7 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7))
				} else if L.GetTop() == 8 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8))
				} else if L.GetTop() == 9 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9))
				} else if L.GetTop() == 10 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10))
				} else if L.GetTop() == 11 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11))
				} else if L.GetTop() == 12 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12))
				} else if L.GetTop() == 13 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13))
				} else if L.GetTop() == 14 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14))
				} else if L.GetTop() == 15 {
					p.AnyArray(L.CheckString(2), L.CheckString(3), L.CheckString(4), L.CheckString(5), L.CheckString(6), L.CheckString(7), L.CheckString(8), L.CheckString(9), L.CheckString(10), L.CheckString(11), L.CheckString(12), L.CheckString(13), L.CheckString(14), L.CheckString(15))
				} else {
					L.RaiseError("参数数量不对...")
					return 0
				}

				return 0
			},

			"CloneOne": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() != 2 {
					L.RaiseError("参数数量不对,期望%d个,实际%d个", 2-1, L.GetTop()-1)
					return 0
				}
				p.CloneOne(check_GameUser(L, 2))

				return 0
			},

			"SetName": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() != 2 {
					L.RaiseError("参数数量不对,期望%d个,实际%d个", 2-1, L.GetTop()-1)
					return 0
				}
				p.SetName(L.CheckString(2))

				return 0
			},

			"GetName": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() != 1 {
					L.RaiseError("参数数量不对,期望%d个,实际%d个", 1-1, L.GetTop()-1)
					return 0
				}
				r := p.GetName()

				L.Push(lua.LString(r))
				return 1
			},

			"FFFFOOO": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() != 3 {
					L.RaiseError("参数数量不对,期望%d个,实际%d个", 3-1, L.GetTop()-1)
					return 0
				}
				r := p.FFFFOOO(L.CheckString(2), uint64(L.CheckInt64(3)))

				L.Push(lua.LString(r))
				return 1
			},

			"NoNoName": func(L *lua.LState) int {
				p := check_GameUser(L, 1)
				if L.GetTop() != 1 {
					L.RaiseError("参数数量不对,期望%d个,实际%d个", 1-1, L.GetTop()-1)
					return 0
				}
				p.NoNoName()

				return 0
			},
		}))
}

//自动生成调用lua的方法InitUser
func CallLua_InitUser(_L *lua.LState, p0 *GameUser) error {
	fn, ok := _L.GetGlobal("InitUser").(*lua.LFunction)
	if ok == false {
		return errors.New("找不到函数:InitUser")
	}
	_L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, luaUserData_GameUser(_L, p0))

	return nil
}

//自动生成调用lua的方法DumpUser
func CallLua_DumpUser(_L *lua.LState, p0 *GameUser) error {
	fn, ok := _L.GetGlobal("DumpUser").(*lua.LFunction)
	if ok == false {
		return errors.New("找不到函数:DumpUser")
	}
	_L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, luaUserData_GameUser(_L, p0))

	return nil
}

//自动生成调用lua的方法fib
func CallLua_fib(_L *lua.LState, p0 int64) (ret int64, err error) {
	fn, ok := _L.GetGlobal("fib").(*lua.LFunction)
	if ok == false {
		return ret, errors.New("找不到函数:fib")
	}
	_L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, lua.LNumber(float64(p0)))

	v := _L.Get(-1)
	if intv, ok := v.(lua.LNumber); ok {
		_L.Pop(1)
		return int64(intv), nil
	}
	_L.Pop(1)
	return ret, errors.New("返回值类型不对")
}

//自动生成调用lua的方法SetUserName
func CallLua_SetUserName(_L *lua.LState, p0 *GameUser, p1 string) error {
	fn, ok := _L.GetGlobal("SetUserName").(*lua.LFunction)
	if ok == false {
		return errors.New("找不到函数:SetUserName")
	}
	_L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, luaUserData_GameUser(_L, p0), lua.LString(p1))

	return nil
}
