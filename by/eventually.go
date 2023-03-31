package by

import (
	"bufio"
	"io"
)

func Channelling[T any](actual chan T) T {
	return <-actual
}
func Reading(actual io.Reader, len int) []byte {
	reader := bufio.NewReader(actual.(io.Reader))
	peek, _ := reader.Peek(len)
	return peek
}

func Calling[K any, T any](actual func(T) K, value T) K {
	return actual(value)
}
