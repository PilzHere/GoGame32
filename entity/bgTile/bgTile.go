package bgTile

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/google/uuid"
	"github.com/pilzhere/GoGame32/collider"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/utils"
)

type BgTile struct {
	entity.Entity
	entity.Entityer
	assMan      *utils.AssetsManager
	window      *pixelgl.Window
	Position    pixel.Vec
	sprite      *pixel.Sprite
	theme       int
	textureCell pixel.Rect
}

func NewBgTile(assMan *utils.AssetsManager, window *pixelgl.Window, uMin float64, vMin float64, uMax float64, vMax float64, theme int, x float64, y float64) *BgTile {
	bgTile := &BgTile{
		Entity: entity.Entity{
			Id:        uuid.Must(uuid.NewRandom()),
			ToDestroy: false,
		},
		Entityer:    entity.Entityer(nil),
		assMan:      assMan,
		window:      window,
		Position:    pixel.V(x, y),
		sprite:      nil,
		theme:       theme,
		textureCell: pixel.Rect{Min: pixel.V(uMin, vMin), Max: pixel.V(uMax, vMax)},
	}

	switch theme {
	case 0:
		// Decide what texture to use here.
		break
	}

	bgTile.sprite = pixel.NewSprite(bgTile.assMan.PicTilesheet01, bgTile.textureCell)

	return bgTile
}

func (bl *BgTile) GetColliders() []collider.Collider {
	return nil
}

func (bl *BgTile) Tick(dt float64) {
	return
}

func (bl *BgTile) Render(spriteBatch *pixel.Batch, dt float64) {
	bl.sprite.Draw(spriteBatch, pixel.IM.Moved(bl.Position))
}

func (bl *BgTile) Destroy() {
	return
}
