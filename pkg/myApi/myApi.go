package myApi

import (
	"context"
	"fmt"
	"reflect"
)

// 1. POST Method
// 2. JSON Body
// 3. JSON Body must have (module action)

// func must have two in elem and one or tow out
// the first elem in must be the context.Context impl,the second elem in must be struct or pointer
// the first elem out must be struct or pointer, the seconde elem out must be err.error

type (
	MustHaveParams struct {
		Module string `json:"Module"`
		Action string `json:"Action"`
	}

	Response struct {
		Code    string      `json:"Code"`
		Message string      `json:"Message"`
		Data    interface{} `json:"Data"`
	}
)

type ApiItf interface {
	CheckObjIsApiFunc(handler interface{}) error
}

type ApiImpl struct {
}

var _ ApiItf = (*ApiImpl)(nil)

func (a *ApiImpl) CheckObjIsApiFunc(handler interface{}) (err error) {
	// 1. check obj is Func type
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		err = fmt.Errorf("hanlder is not a function")
		return
	}
	// 2. func must have two in elem and one or tow out
	// 2.1 the first elem in must be the context.Context impl,the second elem in must be struct or pointer
	// 2.2 the first elem out must be struct or pointer, the seconde elem out must be err.error
	handlerType := reflect.TypeOf(handler)
	if handlerType.NumIn() != 2 {
		err = fmt.Errorf("func must have two in elem")
		return
	}
	if handlerType.In(0) != reflect.TypeOf((*context.Context)(nil)).Elem() {
		err = fmt.Errorf("the first elem in must be the context.Context impl")
		return
	}
	if handlerType.In(0).Kind() != reflect.Ptr && handlerType.In(0).Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("the second elem in must be struct or pointer")
		return
	}
	switch handlerType.NumOut() {
	case 1:
		if handlerType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			err = fmt.Errorf("the first elem out must be err.error")
			return
		}
	case 2:
		if handlerType.Out(0).Kind() != reflect.Ptr && handlerType.Out(0).Elem().Kind() != reflect.Struct {
			err = fmt.Errorf("the first elem in must be struct or pointer")
			return
		}
		if handlerType.Out(1) != reflect.TypeOf((*error)(nil)).Elem() {
			err = fmt.Errorf("the first elem out must be err.error")
			return
		}
	}
	return nil
}
