package bar

//go:generate go run github.com/zdandoh/ecs/codegen myecspkg components

import (
	"testing"

	"bar/components"
	ecs "bar/myecspkg"
)

const nPos = 9000
const nPosVel = 1000

func BenchmarkIterArche(b *testing.B) {
	b.StopTimer()

	posEntities := make([]ecs.Entity, nPos)
	for i := 0; i < nPos; i++ {
		e := ecs.NewEntity()
		e.SetPosition(components.Position{X: 1, Y: 2})
		posEntities[i] = e
	}

	posVelEntities := make([]ecs.Entity, nPosVel)
	for i := 0; i < nPosVel; i++ {
		e := ecs.NewEntity()
		e.SetPosition(components.Position{X: 1, Y: 2})
		e.SetVelocity(components.Velocity{X: 1, Y: 2})
		posVelEntities[i] = e
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ecs.Select(func(entity ecs.Entity, pos *components.Position, vel *components.Velocity) {
			pos.X += vel.X
			pos.Y += vel.Y
		})
	}
}
