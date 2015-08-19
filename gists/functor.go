package gists


type statusFunc func() string

func (sf statusFunc) String() string {
	return sf()
}

type status string

func (s status) String() string {
	return string(s)
}


set := func(s fmt.Stringer) {
	...
}

set(statusFunc(func() string {
	return fmt.Sprintf("copying: %d/%d bytes", bytesCopied, sb.Size)
}))

set(status("copied; removing from queue"))