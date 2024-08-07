package router

import (
	"reflect"
)

type IRouterInterface interface {
}

type IRouter struct {
	HandlerMap map[string]reflect.Value
}

// func AAA(ctx context.Context, * / struct) (*, error)
