package config

import (
	"os/user"
	"sync"
)

var u *user.User
var mu sync.Mutex

// User returns information about the current user. Cached
func User() (user.User, error) {
	mu.Lock()
	defer mu.Unlock()

	if u == nil {
		p, err := user.Current()
		if err != nil {
			return user.User{}, err
		}
		u = p
	}

	return *u, nil
}
