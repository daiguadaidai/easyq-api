package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// 将不是默认字段值, 设置为值定的字段值
func ReflectSetNotDefaultFieldValue(prefix string, baseObj, modifyObj interface{}) error {
	v := reflect.ValueOf(baseObj)
	t := reflect.TypeOf(baseObj)

	for i := 0; i < v.NumMethod(); i++ {
		vMethod := v.Method(i)
		tMethod := t.Method(i)
		// 判断方法是否是某个指定开头
		if strings.HasPrefix(tMethod.Name, prefix) && tMethod.Name != prefix {
			filedName := strings.TrimLeft(tMethod.Name, prefix)
			resValues := vMethod.Call(nil)
			if len(resValues) != 1 {
				continue
			}
			// 判断返回的第一个值是否是bool类型
			isDefault, ok := resValues[0].Interface().(bool)
			if !ok {
				continue
			}
			// 是默认值
			if isDefault {
				continue
			}

			// 获取字段值
			value, err := GetInterfaceFieldValue(baseObj, filedName)
			if err != nil {
				return err
			}

			// 将baseObj字段值设置到 modifyObj 中
			err = SetInterfaceFieldValue(modifyObj, filedName, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func SetObjFieldsFromOtherObj(fieldNames []string, fromObj interface{}, toObj interface{}) error {
	for _, fieldName := range fieldNames {
		if err := SetObjFieldFromOtherObj(fieldName, fromObj, toObj); err != nil {
			return err
		}
	}

	return nil
}

func SetObjFieldFromOtherObj(fieldName string, fromObj interface{}, toObj interface{}) error {
	// 获取字段值
	fromFieldValue, err := GetInterfaceFieldValue(fromObj, fieldName)
	if err != nil {
		return err
	}

	// 将baseObj字段值设置到 modifyObj 中
	err = SetInterfaceFieldValue(toObj, fieldName, fromFieldValue)
	if err != nil {
		return err
	}

	return nil
}

func GetInterfaceFieldValue(v interface{}, fieldName string) (interface{}, error) {
	val := reflect.ValueOf(v)
	return reflectGetFieldValue(val, fieldName)
}

func reflectGetFieldValue(val reflect.Value, fieldName string) (interface{}, error) {
	if val.IsValid() && val.CanInterface() {
		typ := val.Type()
		switch typ.Kind() {
		case reflect.Struct:
			return val.FieldByName(fieldName).Interface(), nil
		case reflect.Ptr:
			return reflectGetFieldValue(val.Elem(), fieldName)
		}
	} else {
		return nil, fmt.Errorf("无效的数据结构: %v", val.String())
	}

	return nil, nil
}

func SetInterfaceFieldValue(v interface{}, fieldName string, fieldValue interface{}) error {
	val := reflect.ValueOf(v)
	return reflectSetFieldValue(val, fieldName, fieldValue)
}

func reflectSetFieldValue(val reflect.Value, fieldName string, fieldValue interface{}) error {
	if val.IsValid() && val.CanInterface() {
		typ := val.Type()
		switch typ.Kind() {
		case reflect.Struct:
			reflectSetValue(val, fieldName, fieldValue)
			return nil
		case reflect.Ptr:
			return reflectSetFieldValue(val.Elem(), fieldName, fieldValue)
		}
	} else {
		return fmt.Errorf("无效的数据结构: %v", val.String())
	}

	return nil
}

func reflectSetValue(val reflect.Value, fieldName string, fieldValue interface{}) {
	val.FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
}

func GetStructFieldNames(v interface{}) ([]string, error) {
	val := reflect.ValueOf(v)
	return reflectGetStructFieldNames(val)
}

func reflectGetStructFieldNames(val reflect.Value) ([]string, error) {
	if val.IsValid() && val.CanInterface() {
		typ := val.Type()
		switch typ.Kind() {
		case reflect.Struct:
			fieldNames := make([]string, 0, val.NumField())
			for i := 0; i < val.NumField(); i++ {
				tf := typ.Field(i)
				fieldNames = append(fieldNames, tf.Name)
			}
			return fieldNames, nil
		case reflect.Ptr:
			return reflectGetStructFieldNames(val.Elem())
		}
	} else {
		return nil, fmt.Errorf("无效的数据结构(获取结构体所有字段名): %v", val.String())
	}

	return nil, nil
}

func GetStructFieldNamesWithPrefix(v interface{}, prefix string) ([]string, error) {
	val := reflect.ValueOf(v)
	return reflectGetStructFieldNamesWithPrefix(val, prefix)
}

func reflectGetStructFieldNamesWithPrefix(val reflect.Value, prefix string) ([]string, error) {
	if val.IsValid() && val.CanInterface() {
		typ := val.Type()
		switch typ.Kind() {
		case reflect.Struct:
			fieldNames := make([]string, 0, val.NumField())
			for i := 0; i < val.NumField(); i++ {
				tf := typ.Field(i)
				fieldNames = append(fieldNames, fmt.Sprintf("%v%v", prefix, tf.Name))
			}
			return fieldNames, nil
		case reflect.Ptr:
			return reflectGetStructFieldNamesWithPrefix(val.Elem(), prefix)
		}
	} else {
		return nil, fmt.Errorf("无效的数据结构(获取结构体所有字段名): %v", val.String())
	}

	return nil, nil
}

func InterfaceToStruct(i interface{}, mode interface{}) (interface{}, error) {
	raw, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(mode)
	obj := reflect.New(t.Elem()).Interface()

	if err = json.Unmarshal(raw, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}

	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		// 通过字段名获取字段是否有效
		if dvalue.IsValid() == false {
			continue
		}
		// 判断类型是否一样
		if value.Type().Name() != dvalue.Type().Name() {
			continue
		}

		dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
}
