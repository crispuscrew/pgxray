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

func IsSet[T any](opt Opt[T]) bool {return opt.Set}
func Get[T any](opt Opt[T]) T {return opt.Value}

func (origin Opt[T]) Override(override Opt[T]) Opt[T] {
	if override.Set {
		return override
	}
	return origin
}

func (option *Opt[T]) UnmarshalTOML(fn func(any) error) error {
	var val T
	if err := fn(&val); err != nil {
		return err
	}
	option.Value = val
	option.Set = true
	return nil
}