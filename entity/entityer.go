package entity

import "github.com/faiface/pixel"

type Entityer interface {
	Tick(dt *float64)
	Render(spriteBatch *pixel.Batch, dt *float64)
	Destroy()
}
