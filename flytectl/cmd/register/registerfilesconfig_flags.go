// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package register

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (RegisterFilesConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (RegisterFilesConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in RegisterFilesConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg RegisterFilesConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("RegisterFilesConfig", pflag.ExitOnError)
	cmdFlags.StringVarP(&(filesConfig.version),fmt.Sprintf("%v%v", prefix, "version"), "v", "v1", "version of the entity to be registered with flyte.")
	cmdFlags.BoolVarP(&(filesConfig.skipOnError), fmt.Sprintf("%v%v", prefix, "skipOnError"), "s", *new(bool), "fail fast when registering files.")
	return cmdFlags
}
