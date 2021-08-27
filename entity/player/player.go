package player

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/google/uuid"
	"github.com/pilzhere/GoGame32/collider"
	"github.com/pilzhere/GoGame32/collision"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/gameMath"
	"github.com/pilzhere/GoGame32/screen"
	"github.com/pilzhere/GoGame32/utils"
	"math"
)

type Player struct { // object level
	entity.Entity
	assMan    *utils.AssetsManager
	window    *pixelgl.Window
	Position  pixel.Vec
	Velocity  pixel.Vec
	sprite    *pixel.Sprite
	moveSpeed float64
	//rect1       pixel.Rect
	spriteRect1 *pixel.Sprite
	spriteRect2 *pixel.Sprite
	//rectPos     pixel.Vec
	collider.Collider
	//rect2       pixel.Rect
	rects                       []collider.Collider
	previousPosition            pixel.Vec
	interpolatedPosition        pixel.Vec
	accumulator                 float64
	screen                      *screen.Screener
	justJumped                  bool
	isGrounded                  bool
	tileSpawnBounds             pixel.Rect
	tileMovingLeft01Bounds      pixel.Rect
	tileMovingRightIdle01Bounds pixel.Rect
	tileMovingRight01Bounds     pixel.Rect
	tileMovingRight02Bounds     pixel.Rect
	tileMovingRight03Bounds     pixel.Rect
	tileMovingRight04Bounds     pixel.Rect
	tileMovingRight05Bounds     pixel.Rect
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
		Velocity:  pixel.V(0, 0),
		moveSpeed: 10.0, // Keep above 10 else player get stuck on sides of block under sometimes when moving left...
		/*rect1: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},*/
		/*rect1: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},*/
		/*rect2: pixel.Rect{
			Min: pixel.Vec{X: x - 8, Y: y - 8},
			Max: pixel.Vec{X: x + 8, Y: y + 8},
		},*/
		//rectPos: pixel.V(x, y),
		Collider: collider.Collider{
			CatBits:  collision.PlayerBit,
			MaskBits: collision.PlayerMask,
			Rect: pixel.Rect{
				Min: pixel.Vec{X: x - 4, Y: y - 4},
				Max: pixel.Vec{X: x + 6, Y: y + 6},
			},
			Position: pixel.V(x, y),
		},
		rects:                []collider.Collider{},
		interpolatedPosition: pixel.V(x, y),
		previousPosition:     pixel.V(x, y),
	}

	player.Collider.ParentInterface = player

	/*if !loadedTexture {
		pic, _ = utils.LoadPicture("assets/block.png")
		loadedTexture = true
	}*/

	player.tileSpawnBounds = pixel.R(0, 0, 16, 16)
	player.tileMovingLeft01Bounds = pixel.R(16, 0, 32, 16)
	player.tileMovingRightIdle01Bounds = pixel.R(16, 16, 32, 32)
	player.tileMovingRight01Bounds = pixel.R(32, 16, 32+16, 32)
	player.tileMovingRight02Bounds = pixel.R(48, 16, 48+16, 32)
	player.tileMovingRight03Bounds = pixel.R(64, 16, 64+16, 32)
	player.tileMovingRight04Bounds = pixel.R(80, 16, 80+16, 32)
	player.tileMovingRight05Bounds = pixel.R(96, 16, 96+16, 32)

	player.sprite = pixel.NewSprite(player.assMan.PicTilesheet01, player.tileSpawnBounds)

	//tileBounds1 := pixel.R(128-16, 0, 128+16, 16)
	//player.spriteRect1 = pixel.NewSprite(player.assMan.PicTilesheet01, tileBounds1)

	//tileBounds2 := pixel.R(128-16, 128-16, 128, 128)
	//player.spriteRect2 = pixel.NewSprite(player.assMan.PicTilesheet01, tileBounds2)

	//block.moveSpeed = 100.0 // can also be set here.

	player.Collider.Position = player.Position
	player.Collider.Rect.Min.X = player.Collider.Position.X - 4
	player.Collider.Rect.Min.Y = player.Collider.Position.Y - 6
	player.Collider.Rect.Max.X = player.Collider.Position.X + 4
	player.Collider.Rect.Max.Y = player.Collider.Position.Y + 6

	return player
}

func (p *Player) GetColliders() []collider.Collider {
	// DO WE EVEN NEED RECTS IN STRUCT?

	//var r []pixel.Rect
	var rects = append(p.rects, p.Collider)
	return rects
}

func (p *Player) HandleInput(dt float64) {
	p.justJumped = false
	//p.previousPosition = p.Position

	if p.window.Pressed(pixelgl.KeyW) {
		//p.Position.Y += p.moveSpeed * *dt
	}

	if p.window.Pressed(pixelgl.KeyS) {
		//p.Position.Y -= p.moveSpeed * *dt
	}

	if p.window.Pressed(pixelgl.KeyA) {
		p.Velocity.X -= p.moveSpeed * dt

		p.sprite.Set(p.assMan.PicTilesheet01, p.tileMovingLeft01Bounds)
	}

	if p.window.Pressed(pixelgl.KeyD) {
		p.Velocity.X += p.moveSpeed * dt

		p.sprite.Set(p.assMan.PicTilesheet01, p.tileMovingRightIdle01Bounds)
	}

	if p.window.JustPressed(pixelgl.KeySpace) {
		if p.isGrounded {
			p.justJumped = true
			p.Velocity.Y = 2.5 // impulse
		}
	}

	deadZoneX := 0.1

	if p.Velocity.X < 0 {
		if p.Velocity.X < -deadZoneX {
			p.Velocity.X += p.moveSpeed / 2 * dt
		} else {
			p.Velocity.X = 0
		}
	} else if p.Velocity.X > 0 {
		if p.Velocity.X > deadZoneX {
			p.Velocity.X -= p.moveSpeed / 2 * dt
		} else {
			p.Velocity.X = 0
		}
	}

	gravityY := 3.0

	p.Velocity.Y -= gravityY * 3 * dt

	maxVelocityX := 1.0
	maxVelocityY := 6.0

	if p.Velocity.X > maxVelocityX {
		p.Velocity.X = maxVelocityX
	} else if p.Velocity.X < -maxVelocityX {
		p.Velocity.X = -maxVelocityX
	}

	if p.Velocity.Y > maxVelocityY {
		p.Velocity.Y = maxVelocityY
	} else if p.Velocity.Y < -maxVelocityY {
		p.Velocity.Y = -maxVelocityY
	}

	p.Position = p.Position.Add(p.Velocity)

	if p.window.JustPressed(pixelgl.KeyRight) {
		fmt.Println("Right")

	}

	p.isGrounded = false
}

func (p *Player) Tick(dt float64) { // (bl *Block) object level
	fmt.Println("player tick")
	fmt.Println(p.Velocity.Y)
	p.Collider.Position = p.Position
	p.Collider.Rect.Min.X = p.Collider.Position.X - 4
	p.Collider.Rect.Min.Y = p.Collider.Position.Y - 6
	p.Collider.Rect.Max.X = p.Collider.Position.X + 4
	p.Collider.Rect.Max.Y = p.Collider.Position.Y + 6
}

func (p *Player) OnCollision(otherCollider collider.Collider) {
	//fmt.Println(p.Collider.CatBits, "collided with ", otherCollider.CatBits)
	if otherCollider.CatBits == collision.TerrainBit {
		//fmt.Println("player collides with terrain")

		if !p.justJumped {

			var intersection = collision.GetIntersectionDepth(p.Collider, otherCollider)

			if math.Abs(intersection.X) < math.Abs(intersection.Y) {
				thisColliderWidth := math.Abs(p.Collider.Rect.Max.X - p.Collider.Rect.Min.X)

				if gameMath.Sign(intersection.X) < 0 {
					fmt.Println("Collision Player RIGHT")
					targetLeft := otherCollider.Rect.Min.X
					p.Position.X = targetLeft - thisColliderWidth/2

					//hitWallRight = true
				} else {
					fmt.Println("Collision Player LEFT")
					targetRight := otherCollider.Rect.Max.X
					p.Position.X = targetRight + thisColliderWidth/2

					//hitWallLeft = true
				}

				p.Velocity.X = 0

			} else if math.Abs(intersection.X) > math.Abs(intersection.Y) {
				thisColliderHeight := math.Abs(p.Collider.Rect.Max.Y - p.Collider.Rect.Min.Y)

				if gameMath.Sign(intersection.Y) < 0 {
					fmt.Println("Collision Player BOTTOM")
					targetTop := otherCollider.Rect.Min.Y
					p.Position.Y = targetTop - thisColliderHeight/2

					//hitWallBottom = true
				} else {
					fmt.Println("Collision Player TOP")
					targetBottom := otherCollider.Rect.Max.Y
					p.Position.Y = targetBottom + thisColliderHeight/2

					//hitWallTop = true
				}

				p.Velocity.Y = 0

				p.isGrounded = true
			}

			p.Collider.Position = p.Position
			p.Collider.Rect.Min.X = p.Collider.Position.X - 4
			p.Collider.Rect.Min.Y = p.Collider.Position.Y - 6
			p.Collider.Rect.Max.X = p.Collider.Position.X + 4
			p.Collider.Rect.Max.Y = p.Collider.Position.Y + 6
		}
	}
}

func (p *Player) Render(spriteBatch *pixel.Batch, dt float64) {
	fmt.Println("posY: ", p.Position.Y)

	/*p.rect1.Min.X = p.rectPos.X - 8
	p.rect1.Min.Y = p.rectPos.Y - 8
	p.rect1.Max.X = p.rectPos.X + 8
	p.rect1.Max.Y = p.rectPos.Y + 8*/

	//fmt.Println("playerPos: ", p.Position)
	//fmt.Println("playerRect: ", p.rectPos)
	//fmt.Println("rect1Rect: ", p.rect1)
	//fmt.Println("rect2Rect: ", p.rect2)

	//fmt.Println("colliderXY: ", p.Collider.Rect.Min.X, p.Collider.Rect.Min.Y)

	//spr2Pos := pixel.Vec{64, 64}

	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center())) // Center sprite in current view matrix.
	//p.sprite.Draw(p.window, pixel.IM.Moved(p.Position)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(pos)) // Center sprite in current view matrix.
	//sprite.Draw(window, pixel.IM.Moved(window.Canvas().Bounds().Center()).Moved(pos)) // Center sprite in current view matrix.

	/*if p.rect1.Intersects(p.rect2) {
		fmt.Println("INTERSECTION")
	}*/

	p.sprite.Draw(spriteBatch, pixel.IM.Moved(p.Position.Add(pixel.V(0, 2))))
	//p.spriteRect1.Draw(spriteBatch, pixel.IM.Moved(p.rectPos))

	//p.spriteRect2.Draw(spriteBatch, pixel.IM.Moved(p.rect2.Center().Sub(pixel.Vec{8, 0})))
}

//Move man idk
/*func (bl *Block) Move(vec *pixel.Vec, dt *float64) {
	vec.Add(pixel.Vec{Y: bl.moveSpeed * *dt})
}*/

func (p *Player) Destroy() {

}
