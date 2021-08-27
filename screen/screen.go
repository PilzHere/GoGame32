package screen

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pilzhere/GoGame32/collider"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/utils"
)

type Screen struct {
	Entities    []entity.Entityer
	Colliders   []collider.Collider
	Window      *pixelgl.Window
	SpriteBatch *pixel.Batch
	AssMan      *utils.AssetsManager
}
