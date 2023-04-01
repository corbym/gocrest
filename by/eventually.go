package by

import (
	"bufio"
	"io"
)

// Channelling channels any channel of type T and returns the value from the channel
func Channelling[T any](actual chan T) T {
	return <-actual
}

// Reading peeks at the value of a reader by reading `len` bytes ahead.
func Reading(actual io.Reader, len int) []byte {
	reader := bufio.NewReader(actual.(io.Reader))
	peek, _ := reader.Peek(len)
	return peek
}

// Calling calls the function passed and returns the value
func Calling[K any, T any](actual func(T) K, value T) K {
	return actual(value)
}
