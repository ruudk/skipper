package weight

import (
	"bufio"
	"net/http"
	"strings"
	"testing"
)

func TestWeightArgs(t *testing.T) {
	for _, ti := range []struct {
		msg  string
		args []interface{}
		err  bool
	}{{
		"no args",
		nil,
		true,
	}, {
		"too many args",
		[]interface{}{"name", "value"},
		true,
	}, {
		"invalid value",
		[]interface{}{"string"},
		true,
	}, {
		"ok float to int",
		[]interface{}{500.99},
		false,
	}, {
		"ok",
		[]interface{}{500},
		false,
	}} {
		func() {
			p, err := New().Create(ti.args)
			if ti.err && err == nil {
				t.Error(ti.msg, "failed to fail")
			} else if !ti.err && err != nil {
				t.Error(ti.msg, err)
			}

			if err != nil {
				return
			}

			if p == nil {
				t.Error(ti.msg, "failed to create filter")
			}
		}()
	}
}

func TestWeightMatch(t *testing.T) {
	p, err := New().Create([]interface{}{500})
	if err != nil {
		t.Error(err)
		return
	}

	r, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET / HTTP/1.0\n\n")))
	if err != nil {
		t.Error(err)
		return
	}

	if m := p.Match(r); m != true {
		t.Error("failed to match", m)
	}
}
