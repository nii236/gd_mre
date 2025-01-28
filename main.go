package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

type RotateSprite struct {
	classdb.Extension[RotateSprite, Sprite2D.Instance]
	classdb.Tool
}

func (r *RotateSprite) Ready() {
	fmt.Println("RotateSprite Ready")
}
func (r *RotateSprite) Process(delta Float.X) {
	current := r.Super().AsNode2D().Rotation()
	r.Super().AsNode2D().SetRotation(current + Float.X(math.Pi*delta))
}

func main() {
	classdb.Register[RotateSprite]()
	startup.Loader()
	startup.Engine()
}
