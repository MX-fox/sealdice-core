package dice

// 用户变量相关

import (
	"reflect"
	"sealdice-core/dice/model"
	"strings"
)

func (ctx *MsgContext) LoadPlayerVars() *PlayerVariablesItem {
	if ctx.Player != nil {
		return LoadPlayerVars(ctx.Session, ctx.Player.UserId)
	}
	return nil
}

func (ctx *MsgContext) LoadGroupVars() {
	g := ctx.Group
	if g.ValueMap == nil {
		g.ValueMap = map[string]*VMValue{}
	}
}

func VarSetValueStr(ctx *MsgContext, s string, v string) {
	VarSetValue(ctx, s, &VMValue{VMTypeString, v})
}

func VarSetValueDNDComputed(ctx *MsgContext, s string, val int64, expr string) {
	vd := &VMComputedValueData{
		BaseValue: VMValue{
			VMTypeInt64,
			val,
		},
		Expr: expr,
	}
	VarSetValue(ctx, s, &VMValue{VMTypeComputedValue, vd})
}

func VarSetValueInt64(ctx *MsgContext, s string, v int64) {
	VarSetValue(ctx, s, &VMValue{VMTypeInt64, v})
}

func VarSetValueAuto(ctx *MsgContext, s string, v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(int))})
	case reflect.Int8:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(int8))})
	case reflect.Int16:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(int16))})
	case reflect.Int32:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(int32))})
	case reflect.Int64:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(int64))})
	case reflect.Uint:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint))})
	case reflect.Uint8:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint8))})
	case reflect.Uint16:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint16))})
	case reflect.Uint32:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint32))})
	case reflect.Uint64:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint64))})
	case reflect.Float32:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(uint64))})
	case reflect.Float64:
		VarSetValue(ctx, s, &VMValue{VMTypeInt64, int64(v.(float64))})
	case reflect.String:
		VarSetValue(ctx, s, &VMValue{VMTypeString, v.(string)})
	}
}

func VarSetValue(ctx *MsgContext, s string, v *VMValue) {
	name := ctx.Player.GetValueNameByAlias(s, nil)
	vClone := *v

	// 临时变量
	if strings.HasPrefix(s, "$t") {
		ctx.Player.ValueMapTemp[s] = &vClone
		return
	}

	// 个人变量
	if strings.HasPrefix(s, "$m") {
		if ctx.Session != nil && ctx.Player != nil {
			playerVars := ctx.LoadPlayerVars()
			playerVars.ValueMap[s] = &vClone
		}
		return
	}

	// 群变量
	if ctx.Group != nil && strings.HasPrefix(s, "$g") {
		ctx.LoadGroupVars()
		ctx.Group.ValueMap[s] = &vClone
		return
	}

	ctx.Player.ValueMap[name] = &vClone
}

func VarDelValue(ctx *MsgContext, s string) {
	name := ctx.Player.GetValueNameByAlias(s, nil)

	// 临时变量
	if strings.HasPrefix(s, "$t") {
		delete(ctx.Player.ValueMapTemp, s)
		return
	}

	// 个人变量
	if strings.HasPrefix(s, "$m") {
		if ctx.Session != nil && ctx.Player != nil {
			playerVars := ctx.LoadPlayerVars()
			delete(playerVars.ValueMap, s)
		}
	}

	// 群变量
	if ctx.Group != nil && strings.HasPrefix(s, "$g") {
		g := ctx.Group
		if g.ValueMap == nil {
			g.ValueMap = map[string]*VMValue{}
		}

		delete(ctx.Group.ValueMap, s)
		return
	}

	delete(ctx.Player.ValueMap, name)
}

func VarGetValue(ctx *MsgContext, s string) (*VMValue, bool) {
	name := ctx.Player.GetValueNameByAlias(s, nil)

	// 临时变量
	if strings.HasPrefix(s, "$t") {
		v, exists := ctx.Player.ValueMapTemp[s]
		return v, exists
	}

	// 个人全局变量
	if strings.HasPrefix(s, "$m") {
		if ctx.Session != nil && ctx.Player != nil {
			playerVars := ctx.LoadPlayerVars()
			a, b := playerVars.ValueMap[s]
			return a, b
		}
	}

	// 群变量
	if ctx.Group != nil && strings.HasPrefix(s, "$g") {
		g := ctx.Group
		if g.ValueMap == nil {
			g.ValueMap = map[string]*VMValue{}
		}

		v, exists := ctx.Group.ValueMap[s]
		return v, exists
	}

	// 个人群变量
	if ctx.Player != nil {
		v, e := ctx.Player.ValueMap[name]
		return v, e
	}
	return nil, false
}

func (i *PlayerInfo) GetValueNameByAlias(s string, alias map[string][]string) string {
	name := s

	if alias == nil {
		// 当私聊的时候，i就会是nil
		if i != nil && i.TempValueAlias != nil {
			alias = *i.TempValueAlias
		}
	}

	for k, v := range alias {
		if strings.EqualFold(s, k) {
			break // 名字本身就是确定值，不用修改
		}
		for _, i := range v {
			if strings.EqualFold(s, i) {
				name = k
				break
			}
		}
	}

	return name
}

func (i *PlayerInfo) SetValueInt64(s string, value int64, alias map[string][]string) {
	name := i.GetValueNameByAlias(s, alias)
	VarSetValue(&MsgContext{Player: i}, name, &VMValue{VMTypeInt64, value})
}

func (i *PlayerInfo) GetValueInt64(s string, alias map[string][]string) (int64, bool) {
	var ret int64
	name := i.GetValueNameByAlias(s, alias)
	v, exists := VarGetValue(&MsgContext{Player: i}, name)

	if exists {
		ret = v.Value.(int64)
	}
	return ret, exists
}

func LoadPlayerVars(s *IMSession, id int64) *PlayerVariablesItem {
	if s.PlayerVarsData == nil {
		s.PlayerVarsData = map[int64]*PlayerVariablesItem{}
	}

	if _, exists := s.PlayerVarsData[id]; !exists {
		s.PlayerVarsData[id] = &PlayerVariablesItem{
			Loaded: false,
		}
	}

	vd, _ := s.PlayerVarsData[id]
	if vd.ValueMap == nil {
		vd.ValueMap = map[string]*VMValue{}
	}

	if vd.Loaded == false {
		vd.Loaded = true
		data := model.AttrUserGetAll(s.Parent.DB, id)
		err := JsonValueMapUnmarshal(data, &vd.ValueMap)
		if err != nil {
			s.Parent.Logger.Error(err)
		}
	}

	return vd
}
