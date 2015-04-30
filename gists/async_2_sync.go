package gists

func (d *driver) asyncDriver(f func()) {
	d.pendingDriver <- f
	d.wake()
}

func (d *driver) syncDriver(f func()) {
	c := make(chan bool, 1)
	d.asyncDriver(func() { f(); c <- true })
	<-c
}
