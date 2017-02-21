package logger

var redirectStdErr bool

func RedirectStdErr() {
	redirectStdErr = true
}
