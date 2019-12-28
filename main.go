package main

import (
	"goshorturl/app"
)

func main() {
	a := app.App{}
	a.Initialize()
	a.Run(":8000")
}
