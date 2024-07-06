package main

import (
	"time"

	"gopkg.in/yaml.v3"
)

type Optional[T any] struct {
	Value T
	Set   bool
}

func (o Optional[T]) IsSet() bool {
	return o.Set
}

func (o Optional[T]) Get() (T, bool) {
	return o.Value, o.Set
}

func (o *Optional[T]) SetValue(value T) {
	o.Value = value
	o.Set = true
}

func (o *Optional[T]) Clear() {
	var zero T
	o.Value = zero
	o.Set = false
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{Value: value, Set: true}
}

type OptionalTime struct {
	Optional[time.Time]
}

func NewOptionalTime(t time.Time) OptionalTime {
	return OptionalTime{Optional: NewOptional(t)}
}

func (ot OptionalTime) Format(layout string) string {
	if ot.Set {
		return ot.Value.Format(layout)
	}
	return ""
}

func (ot *OptionalTime) UnmarshalYAML(node *yaml.Node) error {
	if node.Kind == yaml.ScalarNode && node.Value == "" {
		ot.Clear()
		return nil
	}
	var t time.Time
	err := node.Decode(&t)
	if err != nil {
		return err
	}
	ot.SetValue(t)
	return nil
}

func (ot OptionalTime) MarshalYAML() (interface{}, error) {
	if !ot.Set {
		return nil, nil
	}
	return ot.Value, nil
}
