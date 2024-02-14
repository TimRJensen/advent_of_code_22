package common

import (
	"fmt"
)

/**
 * Stack
 */
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0, 256)}
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) Pop() (result T, flag bool) {
	if stack.IsEmpty() {
		return *new(T), false
	}
	result = stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return result, true
}

func (stack *Stack[T]) String() string {
	return fmt.Sprintf("%T%v", stack, stack.items)
}

/**
 * utilities
 */
func MapFilter[K comparable, V any](m map[K]V, f func(v V) bool) (result map[K]V) {
	for key, val := range m {
		if f(val) {
			result[key] = val
		}
	}

	return result
}

func MapReduce[K comparable, V any, R any](m map[K]V, f func(v V, r R) R, r R) (result R) {
	result = r

	for _, val := range m {
		result = f(val, result)
	}

	return result
}

func SliceReduce[V any, R any](v []V, f func(v V, r R) R, r R) (result R) {
	result = r

	for _, val := range v {
		result = f(val, result)
	}

	return result
}
