package server

// A Reloader is an application that can reload
type Reloader interface {
	Reload() error
}
