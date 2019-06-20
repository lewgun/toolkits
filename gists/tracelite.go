package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"
)

type item struct {
	Tick time.Time
	Mark string
}

const KeyTrace = "trace"

var Root = newTrace(0, "root")

type trace struct {
	parent   *trace
	children []*trace

	depth  int
	prefix string

	id    string
	items []item

	flag bool
}

func newTrace(depth int, id string) *trace {
	t := &trace{
		id:    id,
		depth: depth,
		items: []item{{Mark: "begin", Tick: time.Now()}},
	}
	return t
}

func NewTrace(parent *trace, id string) *trace {
	if parent == nil {
		parent = newTrace(0, "root")
	}

	if id == "" {
		id = parent.id
	}

	t := newTrace(parent.depth+1, id)
	parent.children = append(parent.children, t)
	return t

}

func (t *trace) Tick(mark string) {
	t.items = append(t.items, item{
		Mark: mark,
		Tick: time.Now(),
	})

}

func (t *trace) End() {
	if t.flag {
		return
	}

	t.flag = true
	t.items = append(t.items, item{Mark: "end", Tick: time.Now()})
}

const format = "%s\t%s\t%.2f\t"

func (t *trace) String() string {

	buf := &bytes.Buffer{}
	for i := 0; i < t.depth; i++ {
		buf.WriteString("  ") // 2 spaces
	}
	prefix := buf.String()

	buf.Truncate(0)

	w := tabwriter.NewWriter(buf, 0, 0, 2, ' ', 0)
	if len(t.items) > 0 {
		begin := t.items[0]
		end := t.items[len(t.items)-1]
		fmt.Fprintln(w, fmt.Sprintf(format, prefix, "* "+t.id, end.Tick.Sub(begin.Tick).Seconds()))
		//fmt.Fprintln(w, fmt.Sprintf(format, prefix, "· "+t.id, end.Tick.Sub(begin.Tick).Seconds()))

	}

	for i := 1; i < len(t.items)-1; i++ {
		fmt.Fprintln(w, fmt.Sprintf(format, prefix, t.items[i].Mark, t.items[i].Tick.Sub(t.items[i-1].Tick).Seconds()))
	}

	fmt.Fprintln(w, "")

	for i := 0; i < len(t.children); i++ {
		fmt.Fprint(w, t.children[i].String())
	}

	w.Flush()

	return buf.String()

}

func (t *trace) Output() {
	fmt.Println(t)

}

func main() {

	t := NewTrace(Root, "123")
	defer t.Output()
	defer t.End()
	t.Tick("1-1111111")
	time.Sleep(2e9)
	t.Tick("1-2222222")

	t2 := NewTrace(t, "456")
	defer t2.End()
	t2.Tick("2-1111111")
	time.Sleep(3e9)
	t2.Tick("2-2222222")

	t3 := NewTrace(t2, "789")
	defer t3.End()
	t3.Tick("3-1111111")
	time.Sleep(3e9)
	t3.Tick("3-2222222")

	t4 := NewTrace(t, "abc")
	defer t4.End()
	t4.Tick("4-1111111")
	time.Sleep(3e9)
	t4.Tick("4-2222222")

	t5 := NewTrace(t, "abc")
	defer t5.End()
	t5.Tick("5-1111111")
	time.Sleep(3e9)
	t5.Tick("5-2222222")
}
