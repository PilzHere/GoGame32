package block

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/google/uuid"
	"github.com/pilzhere/GoGame32/collider"
	"github.com/pilzhere/GoGame32/collision"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/utils"
)

type Block struct { // object level
	entity.Entity
	entity.Entityer
	assMan      *utils.AssetsManager
	window      *pixelgl.Window
	Position    pixel.Vec
	sprite      *pixel.Sprite
	theme       int
	textureCell pixel.Rect
	collider.Collider
	rects []collider.Collider
}

func NewBlock(assMan *utils.AssetsManager, window *pixelgl.Window, uMin float64, vMin float64, uMax float64, vMax float64, theme int, x float64, y float64) *Block {
	block := &Block{
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
		/*rect: pixel.Rect{
			Min: pixel.Vec{X: x-8, Y: y-8},
			Max: pixel.Vec{X: x+8, Y: y+8},
		},*/
		Collider: collider.Collider{
			CatBits:  collision.TerrainBit,
			MaskBits: collision.TerrainMask,
			Rect: pixel.Rect{
				Min: pixel.Vec{X: x - 8, Y: y - 8},
				Max: pixel.Vec{X: x + 8, Y: y + 8},
			},
		},
		rects: []collider.Collider{},
	}

	block.Collider.ParentInterface = block

	//fmt.Println("assMan block: ", block.assMan)

	switch theme {
	case 0:
		// Decide what texture to use here.
		break
	}

	block.sprite = pixel.NewSprite(block.assMan.PicTilesheet01, block.textureCell)

	//block.moveSpeed = 100.0 // can also be set here.

	return block
}

func (bl *Block) GetColliders() []collider.Collider {
	// DO WE EVEN NEED RECTS IN STRUCT?

	//var r []pixel.Rect
	var rects = append(bl.rects, bl.Collider)
	return rects
}

func (bl *Block) OnCollision(otherCollider collider.Collider) {
	//fmt.Println(bl.Collider.CatBits, "collided with ", otherCollider.CatBits)
}

func (bl *Block) Tick(dt float64) { // (bl *Block) = object level
	//fmt.Println(bl.Id)
	//bl.Tick2()

	//bl.Entity.

	//bl.Entityer.Tick(dt)
}

func (bl *Block) Render(spriteBatch *pixel.Batch, dt float64) {
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center())) // Center sprite in current view matrix.
	//bl.sprite.Draw(bl.window, pixel.IM.Moved(bl.Position)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(pos)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center()).Moved(pos)) // Center sprite in current view matrix.

	bl.sprite.Draw(spriteBatch, pixel.IM.Moved(bl.Position))
}

func (bl *Block) Destroy() {

}
