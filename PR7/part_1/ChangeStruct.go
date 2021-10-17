package part_1

import (
	"errors"
	"reflect"
)

/*
	Написать функцию, которая принимает на вход структуру in (struct или кастомную struct)
	и values map[string]interface{} (key - название поля структуры, которому нужно присвоить
	value этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
	пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции
	тесты (чем больше, тем лучше - зачтется в плюс).
*/

func ChangeStruct(in interface{}, values map[string]interface{}) error {
	if in == nil || values == nil {
		return nil
	}

	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("'in' not a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		valueFromMap := values[typeField.Name]
		if valueFromMap == nil {
			continue
		}
		if typeField.Type.Kind() == reflect.Struct {
			v, ok := valueFromMap.(map[string]interface{})
			_ = v
			if !ok {
				continue
			}
			err := ChangeStruct(val.Field(i).Addr().Interface(), v)
			if err != nil {
				return err
			}
			continue
		}
		if reflect.TypeOf(valueFromMap).Kind() != typeField.Type.Kind() {
			return errors.New("value of type " + reflect.TypeOf(valueFromMap).Kind().String() + " is not assignable to type " + typeField.Type.Kind().String())
		}
		val.Field(i).Set(reflect.ValueOf(valueFromMap))
	}
	return nil
}