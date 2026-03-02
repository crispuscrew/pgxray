package cfg

import (
	"github.com/crispuscrew/pgxray/internal/opt"

	"reflect"
	"os"
	"path/filepath"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
)

type Field struct {
	Value reflect.Value
	Meta  reflect.StructField
}

func ForEachField[T any](val *T, fn func(field Field)) {
	value := reflect.ValueOf(val).Elem()
	valType := reflect.TypeOf(*val)
	for i := range value.NumField() {
		fn(Field{Value: value.Field(i), Meta: valType.Field(i)})
	}
}

// TODO: implement --help info for keybinds config file and CLI overrides
func (kb *Keybind) UnmarshalTOML(fn func(any) error) error {
	var keys []string
	if err := fn(&keys); err != nil {
		return err
	}
	*kb = Keybind{
		key.NewBinding(key.WithKeys(keys...)),
	}
	return nil
}

func resolvePath(path opt.Opt[string], envVar string, defaultPath string) (string, error) {
	if path.IsSet() {
		return resolveHome(path.Get())
	}
	if os.Getenv(envVar) != "" {
		return resolveHome(os.Getenv(envVar))
	}
	return resolveHome(defaultPath)
}

func resolveHome(path string) (string, error) {                                                                                                    
	if strings.HasPrefix(path, "~/") {                                                                                                             
		home, err := os.UserHomeDir()                                                                                                              
		if err != nil {
			return "", fmt.Errorf("resolvePath: %w", err)
		}
		return filepath.Join(home, path[2:]), nil
	}
	return filepath.Abs(path)
}

func merge[T any](base, override T) T {
	ForEachFieldPair := func (a, b *T, fn func(a, b Field)) {                                                                                              
		av := reflect.ValueOf(a).Elem()                                                                                                                
		bv := reflect.ValueOf(b).Elem()                                                                                                                
		t := reflect.TypeOf(*a)
		for i := range av.NumField() {
			fn(
				Field{av.Field(i), t.Field(i)},
				Field{bv.Field(i), t.Field(i)},
			)
		}
	}
	type settable interface {
		IsSet() bool                                                          
	}
	ForEachFieldPair(&base, &override, func(field, overrideField Field) {
		set, ok := overrideField.Value.Interface().(settable)
		if ok && set.IsSet() {
			field.Value.Set(overrideField.Value)
		}
	})
	return base
}