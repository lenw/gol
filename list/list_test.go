package list

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	l := New()
	l.Add("Hello")
	l.Add("World")

	i := l.Iterator()
	for i.HasNext() {
		fmt.Printf("%+v\n", i.Item())
		i.Next()
	}
}
