package events

import (
	"reflect"
	"testing"
	"time"
)

func TestEvent(t *testing.T) {
	t.Run("Clone", func(t *testing.T) {
		e1 := &Event{
			Message: "Hello World",
			Args:    Args{{"hello", "world"}},
			Time:    time.Now(),
		}
		e2 := e1.Clone()

		if e1 == e2 {
			t.Error("Clone cannot return a value with the same address as the original")
		}

		if !reflect.DeepEqual(e1, e2) {
			t.Errorf("%#v", e2)
		}
	})
}

func TestArgs(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		args := Args{
			{"hello", "world"},
			{"answer", 42},
		}

		if v, ok := args.Get("answer"); !ok {
			t.Error("expected answer but got nothing")
		} else if !reflect.DeepEqual(v, 42) {
			t.Error("expected answer=42 but got", v)
		}

		if v, ok := args.Get("question"); ok {
			t.Error("expected no question but got", v)
		}
	})
}