package playScreen

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
	"github.com/pilzhere/GoGame32/collider"
	"github.com/pilzhere/GoGame32/collision"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/entity/bgTile"
	"github.com/pilzhere/GoGame32/entity/block"
	"github.com/pilzhere/GoGame32/entity/player"
	"github.com/pilzhere/GoGame32/screen"
	"github.com/pilzhere/GoGame32/utils"
	"golang.org/x/image/colornames"
	"sort"
)

type PlayScreen struct {
	screen.Screen
	currentMap   *tiled.Map
	currentTheme int
	*player.Player
}

func NewPlayScreen(spriteBatch *pixel.Batch, assMan *utils.AssetsManager, window *pixelgl.Window) *PlayScreen {
	ps := &PlayScreen{
		Screen: screen.Screen{
			Entities:    []entity.Entityer{},
			Colliders:   []collider.Collider{},
			SpriteBatch: spriteBatch,
			Window:      window,
			AssMan:      assMan,
		},
		Player: nil,
	}

	ps.currentMap, _ = tiled.LoadFromFile("assets/maps/map01.tmx")
	ps.currentTheme = ps.currentMap.Properties.GetInt("theme")
	height := float64(ps.currentMap.Tilesets[0].Image.Height)
	//width := float64(ps.currentMap.Tilesets[0].Image.Width)
	cellSize := 16.0

	//println("assMan: ", ps.AssMan)

	//spriteBatch = pixel.NewBatch(&pixel.TrianglesData{}, assMan.PicTilesheet01)
	//spriteBatch.MakePicture(assMan.PicTilesheet01)

	// background
	for _, e := range ps.currentMap.ObjectGroups[2].Objects {
		switch e.Name {
		case "bgWall01":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 0, 48, 16, 64, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgWall02":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 0, 64, 16, 64+16, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgWallPillar01":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 16, 48, 32, 64, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgWallPillar02":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 16, 64, 32, 64+16, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgWallHole01":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 32, 48, 48, 64, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgWallHole02":
			ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 32, 64, 48, 64+16, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		}
	}

	// frame
	for _, e := range ps.currentMap.ObjectGroups[0].Objects {
		switch e.Name {
		case "bgBlock01":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 0, height, cellSize, height-cellSize, ps.currentTheme, e.X+16+8, e.Y-8))
			//ps.AddEntity(bgTile.NewBgTile(ps.AssMan, ps.Window, 0, height, cellSize, height-cellSize, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "bgBlock02":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 64, 64, 16, 16, ps.currentTheme, e.X+16+8, e.Y-8))
			break
		case "blockScreenFloor":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 3*cellSize, 80, 4*cellSize, 80+16, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenCeiling":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 3*cellSize, 112, 4*cellSize, 112+16, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenLeft":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 2*cellSize, 96, 3*cellSize, 96+16, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenRight":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 64, 96, 64+16, 96+16, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenTopLeft":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 2*cellSize, height-16, 3*cellSize, height, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenTopRight":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 4*cellSize, height-16, 5*cellSize, height, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenBottomLeft":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 2*cellSize, 80, 3*cellSize, 96, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		case "blockScreenBottomRight":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 4*cellSize, 80, 5*cellSize, 96, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		}
	}

	// stage
	for _, e := range ps.currentMap.ObjectGroups[1].Objects {
		switch e.Name {
		case "levelBlock01":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 4*cellSize, 80, 5*cellSize, 96, ps.currentTheme, e.X+16+8, 240-e.Y+8))
			break
		}
	}

	// Add player
	ps.Player = player.NewPlayer(ps.AssMan, ps.Window, 32, 32)
	ps.AddEntity(ps.Player)

	return ps
}

func (ps *PlayScreen) AddEntity(entity entity.Entityer) {
	ps.Entities = append(ps.Entities, entity)
}

func (ps *PlayScreen) RemoveEntity(entity entity.Entityer) {
	for i, e := range ps.Entities {
		if e == entity {
			ps.Entities = append(ps.Entities[:i], ps.Entities[i+1:]...)
			break
		}
	}
}

func (ps *PlayScreen) HandleInput(dt float64) {
	if ps.Player != nil {
		ps.Player.HandleInput(dt)
	}
}

func (ps *PlayScreen) CheckColliders() {
	for _, ent := range ps.Entities {
		ps.Colliders = append(ps.Colliders, ent.GetColliders()...)
	}

	//fmt.Println("TICK")

	for _, thisCollider := range ps.Colliders {
		for _, otherCollider := range ps.Colliders {
			if thisCollider.MaskBits&otherCollider.CatBits == otherCollider.CatBits { // Filter exists in mask.
				if thisCollider.Rect != otherCollider.Rect { // Don't collide with self.
					if thisCollider.Rect.Intersects(otherCollider.Rect) {
						fmt.Println("INTERSECTION")
						otherCollider.Intersection = collision.GetIntersectionDepth(thisCollider, otherCollider)
						otherCollider.IntersectionArea = collision.GetIntersectionArea(otherCollider.Intersection)

						thisCollider.IntersectingColliders = append(thisCollider.IntersectingColliders, otherCollider)
					}
				}
			}
		}

		if len(thisCollider.IntersectingColliders) > 0 {
			fmt.Println("length: ", len(thisCollider.IntersectingColliders))
		}

		// Sort array after closest intersection.
		sort.SliceStable(thisCollider.IntersectingColliders, func(i, j int) bool {
			return thisCollider.IntersectingColliders[i].IntersectionArea > thisCollider.IntersectingColliders[j].IntersectionArea
		})

		if len(thisCollider.IntersectingColliders) != 0 {
			for _, interCollider := range thisCollider.IntersectingColliders {
				//fmt.Println("this: ", thisCollider.CatBits, " intersectingArea: ", interCollider.IntersectionArea)
				thisCollider.OnCollision(interCollider)
				//interCollider.OnCollision(thisCollider) // ?
			}

			thisCollider.IntersectingColliders = thisCollider.IntersectingColliders[:0] // Clear array.
		}
	}
}

func (ps *PlayScreen) Tick(dt float64) {
	for _, ent := range ps.Entities {
		ent.Tick(dt)
	}

	// Collision detections.
	ps.CheckColliders() // FIXME Adding to and clearing array each frame is not optimal, but it will do for now.
	ps.Colliders = ps.Colliders[:0]
}

func (ps *PlayScreen) Render(dt float64) {
	ps.Window.Clear(colornames.Black)

	ps.SpriteBatch.Clear()

	for _, ent := range ps.Entities {
		ent.Render(ps.SpriteBatch, dt)
	}

	ps.SpriteBatch.Draw(ps.Window)
}

func (ps *PlayScreen) Destroy() {
	fmt.Println("DESTROYING PLAYSCREEN")
	ps.currentMap = nil
	ps.Entities = nil
	//screen.Entities = nil
}
