package main

import (
	"lgdisc/controller"
)

func main() {
	r := controller.MakeRoute()
	r.Run(":8080")
}
