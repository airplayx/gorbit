package validator

import (
	"errors"
	"reflect"
	"time"
)

const (
	kFuncSuffix = "Validator"
)

var (
	ErrNilObject = errors.New("validator: receiving nil object")
)

func Check(obj interface{}) error {
	var objType = reflect.TypeOf(obj)
	var objValue = reflect.ValueOf(obj)
	var objValueKind = objValue.Kind()

	for {
		if objValueKind == reflect.Ptr && objValue.IsNil() {
			return ErrNilObject
		}
		if objValueKind == reflect.Ptr {
			objValue = objValue.Elem()
			objType = objType.Elem()
			objValueKind = objValue.Kind()
			continue
		}
		break
	}
	return check(objType, objValue, objValue)
}

func check(objType reflect.Type, parent, current reflect.Value) error {
	var numField = objType.NumField()
	for i := 0; i < numField; i++ {
		var fieldStruct = objType.Field(i)
		var fieldValue = current.Field(i)

		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		if fieldValue.Kind() == reflect.Struct && fieldValue.Type() != reflect.TypeOf(time.Time{}) {
			if err := check(fieldValue.Type(), parent, fieldValue); err != nil {
				return err
			}
			continue
		}

		var mName = fieldStruct.Name + kFuncSuffix
		var mValue = methodByName(mName, parent, current)

		if mValue.IsValid() {
			var pValue []reflect.Value
			if fieldValue.IsValid() {
				pValue = []reflect.Value{fieldValue}
			} else {
				pValue = []reflect.Value{reflect.New(fieldStruct.Type).Elem()}
			}
			var rValueList = mValue.Call(pValue)

			if !rValueList[0].IsNil() {
				var err = rValueList[0].Interface().(error)
				return err
			}
		}
	}
	return nil
}

func methodByName(mName string, parent, current reflect.Value) reflect.Value {
	var mValue = parent.MethodByName(mName)
	if mValue.IsValid() == false {
		if parent.CanAddr() {
			mValue = parent.Addr().MethodByName(mName)
		}
	}
	if mValue.IsValid() == false && parent != current {
		return methodByName(mName, current, current)
	}
	return mValue
}
