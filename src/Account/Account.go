package Account

import (
	"fmt"
	"errors"
	"strings"
)

//DummyType is original struct
type DummyType struct {
	name  string
	value int
}

//InitDummy X aaaaa
func InitDummy(name string, value int) *DummyType {
	dumm := DummyType{name: name, value: value}
	return &dumm
}

//CallingValue X blahblah
func (d DummyType) CallingValue() (string, int) {
	return d.name, d.value
}

//String X blaaahhh
func (d DummyType) String() string {
	return fmt.Sprint("This struct value : ")
}
