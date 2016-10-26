package LRU

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New(8)
	if l.capacity != 8 {
		t.Errorf("New failed")
	}
}

func TestSetGet(t *testing.T) {
	l := New(8)
	for i := 0; i < 10; i++ {
		l.Set(i, i)
	}
	if _, _, ok := l.Get(0); ok == false {
		t.Errorf("Get error")
	}
}
