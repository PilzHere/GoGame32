package block

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/google/uuid"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/utils"
)

type Block struct { // object level
	entity.Entity
	assMan      *utils.AssetsManager
	window      *pixelgl.Window
	Position    pixel.Vec
	sprite      *pixel.Sprite
	theme       int
	textureCell pixel.Rect
}

// Global Vars and Constants are stored in the lifetime of the program.
<<<<<<< HEAD
var (

)

func NewBlock(assMan *utils.AssetsManager, window *pixelgl.Window, uMin float64, vMin float64, uMax float64, vMax float64, theme int, x float64, y float64) *Block {
	block := &Block{
		Entity:   entity.Entity{
			Id:        uuid.Must(uuid.NewRandom()),
			ToDestroy: false,
		},
		assMan:   assMan,
		window:   window,
		Position: pixel.V(x, y),
=======
var ()

func NewBlock(assMan *utils.AssetsManager, window *pixelgl.Window, uMin float64, vMin float64, uMax float64, vMax float64, theme int, x float64, y float64) *Block {
	block := &Block{
		Entity: entity.Entity{
			Id:        uuid.Must(uuid.NewRandom()),
			ToDestroy: false,
		},
		assMan:      assMan,
		window:      window,
		Position:    pixel.V(x, y),
>>>>>>> 0d061366faa80267ab79a31e2e4f355d8886141f
		sprite:      nil,
		theme:       theme,
		textureCell: pixel.Rect{Min: pixel.V(uMin, vMin), Max: pixel.V(uMax, vMax)},
	}

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

func (bl *Block) Tick(dt *float64) { // (bl *Block) = object level
	//fmt.Println(bl.Id)
	//bl.Tick2()
}

func (bl *Block) Render(spriteBatch *pixel.Batch, dt *float64) {
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center())) // Center sprite in current view matrix.
	//bl.sprite.Draw(bl.window, pixel.IM.Moved(bl.Position)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(pos)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center()).Moved(pos)) // Center sprite in current view matrix.

	bl.sprite.Draw(spriteBatch, pixel.IM.Moved(bl.Position))
}

func (bl *Block) Destroy() {

}
