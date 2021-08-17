package utils

import (
	"github.com/faiface/pixel"
)

type AssetsManager struct {
	//tilesheet01Path string
	//block01Path     string
	PicTilesheet01 pixel.Picture
	PicBlock01     pixel.Picture
	Bg01           pixel.Picture
}

var (
//PicTileSheet01 = pixel.Picture(nil)
//PicBlock01 = pixel.Picture(nil)
//PicTilesheet01 pixel.Picture
//PicBlock01     pixel.Picture
)

/*func (am AssetsManager) LoadAllAssets() {
	am.tilesheet01Path = "assets/tilesheet01.png"
	am.block01Path = "assets/block.png"

	am.LoadTextures()

	fmt.Println("All assets loaded.")
}*/

/*func (am AssetsManager) LoadTextures() {
	PicTilesheet01, _ = LoadPicture(am.tilesheet01Path)
	am.PicTilesheet01 = PicTilesheet01

	PicBlock01, _ = LoadPicture(am.block01Path)
	am.PicBlock01 = PicBlock01
}*/
