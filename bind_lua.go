package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"os"
	"go/format"
	"path"
)

func getLuaType(goType string) string {
	switch goType {
	case "string":
		return "lua.LString"
	case "uint64", "uint32", "int64", "int32", "int", "float32", "float64":
		return "lua.LNumber"
	case "bool":
		return "lua.LBool"
	default:
		return "lua.LUserData"
	}
}

var Cfg = &LuaBindCfg{}

type LuaBindCfg struct {
	OutPackageName string
	InPackageName  string
	AutoInit       bool
	ImportString   []string

	initFunction []string
	bindList     string
}

//生成调用go的参数
func _buildGoCallParam(gc *GoCodeStruct, fc *GoCodeFunction, param GoCodeParam, pIndex int) string {
	call_params := ``
	pNum := strconv.FormatInt(int64(pIndex+1), 10)
	if gc != nil {
		pNum = strconv.FormatInt(int64(pIndex+2), 10)
	}
	if pIndex > 0 {
		call_params += ", "
	}

	switch param.TypeName {
	case "string":
		call_params += strings.Replace("L.CheckString(#PM#)", "#PM#", pNum, -1)
	case "uint64", "int64":
		call_params += strings.Replace("#PT#(L.CheckInt64(#PM#))", "#PM#", pNum, -1)
	case "uint32", "uint", "int32", "int":
		call_params += strings.Replace("#PT#(L.CheckInt(#PM#))", "#PM#", pNum, -1)
	case "float32", "float64":
		call_params += strings.Replace("#PT#(L.CheckNumber(#PM#))", "#PM#", pNum, -1)
	case "bool":
		call_params += strings.Replace("#PT#(L.CheckBool(#PM#))", "#PM#", pNum, -1)
	default:
		s := strings.Replace(`check_#PType#(L, #PM#)`, "#PM#", pNum, -1)
		s = strings.Replace(s, "#PType#", param.TypeName, -1)
		call_params += s
	}
	call_params = strings.Replace(call_params, "#PT#", param.TypeName, -1)
	return call_params
}

type GoCodeParam struct {
	ParamName  string
	TypeName   string
	IsEllipsis bool
}

type GoCodeFunction struct {
	GoName                string
	LuaName               string
	Params                []GoCodeParam
	Returns               []string
	GetterSetter          bool
	GetterSetter_Star     bool
	GetterSetter_UserData bool
	CallLuaFunction       bool
}

func (fc *GoCodeFunction) String() string {
	if fc.GetterSetter{
		str := "属性:"+fc.GoName
		for _, v := range fc.Returns {
			str += " " + v
		}
		return str
	}else {
		str := fc.GoName
		str += "("
		for k, v := range fc.Params {
			if k > 0 {
				str += ", "
			}
			str += v.TypeName
		}
		str += ")"
		for _, v := range fc.Returns {
			str += " " + v
		}
		return str
	}
}

func (fc *GoCodeFunction) BuildLuaBind(gc *GoCodeStruct) string {
	isEllipsis := false
	for _, p := range fc.Params {
		if p.IsEllipsis {
			isEllipsis = true
		}
	}

	methods_call := ``
	method_code := ``

	if isEllipsis == true {
		for pn := 1; pn < 15; pn++ {
			call_params := ``

			one_call_go := ``
			if len(fc.Returns) > 0 {
				one_call_go = `if L.GetTop() == #ParamArrayNum# {
					r := #GoCaller##GoName#(#call_params#)
				}`
			} else {
				one_call_go = `if L.GetTop() == #ParamArrayNum# {
					#GoCaller##GoName#(#call_params#)
				}`
			}
			if pn > 1 {
				one_call_go = "else " + one_call_go
			}
			topNum := 0
			for k, p := range fc.Params {
				topNum = k
				if p.IsEllipsis {
					for i := 0; i < pn; i++ {
						call_params += _buildGoCallParam(gc, fc, p, k+i)
					}
				} else {
					call_params += _buildGoCallParam(gc, fc, p, k)
				}
			}
			one_call_go = strings.Replace(one_call_go, "#call_params#", call_params, -1)
			one_call_go = strings.Replace(one_call_go, "#GoName#", fc.GoName, -1)
			one_call_go = strings.Replace(one_call_go, "#ParamArrayNum#", strconv.FormatUint(uint64(topNum+pn+1), 10), -1)

			methods_call += one_call_go
		}
		methods_call += `else{
	L.RaiseError("参数数量不对...")
	return 0
}`
	} else if fc.CallLuaFunction {
		methods_call = `
//自动生成调用lua的方法#LuaFuncName#
func #GoFuncName#(_L *lua.LState#GoParams#) (#GoReturns# error) {
	fn, ok := _L.GetGlobal("#LuaFuncName#").(*lua.LFunction)
	if ok == false {
		return #DefaultReturn# errors.New("找不到函数:#LuaFuncName#")
	}
	err:=_L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	} #LuaParams#)

	#LuaReturns#
}
`
		GoParams := ""
		for _, v := range fc.Params {
			GoParams += ", "
			GoParams += v.ParamName + " " + v.TypeName
		}
		GoReturns := ""
		DefaultReturn := ""
		if len(fc.Returns) == 1 {
			GoReturns = "ret " + fc.Returns[0] + ", err"
			DefaultReturn = "ret, "
		}
		LuaParams := ""
		for _, v := range fc.Params {
			LuaParams += ","
			switch v.TypeName {
			case "string":
				LuaParams += "lua.LString(" + v.ParamName + ")"
			case "uint64", "uint32", "int64", "int32", "int", "float32", "float64":
				LuaParams += "lua.LNumber(float64(" + v.ParamName + "))"
			case "bool":
				LuaParams += "lua.LBool((" + v.ParamName + "))"
			default:
				t := strings.TrimPrefix(v.TypeName, "*")
				st := strings.Split(t, ".")
				t = st[len(st)-1]
				LuaParams += "luaUserData_" + t + "(_L," + v.ParamName + ")"
			}
		}
		LuaReturns := ""
		if len(fc.Returns) == 1 {
			LuaReturns = `
			v := _L.Get(-1)
			if intv, ok := v.(#LType#); ok {
				_L.Pop(1)
				return #GType#(intv), err
			}
			_L.Pop(1)
			return ret, errors.New("返回值类型不对")`
			LuaReturns = strings.Replace(LuaReturns, "#LType#", getLuaType(fc.Returns[0]), -1)
			LuaReturns = strings.Replace(LuaReturns, "#GType#", fc.Returns[0], -1)
		} else {
			LuaReturns = `return err`
		}
		methods_call = strings.Replace(methods_call, "#GoFuncName#", fc.GoName, -1)
		methods_call = strings.Replace(methods_call, "#LuaFuncName#", fc.LuaName, -1)
		methods_call = strings.Replace(methods_call, "#GoParams#", GoParams, -1)
		methods_call = strings.Replace(methods_call, "#LuaParams#", LuaParams, -1)
		methods_call = strings.Replace(methods_call, "#LuaReturns#", LuaReturns, -1)
		methods_call = strings.Replace(methods_call, "#GoReturns#", GoReturns, -1)
		methods_call = strings.Replace(methods_call, "#DefaultReturn#", DefaultReturn, -1)

		//#GoParams#
		//#GoReturns#
		//#LuaParams#	lua.LNumber(float64(n))
		//#LuaReturns#
		return methods_call
	} else if fc.GetterSetter {
		if fc.GetterSetter_Star {
			methods_call = `
	if L.GetTop() > 2 {
		L.RaiseError("参数数量不对...")
		return 0
	}
    if L.GetTop() == 2 {
        #GoCaller##GoName# = #call_params#
        return 0
    }
	r:=#GoCaller##GoName# 
`
		} else {
			methods_call = `
	if L.GetTop() > 2 {
		L.RaiseError("参数数量不对...")
		return 0
	}
    if L.GetTop() == 2 {
        #GoCaller##GoName# = *#call_params#
        return 0
    }
	r:=&#GoCaller##GoName# 
`
		}
		call_params := ``
		for k, p := range fc.Params {
			call_params += _buildGoCallParam(gc, fc, p, k)
		}
		methods_call = strings.Replace(methods_call, "#GoName#", fc.GoName, -1)
		methods_call = strings.Replace(methods_call, "#call_params#", call_params, -1)
	} else {

		call_params := ``
		for k, p := range fc.Params {
			call_params += _buildGoCallParam(gc, fc, p, k)
		}

		methods_call = `if L.GetTop() != #ParamNum# {
					L.RaiseError("参数数量不对,期望%d个,实际%d个",#ParamNum#-#ThisNum#, L.GetTop()-#ThisNum#)
					return 0
				}
				#call_go#
`

		call_go := ""

		if len(fc.Returns) > 0 {
			call_go = `r := #GoCaller##GoName#(#call_params#)`
		} else {
			call_go = `#GoCaller##GoName#(#call_params#)`
		}
		call_go = strings.Replace(call_go, "#GoName#", fc.GoName, -1)
		call_go = strings.Replace(call_go, "#call_params#", call_params, -1)

		methods_call = strings.Replace(methods_call, "#call_go#", call_go, -1)

	}

	methods_return := ``
	for _, r := range fc.Returns {
		//L.Push(lua.LString(p.GetName()))
		switch r {
		case "string":
			methods_return += "L.Push(lua.LString(r))"
		case "uint64", "uint32", "int64", "int32", "int", "float32", "float64":
			methods_return += "L.Push(lua.LNumber(r))"
		case "bool":
			methods_return += "L.Push(lua.LBool(r))"
		default:
			ret_str := `ud := L.NewUserData()
				ud.Value = r
				L.SetMetatable(ud, L.GetTypeMetatable("#ReturnType#"))
				L.Push(ud)`
			methods_return += strings.Replace(ret_str, "#ReturnType#", r, -1)
		}
	}

	if gc == nil {
		methods_call = strings.Replace(methods_call, "#GoCaller#", "", -1)

		methods_call = strings.Replace(methods_call, "#ParamNum#", strconv.FormatInt(int64(len(fc.Params)), 10), -1)
		methods_call = strings.Replace(methods_call, "#ThisNum#", "0", -1)

		method_code = `
		L.SetField(mt, "#GoName#",
			L.NewFunction(func(L *lua.LState) int {
				#methods_call#
				#methods_return#
				return #return_num#
			}))`

	} else {
		methods_call = strings.Replace(methods_call, "#GoCaller#", "p.", -1)

		methods_call = strings.Replace(methods_call, "#ParamNum#", strconv.FormatInt(int64(len(fc.Params)+1), 10), -1)
		methods_call = strings.Replace(methods_call, "#ThisNum#", "1", -1)

		method_code = `
		"#GoName#": func(L *lua.LState) int {
				p := check_#Type#(L, 1)
				#methods_call#
				#methods_return#
				return #return_num#
			},
`
	}

	src := method_code
	src = strings.Replace(src, "#LuaName#", fc.LuaName, -1)
	src = strings.Replace(src, "#GoName#", fc.GoName, -1)
	src = strings.Replace(src, "#methods_call#", methods_call, -1)
	src = strings.Replace(src, "#methods_return#", methods_return, -1)
	src = strings.Replace(src, "#return_num#", strconv.FormatInt(int64(len(fc.Returns)), 10), -1)
	return src
}

type GoCodeStruct struct {
	GoName          string
	GoNameWithPack  string
	LuaName         string
	Functions       []GoCodeFunction
	StaticFunctions []GoCodeFunction

	foundStruct bool
}

func (gc *GoCodeStruct) BuildLuaBind() string {
	Cfg.bindList += "\n=================================\n绑定对象:" + gc.GoName + "\n"
	init_code := `luabind_#Type#(L/*LuaState*/)`
	init_code = strings.Replace(init_code, "#Type#", gc.GoName, -1)
	Cfg.initFunction = append(Cfg.initFunction, init_code)

	check_code := `
func check_#Type#(L *lua.LState, n int) *#TypeWithPackage# {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*#TypeWithPackage#); ok {
		return v
	}
	L.ArgError(1, "#Type# Expected")
	return nil
}
func luaUserData_#Type#(_L *lua.LState, d *#TypeWithPackage#) *lua.LUserData {
	ld := _L.NewUserData()
	ld.Value = d
	_L.SetMetatable(ld, _L.GetTypeMetatable("#Type#"))
	return ld
}

`

	if Cfg.AutoInit {
		//check_code = init_code + check_code
	} else {
		//check_code = "/*调用这个函数注册Go对象到Lua\n" + init_code + "\n*/" + check_code
	}

	bind_code := `
func luabind_#Type#(L *lua.LState) {
	mt := L.NewTypeMetatable("#Type#")
	L.SetGlobal("#Type#", mt)
	#static_code#
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(),
		map[string]lua.LGFunction{
		#methods#
	}))
}
`
	static_code := ``
	Cfg.bindList += "静态函数列表：\n"
	for _, f := range gc.StaticFunctions {
		Cfg.bindList += "	" + f.String() + "\n"
		static_code += f.BuildLuaBind(nil)
	}

	methods := ""
	Cfg.bindList += "函数列表：\n"
	for _, f := range gc.Functions {
		Cfg.bindList += "	" + f.String() + "\n"
		methods += f.BuildLuaBind(gc)
	}

	src := `

#check_code#
#bind_code#
`
	src = strings.Replace(src, "#check_code#", check_code, -1)
	src = strings.Replace(src, "#bind_code#", bind_code, -1)
	src = strings.Replace(src, "#static_code#", static_code, -1)
	src = strings.Replace(src, "#methods#", methods, -1)

	src = strings.Replace(src, "#Type#", gc.GoName, -1)
	src = strings.Replace(src, "#TypeWithPackage#", gc.GoNameWithPack, -1)

	return src
}

var code = make(map[string]*GoCodeStruct)
var callLua = make(map[string]*GoCodeFunction)

func MakeLuaBindFile(sourceFile string, outDir string) {
	dat, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		panic(err)
	}
	src := string(dat)
	code := MakeLuaBindCode(src)
	code = strings.Replace(code, "#GoSrcFile#", sourceFile, -1)
	code = strings.Replace(code, "#RegFuncName#", "LuaBind_"+strings.Split(path.Base(sourceFile), ".")[0], -1)

	fmtCode, err := format.Source([]byte(code))
	if err != nil {
		fmt.Println(err)
		fmtCode = []byte(code)
	}

	//fmt.Println(string(fmtCode))
	//return

	outFile := "lua_" + path.Base(sourceFile)
	outFile = path.Join(outDir, outFile)
	f, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(f, string(fmtCode))
	f.Close()
}

func MakeLuaBindCode(src string) string {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	//ast.Print(fset, f)
	for _, v := range file.Comments {
		for _, c := range v.List {
			if strings.HasPrefix(strings.TrimLeft(c.Text, "// "), "@bindlua:") {
				decl := strings.Split(c.Text, ":")[1]
				ss := strings.Split(decl, "(")
				if len(ss) == 2 {
					fc := &GoCodeFunction{}
					fc.GoName = "CallLua_" + ss[0]
					fc.LuaName = ss[0]
					fc.CallLuaFunction = true
					ps := strings.Split(ss[1], ")")
					for k, p := range strings.Split(ps[0], ",") {
						fc.Params = append(fc.Params, GoCodeParam{
							ParamName: "p" + strconv.FormatInt(int64(k), 10),
							TypeName:  strings.Trim(p, " "),
						})
					}
					if len(ps) == 2 && len(strings.Trim(ps[1], " ")) > 0 {
						fc.Returns = append(fc.Returns, ps[1])
					}
					//fmt.Println(fc)
					callLua[fc.LuaName] = fc
					Cfg.ImportString = append(Cfg.ImportString, `errors`)
				}
			}
		}
	}
	for _, v := range file.Decls {
		//fmt.Println("--")
		if f, ok := v.(*ast.FuncDecl); ok {
			if f.Doc != nil &&
				len(f.Doc.List) > 0 {
				fc := &GoCodeFunction{}
				foundFunction := false

				for _, v := range f.Doc.List {
					index := strings.Index(v.Text, "@lua")
					if index > 0 {
						fc.GoName = f.Name.Name
						fc.LuaName = f.Name.Name
						ss := strings.Split(v.Text[index+4:], ":")
						if len(ss) > 1 {
							fc.LuaName = ss[1]
						}
						foundFunction = true
					} else {
						foundFunction = false
					}
				}

				if f.Type.Params != nil {
					for _, v := range f.Type.Params.List {
						if ti, ok := v.Type.(*ast.Ident); ok {
							//fmt.Println(ti.Name)
							for _, p := range v.Names {
								//fmt.Println(p.Name + " --> " + ti.Name)
								fc.Params = append(fc.Params, GoCodeParam{ParamName: p.Name, TypeName: ti.Name})
							}
						} else if ti, ok := v.Type.(*ast.StarExpr); ok {
							if x, ok := ti.X.(*ast.Ident); ok {
								for _, p := range v.Names {
									//fmt.Println(p.Name + " --> " + ti.Name)
									fc.Params = append(fc.Params, GoCodeParam{ParamName: p.Name, TypeName: x.Name})
								}
							}
						} else if el, ok := v.Type.(*ast.Ellipsis); ok {
							if idt, ok := el.Elt.(*ast.Ident); ok {
								for _, p := range v.Names {
									//fmt.Println(p.Name + " --> " + ti.Name)
									fc.Params = append(fc.Params, GoCodeParam{ParamName: p.Name, TypeName: idt.Name, IsEllipsis: true})
								}
							}
						} else if el, ok := v.Type.(*ast.ArrayType); ok {
							el.End()
						} else if el, ok := v.Type.(*ast.SelectorExpr); ok {
							el.End()
						} else {
							v.Type.(*ast.StarExpr).End()
						}
					}
				}

				if f.Type.Results != nil {
					for _, v := range f.Type.Results.List {
						//fmt.Println(v.Type)
						if n, ok := v.Type.(*ast.Ident); ok {
							//fmt.Println(n.Name)
							fc.Returns = append(fc.Returns, n.Name)
						} else if n, ok := v.Type.(*ast.StarExpr); ok {
							if x, ok := n.X.(*ast.Ident); ok {
								fc.Returns = append(fc.Returns, x.Name)
							} else if x, ok := n.X.(*ast.SelectorExpr); ok {
								fc.Returns = append(fc.Returns, x.Sel.Name)
							}
						}
					}
				}

				if !foundFunction {
					continue
				}

				if f.Recv != nil &&
					len(f.Recv.List) > 0 {
					//fmt.Println(f.Type.Params.List)
					exprIdent, okIdent := f.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident)
					if okIdent {
						if st, ok := code[exprIdent.Name]; ok {
							st.Functions = append(st.Functions, *fc)
							//fmt.Println(fc)
						}
					}
				} else if f.Recv == nil {
					////
					if st, ok := code[fc.LuaName]; ok {
						st.StaticFunctions = append(st.StaticFunctions, *fc)
					}
				}
			}
		} else if f, ok := v.(*ast.GenDecl); ok {
			if f.Doc != nil {
				goCode := &GoCodeStruct{}
				for _, v := range f.Doc.List {
					index := strings.Index(v.Text, "@lua")
					if index > 0 {
						if len(f.Specs) > 0 {
							if ts, ok := f.Specs[0].(*ast.TypeSpec); ok {
								//fmt.Println(ts.Name)
								if ft, ok := ts.Type.(*ast.StructType); ok {
									//fmt.Println(st.Fields.List)
									goCode.foundStruct = true
									goCode.GoName = ts.Name.Name
									if Cfg.InPackageName != "" {
										goCode.GoNameWithPack = Cfg.InPackageName + "." + ts.Name.Name
									} else {
										goCode.GoNameWithPack = ts.Name.Name
									}
									if ft.Fields != nil {
										for _, field := range ft.Fields.List {
											if field.Tag != nil && strings.Index(field.Tag.Value, "lua:") > 0 {
												for _, fname := range field.Names {
													//判断首字母大写
													vv := []rune(fname.Name)
													if len(vv) >= 1 && vv[0] >= 65 && vv[0] <= 90 {
														fc := &GoCodeFunction{}
														fc.GoName = fname.Name
														fc.LuaName = fname.Name

														if fType, ok := field.Type.(*ast.Ident); ok {
															fc.Returns = append(fc.Returns, fType.Name)
															fc.Params = append(fc.Params, GoCodeParam{
																ParamName: fname.Name,
																TypeName:  fType.Name,
															})
															fc.GetterSetter = true
															fc.GetterSetter_Star = false
															switch(fType.Name) {
															case "uint64", "uint32", "int64", "int32", "int", "float32", "float64":
																fc.GetterSetter_Star = true
															case "string", "bool":
																fc.GetterSetter_Star = true
															}
															goCode.Functions = append(goCode.Functions, *fc)
														} else if fStartType, ok := field.Type.(*ast.StarExpr); ok {
															if fType, ok := fStartType.X.(*ast.Ident); ok {
																fc.Returns = append(fc.Returns, fType.Name)
																fc.Params = append(fc.Params, GoCodeParam{
																	ParamName: fname.Name,
																	TypeName:  fType.Name,
																})
																fc.GetterSetter = true
																fc.GetterSetter_Star = true
																goCode.Functions = append(goCode.Functions, *fc)
															}
														}
													}
												}
											}
										}
									}
								}
							}

							//goCode.GoName = f.Specs
							ss := strings.Split(v.Text[index+4:], ":")
							if len(ss) > 1 {
								goCode.LuaName = ss[1]
							} else {
								goCode.LuaName = goCode.GoName
							}
						}
					}
				}

				if goCode.foundStruct {
					code[goCode.GoName] = goCode
				}
			}
		}
	}
	fmt.Println("----------\n\n")

	luaBind := `/*
自动生成的Lua绑定, 对应GO文件#GoSrcFile#
*/
`
	packageName := Cfg.OutPackageName
	if packageName == "" {
		packageName = file.Name.Name
	}
	luaBind += "package " + packageName

	luaBind += `
	import "github.com/yuin/gopher-lua"
	#import_code#

	/*#bind_list#*/

	#init_code#
`
	import_code := ""
	has_import := make(map[string]bool)
	for _, v := range Cfg.ImportString {
		if v != "" && has_import[v] == false {
			has_import[v] = true
			import_code += `import "` + v + `"
`
		}
	}
	luaBind = strings.Replace(luaBind, "#import_code#", import_code, -1)

	for _, v := range code {
		luaBind += v.BuildLuaBind() + "\n"
	}
	for _, c := range callLua {
		//fmt.Println(c.BuildLuaBind(nil))
		luaBind += c.BuildLuaBind(nil) + "\n"
	}

	init_code := ""
	if Cfg.AutoInit {
		init_code = "func init(){\n"
	} else {
		init_code = "//调用这个函数注册Go对象到Lua\nfunc #RegFuncName#(L *lua.LState){\n"
	}
	for _, f := range Cfg.initFunction {
		init_code += f + "\n"
	}
	init_code += "}"
	luaBind = strings.Replace(luaBind, "#init_code#", init_code, -1)

	luaBind = strings.Replace(luaBind, "#bind_list#", Cfg.bindList, -1)
	//fmt.Println(luaBind)
	return luaBind
}
