package entity

import (
	"github.com/faiface/pixel"
	"github.com/pilzhere/GoGame32/collider"
)

type Entityer interface {
	Tick(dt float64)
	Render(spriteBatch *pixel.Batch, dt float64)
	GetColliders() []collider.Collider
	Destroy()
}
