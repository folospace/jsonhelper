package jsonhelper

import (
    "encoding/json"
    "reflect"
)

func FillDefaultToNil(origin interface{}) interface{} {
    originVal := reflect.ValueOf(origin)

    switch originVal.Kind() {
    case reflect.Ptr:
        //fill ptr with default
        if originVal.IsNil() {
            originElemType := reflect.TypeOf(origin).Elem()
            if originElemType.Kind() == reflect.Ptr {
                return origin
            }
            if originElemType.Kind() == reflect.Struct {
                if _, ok := origin.(json.Marshaler); ok == false {
                    return origin
                }
            }
            ret := reflect.New(originElemType)
            ret.Elem().Set(reflect.Indirect(reflect.ValueOf(FillDefaultToNil(ret.Elem().Interface()))))
            return ret.Interface()
        } else {
            originVal.Elem().Set(reflect.Indirect(reflect.ValueOf(FillDefaultToNil(originVal.Elem().Interface()))))
            return origin
        }
    case reflect.Struct:
        //fill struct with default
        if _, ok := origin.(json.Marshaler); ok {
            return origin
        }
        retPtr := reflect.New(reflect.TypeOf(origin))
        ret := retPtr.Elem()

        for i := 0; i < originVal.NumField(); i++ {
            retElement := ret.Field(i)
            originElement := originVal.Field(i)
            if retElement.CanSet() == false {
                continue
            }
            if originElement.Kind() == reflect.Ptr {
                retElement.Set(reflect.ValueOf(FillDefaultToNil(originElement.Interface())))
            } else {
                retElement.Set(reflect.Indirect(reflect.ValueOf(FillDefaultToNil(originElement.Interface()))))
            }
        }
        return retPtr.Elem().Interface()

    case reflect.Map:
        //fill map with default
        ret := reflect.MakeMap(reflect.TypeOf(origin))
        rangeOverMap := originVal.MapRange()
        for rangeOverMap.Next() {
            k := rangeOverMap.Key()
            v := rangeOverMap.Value()
            if v.Kind() == reflect.Ptr {
                ret.SetMapIndex(k, reflect.ValueOf(FillDefaultToNil(v.Interface())))
            } else {
                ret.SetMapIndex(k, reflect.Indirect(reflect.ValueOf(FillDefaultToNil(v.Interface()))))
            }
        }
        return ret.Interface()

    case reflect.Slice:
        //fill slice with default
        ret := reflect.MakeSlice(originVal.Type(), 0, originVal.Len())
        if !originVal.IsZero() {
            for j := 0; j < originVal.Len(); j++ {
                temp := originVal.Index(j)
                var retElement reflect.Value
                if temp.Kind() == reflect.Ptr {
                    retElement = reflect.ValueOf(FillDefaultToNil(temp.Interface()))
                } else {
                    retElement = reflect.Indirect(reflect.ValueOf(FillDefaultToNil(temp.Interface())))
                }
                ret = reflect.Append(ret, retElement)
            }
        }
        return ret.Interface()
    default:
        return origin
    }
}