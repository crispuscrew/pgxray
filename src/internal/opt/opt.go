// Package opt provides a simple wrapper type Opt[T] 
// that can be used to represent optional values of any type T.
package opt

import (
	"fmt"
	"strconv"
)

type Opt[T any] struct {
	Value 	T
	Set  	bool
}

func Set[T any](val T) Opt[T] {
	return Opt[T]{
		Value: val,
		Set:   true,
	}
}

func (opt Opt[T]) IsSet() 	bool 	{ return opt.Set 	}
func (opt Opt[T]) Get() 	T 		{ return opt.Value 	}

// UnmarshalTOML implements the toml.Unmarshaler interface, allowing Opt[T] to be used directly in TOML decoding
func (option *Opt[T]) UnmarshalText(data []byte) error {
	var zero T
	switch any(zero).(type) {
	case string:
		option.Value = any(string(data)).(T)
	case int:
		n, err := strconv.Atoi(string(data))
		if err != nil {
			return fmt.Errorf("cannot parse %q as int: %w", data, err)
		}
		option.Value = any(n).(T)
	default:
		return fmt.Errorf("opt.Opt: unsupported type %T", zero)
	}
	option.Set = true
	return nil
}