package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/startup"
)

type Example struct {
	classdb.Extension[Example, Node2D.Instance] `gd:"Example"`
}

func main() {
	startup.Loader()
	classdb.Register[Example]()
	startup.Engine()
}
