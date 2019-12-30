package golang_helpers

import (
	"io"
	"log"
	"reflect"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Contains(slice interface{}, elem interface{}) bool {
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		if sliceValue.Index(i).Interface() == reflect.ValueOf(elem).Interface() {
			return true
		}
	}
	return false
}
