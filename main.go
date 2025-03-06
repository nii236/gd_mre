package main

import (
	"fmt"
	"log/slog"

	"graphics.gd/classdb"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/RayCast3D"
	"graphics.gd/classdb/Timer"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

type MRESpawner struct {
	classdb.Extension[MRESpawner, Node.Instance] `gd:"MRESpawner"`
	Cooldown                                     Timer.Instance
	BulletScene                                  PackedScene.Instance
}

func (b *MRESpawner) Ready() {
	fmt.Println("MRESpawner Ready")
	b.Cooldown.SetWaitTime(1)
	b.Cooldown.SetOneShot(true)
}

func (b *MRESpawner) Process(delta Float.X) {
	if !b.Cooldown.IsStopped() {
		return
	}
	b.Cooldown.Start()
	b.Fire()
}

func (b *MRESpawner) Fire() {
	scene := b.BulletScene.Instantiate()
	b.Super().AsNode().AddChild(scene)
}

type Bullet struct {
	classdb.Extension[Bullet, Node.Instance] `gd:"MREBullet"`
	Mesh                                     MeshInstance3D.Instance
	Ray                                      RayCast3D.Instance
	Speed                                    Float.X
	DeathTimer                               Timer.Instance
}

func (m *Bullet) Ready() {
	fmt.Println("bullet spawned")
	m.DeathTimer.SetAutostart(true)
	m.DeathTimer.SetWaitTime(1)
	m.DeathTimer.SetOneShot(true)
}

func (m *Bullet) Process(delta Float.X) {
	if m.DeathTimer.IsStopped() {
		fmt.Println("timer stopped, remove bullet")
		m.Super().AsNode().QueueFree()
	}
}

func main() {
	classdb.Register[MRESpawner]()
	classdb.Register[Bullet]()
	startup.LoadingScene()
	slog.Info("main ready")
	startup.Scene()
}
