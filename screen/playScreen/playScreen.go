package playScreen

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/entity/block"
	"github.com/pilzhere/GoGame32/entity/player"
	"github.com/pilzhere/GoGame32/screen"
	"github.com/pilzhere/GoGame32/utils"
	"golang.org/x/image/colornames"
)

type PlayScreen struct {
	screen.Screen
	currentMap   *tiled.Map
	currentTheme int
}

func NewPlayScreen(spriteBatch *pixel.Batch, assMan *utils.AssetsManager, window *pixelgl.Window) *PlayScreen {
	fmt.Printf("THIS IS THE CONSTRUCTOR OF PLAYSCREEN %p\n", &PlayScreen{})

	ps := &PlayScreen{
		Screen:screen.Screen{
			Entities: []entity.Entityer{},
			SpriteBatch: spriteBatch,
			Window:   window,
			AssMan:   assMan,
		},
	}

	ps.currentMap, _ = tiled.LoadFromFile("assets/maps/map01.tmx")
	ps.currentTheme = ps.currentMap.Properties.GetInt("theme")
	height := float64(ps.currentMap.Tilesets[0].Image.Height)
	//width := float64(ps.currentMap.Tilesets[0].Image.Width)
	cellSize := 16.0

	//println("assMan: ", ps.AssMan)

	//spriteBatch = pixel.NewBatch(&pixel.TrianglesData{}, assMan.PicTilesheet01)
	//spriteBatch.MakePicture(assMan.PicTilesheet01)

	for _, e := range ps.currentMap.ObjectGroups[0].Objects {
		switch e.Name {
		case "bgBlock01":
			ps.AddEntity(block.NewBlock(ps.AssMan, ps.Window, 0, height, cellSize, height-cellSize, ps.currentTheme, e.X+16+8, e.Y-8))
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

	ps.AddEntity(player.NewPlayer(ps.AssMan, ps.Window, 32, 32))

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

/*func (ps *PlayScreen) GetWindow() *pixelgl.Window {
	return ps.Window
}*/

func (ps *PlayScreen) HandleInput(dt *float64) {

}

func (ps *PlayScreen) Tick(dt *float64) {
	//fmt.Println("-PlaysScreen TICK!")

	for _, ent := range ps.Entities {
		ent.Tick(dt)
		//fmt.Println(&ent)
	}

	//fmt.Println("ents:", len(screen.Entities))
}

func (ps *PlayScreen) Render(dt *float64) {
	ps.Window.Clear(colornames.Skyblue)

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

// DELAR ALLA SCREENS SAMMA INTERFACE!??!?!?!?
// Spela och spawna screens. Alla block försvinner om EN screen försvinner...
