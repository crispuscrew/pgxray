// Package opt provides a simple wrapper type Opt[T] 
// that can be used to represent optional values of any type T.
package opt

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
func (option *Opt[T]) UnmarshalTOML(fn func(any) error) error {
	var val T
	if err := fn(&val); err != nil {
		return err
	}
	option.Value = val
	option.Set = true
	return nil
}