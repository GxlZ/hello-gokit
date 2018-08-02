package di

import (
	"go.uber.org/dig"
)

var Di *dig.Container

func init() {
	Di = dig.New()

}
