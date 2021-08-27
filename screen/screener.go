package screen

import (
	"github.com/pilzhere/GoGame32/entity"
)

type Screener interface {
	AddEntity(entity entity.Entityer)
	RemoveEntity(entity entity.Entityer)
	HandleInput(dt float64)
	CheckColliders()
	Tick(dt float64)
	Render(dt float64)
	Destroy()
}
