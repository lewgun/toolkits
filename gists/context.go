package gists

//http://arslan.io/ten-useful-techniques-in-go
func withLockContext(fn func()) {
	mu.Lock
	defer mu.Unlock()

	fn()
}

func foo() {
	withLockContext(func() {
		// foo related stuff
	})
}

func bar() {
	withLockContext(func() {
		// bar related stuff
	})
}
