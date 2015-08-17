package gists


func ExpectErrorContains(t *testing.T, err error , substr, msg string) {
	errorContains((*testing.T).Errorf, t, err, substr, msg)
}

func AssertErrorContains(t *testing.T, err error, substr, msg string) {
	errorContains((*testing.T).Fatalf, t, err, substr, msg)
}

func errorContains(f func(*testing.T, string, ...interface{}), t *testing.T, err error, substr, msg string) {
	if err == nil {
		f(t, "%s: got nil error; expected error containing %q", msg, substr)
		return
	}
	if !strings.Contains(err.Error(), substr) {
		f(t, "%s: expected error containing %q; got instead error %q", msg, substr, err.Error())
	}
}