package registry

import (
	"testing"
)

func TestRegister(t *testing.T) {
	i := &Instance{
		Env:     "test",
		AppID:   "provider",
		Address: "http://192.168.1.211:8080",
		Status:  "ok",
	}

	Register(i)
}
