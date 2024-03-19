package goreact

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type GoReact struct {
	Data map[string]map[string]string
}

func (gr *GoReact) Init() {
	gr.Data = make(map[string]map[string]string)
}
func Invoke(obj any, name string, args ...any) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	return reflect.ValueOf(obj).MethodByName(name).Call(inputs)
}

func (gr *GoReact) getCache(obj string, method string) string {
	return gr.Data[obj][method]
}

func (gr *GoReact) checkCache(obj string, method string) bool {
	if _, ok := gr.Data[obj]; ok {
		if _, ok1 := gr.Data[obj][method]; ok1 {
			return true
		}
	}
	return false
}

func (gr *GoReact) cache(obj string, method string, data string) string {
	if gr.Data[obj] == nil {
		gr.Data[obj] = make(map[string]string)
	}

	gr.Data[obj][method] = data
	return data
}

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (gr *GoReact) HandleRoute(ctx *gin.Context, obj any, method string) {
	//	cname := getType(obj)
	res := ""
	// if gr.checkCache(cname, method) {
	// 	res = gr.getCache(cname, method)

	// } else {
	rf := Invoke(obj, method, ctx)
	res = rf[0].Interface().(string)
	//gr.cache(cname, method, res)
	//}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(res))
}
