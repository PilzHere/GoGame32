package player

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/google/uuid"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/utils"
)

type Player struct { // object level
	entity.Entity
	assMan      *utils.AssetsManager
	window      *pixelgl.Window
	Position    pixel.Vec
	sprite      *pixel.Sprite
	moveSpeed   float64
	rect        pixel.Rect
	spriteRect1 *pixel.Sprite
	spriteRect2 *pixel.Sprite
	rectPos     pixel.Vec
	rect1       pixel.Rect
	rect2       pixel.Rect
}

var ( // vars and constants are on class level
//loadedTexture bool
//pic           pixel.Picture
)

func NewPlayer(assMan *utils.AssetsManager, window *pixelgl.Window, x float64, y float64) *Player {
	player := &Player{
		Entity: entity.Entity{
			Id:        uuid.Must(uuid.NewRandom()),
			ToDestroy: false,
		},
		assMan:    assMan,
		window:    window,
		Position:  pixel.V(x, y),
		moveSpeed: 100.0,
		rect: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},
		rect1: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},
		rect2: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},
		rectPos: pixel.V(x, y),
	}

	/*if !loadedTexture {
		pic, _ = utils.LoadPicture("assets/block.png")
		loadedTexture = true
	}*/

	tileBounds := pixel.R(0, 0, 16, 16)
	player.sprite = pixel.NewSprite(player.assMan.PicTilesheet01, tileBounds)

	tileBounds1 := pixel.R(128-16, 0, 128+16, 16)
	player.spriteRect1 = pixel.NewSprite(player.assMan.PicTilesheet01, tileBounds1)

	tileBounds2 := pixel.R(128-16, 128-16, 128, 128)
	player.spriteRect2 = pixel.NewSprite(player.assMan.PicTilesheet01, tileBounds2)

	//block.moveSpeed = 100.0 // can also be set here.

	return player
}

func (p *Player) Tick(dt *float64) { // (bl *Block) object level

}

func (p *Player) Render(spriteBatch *pixel.Batch, dt *float64) {
	if p.window.Pressed(pixelgl.KeyW) {
		p.Position.Y += p.moveSpeed * *dt

		p.rectPos.Y += p.moveSpeed * *dt
		//p.rect.Moved(pixel.V(0, p.moveSpeed * *dt))
	}

	if p.window.Pressed(pixelgl.KeyS) {
		p.Position.Y -= p.moveSpeed * *dt

		p.rectPos.Y -= p.moveSpeed * *dt
		//p.rect.Moved(pixel.V(0, -p.moveSpeed * *dt))
	}

	if p.window.Pressed(pixelgl.KeyA) {
		p.Position.X -= p.moveSpeed * *dt

		p.rectPos.X -= p.moveSpeed * *dt
		//p.rect.Moved(pixel.V(-p.moveSpeed * *dt, 0))
	}

	if p.window.Pressed(pixelgl.KeyD) {
		p.Position.X += p.moveSpeed * *dt

		p.rectPos.X += p.moveSpeed * *dt
		//p.rect.Moved(pixel.V(p.moveSpeed * *dt, 0))
	}

	p.rect1.Min.X = p.rectPos.X - 8
	p.rect1.Min.Y = p.rectPos.Y - 8
	p.rect1.Max.X = p.rectPos.X + 8
	p.rect1.Max.Y = p.rectPos.Y + 8

	fmt.Println("playerPos: ", p.Position)
	fmt.Println("playerRect: ", p.rectPos)
	fmt.Println("rect1Rect: ", p.rect1)
	fmt.Println("rect2Rect: ", p.rect2)

	//spr2Pos := pixel.Vec{64, 64}

	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center())) // Center sprite in current view matrix.
	//p.sprite.Draw(p.window, pixel.IM.Moved(p.Position)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(pos)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center()).Moved(pos)) // Center sprite in current view matrix.

	if p.rect1.Intersects(p.rect2) {
		fmt.Println("INTERSECTION")
	}

	p.sprite.Draw(spriteBatch, pixel.IM.Moved(p.Position))
	p.spriteRect1.Draw(spriteBatch, pixel.IM.Moved(p.rectPos))

<<<<<<< HEAD
	p.spriteRect2.Draw(spriteBatch, pixel.IM.Moved(pixel.V(32 - 8, 32)))
=======
	p.spriteRect2.Draw(spriteBatch, pixel.IM.Moved(pixel.V(32-8, 32)))
>>>>>>> 0d061366faa80267ab79a31e2e4f355d8886141f
}

//Move man idk
/*func (bl *Block) Move(vec *pixel.Vec, dt *float64) {
	vec.Add(pixel.Vec{Y: bl.moveSpeed * *dt})
}*/

func (p *Player) Destroy() {

}
