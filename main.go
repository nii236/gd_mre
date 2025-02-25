package main

import (
	"fmt"
	"log/slog"

	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/startup"
)

type MRE struct {
	classdb.Extension[MRE, Node.Instance] `gd:"MinimalExample"`
	Player                                PackedScene.Instance
}

func (m *MRE) AsNode() Node.Instance { return m.AsNode() }

func (m *MRE) Ready() {
	player := m.Player.Instantiate()

	fmt.Println("Player:", player)
	char, ok := classdb.As[CharacterBody2D.Instance](Node.Instance(player))
	if !ok {
		slog.Warn("could not get CharacterBody2D from player")
		return
	}
	fmt.Println("char", char)
	fmt.Println("char.AsNode()", char.AsNode())

	// If you uncomment the following line, it will panic when running in the browser
	fmt.Println("char.AsNode().Name()", char.AsNode().Name())
}

func main() {
	classdb.Register[MRE]()
	startup.LoadingScene()
	slog.Info("main ready")
	startup.Scene()
}
