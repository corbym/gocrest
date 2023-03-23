package by

import (
	"bufio"
	"io"
	"reflect"
)

func Channelling(actual interface{}) interface{} {
	var selectCase = make([]reflect.SelectCase, 1)
	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(actual)
	_, recv, _ := reflect.Select(selectCase)
	return recv.Interface()
}
func Reading(actual io.Reader, len int) interface{} {
	reader := bufio.NewReader(actual.(io.Reader))
	peek, _ := reader.Peek(len)
	return peek
}
