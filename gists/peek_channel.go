type peekChanOpLog struct {
	in <-chan *opLog
	v  *opLog
	ok bool
}

func newPeekChanOpLog(in <-chan *opLog) *peekChanOpLog {
	return &peekChanOpLog{in: in}
}
func (p *peekChanOpLog) peek() (*opLog, bool) {
	if !p.ok {
		p.v, p.ok = <-p.in
	}
	return p.v, p.ok
}

func (p *peekChanOpLog) get() (*opLog, bool) {
	if p.ok {
		p.ok = false
		return p.v, true
	}
	v, ok := <-p.in
	return v, ok
}
