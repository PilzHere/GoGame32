package menuScreen

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pilzhere/GoGame32/entity"
	"github.com/pilzhere/GoGame32/screen"
	"github.com/pilzhere/GoGame32/utils"
	"golang.org/x/image/colornames"
)

var (
//sprite *pixel.Sprite
)

type MenuScreen struct {
	screen.Screen
	sprite   *pixel.Sprite
}

func NewMenuScreen(spriteBatch *pixel.Batch, assMan *utils.AssetsManager, window *pixelgl.Window) *MenuScreen {
	fmt.Printf("THIS IS THE CONSTRUCTOR OF MENUSCREEN %p\n", &MenuScreen{})

	ms := &MenuScreen{
		Screen: screen.Screen{
			Entities: []entity.Entityer{},
			SpriteBatch: spriteBatch,
			Window:   window,
			AssMan:   assMan,
		},
		sprite: nil,
	}

	/*pic, err := utils.LoadPicture("assets/testBg.png")
	if err != nil {
		panic(err)
	}*/

	// Create sprite using image asset
	ms.sprite = pixel.NewSprite(assMan.Bg01, assMan.Bg01.Bounds())

	//spriteBatch.MakePicture(pic)

	return ms
}

func (ms *MenuScreen) AddEntity(entity entity.Entityer) {
	ms.Entities = append(ms.Entities, entity)
}

func (ms *MenuScreen) RemoveEntity(entity entity.Entityer) {
	for i, e := range ms.Entities {
		if e == entity {
			ms.Entities = append(ms.Entities[:i], ms.Entities[i+1:]...)
			break
		}
	}
}

func (ms *MenuScreen) HandleInput(dt *float64) {

}

func (ms *MenuScreen) Tick(dt *float64) {

}

func (ms *MenuScreen) Render(dt *float64) {
	ms.Window.Clear(colornames.Skyblue)

	ms.SpriteBatch.Clear()

	ms.sprite.Draw(ms.SpriteBatch, pixel.IM.Moved(ms.Window.Canvas().Bounds().Center())) // Center sprite in current view matrix.

	ms.SpriteBatch.Draw(ms.Window)
}

func (ms *MenuScreen) Destroy() {
	fmt.Println("DESTROYING MENUSCREEN")
}
