package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/pilzhere/GoGame32/screen"
	"github.com/pilzhere/GoGame32/screen/menuScreen"
	"github.com/pilzhere/GoGame32/screen/playScreen"
	"github.com/pilzhere/GoGame32/utils"
	"strconv"
	"time"
)

var (
	frameTick              *time.Ticker
	now                    time.Time
	dt                     float64
	oneSecondTimer         float64
	fps                    int32
	canvasFbo              *pixelgl.Canvas
	windowScale            float64
	windowScaleOld         float64
	canvasFboWidth         float64
	canvasFboHeight        float64
	windowOriginalPosX     float64
	windowOriginalPosY     float64
	windowIsFullscreen     = false
	windowedWindowOldScale float64
	screens                []screen.Screener
	currentScreen          screen.Screener
	assetsMan              utils.AssetsManager
	windowRepositionTimer  float64
	windowRepositionToggle bool
	spriteBatch            *pixel.Batch
	T                      = 0.0
	DT                     = 0.01
	currentTime            = time.Now()
	accumulator            = 0.0
	sdt                    = 1.0
)

const (
	windowOriginalWidth  float64 = 426
	windowOriginalHeight float64 = 240
	fpsLimit             int     = 60 // FIXME Game is designed for 60 fps ONLY. Increasing or lowering this will alter the physics (mostly jumps) of the game (there is no interpolation in the code).
)

func main() {
	pixelgl.Run(run)
}

func run() {
	windowScale = 1
	canvasFboWidth = windowOriginalWidth
	canvasFboHeight = windowOriginalHeight

	// Create window configuration
	cfg := pixelgl.WindowConfig{
		Title:     "GoGame32",
		Bounds:    pixel.R(0, 0, windowOriginalWidth*windowScale, windowOriginalHeight*windowScale), // Same height as a NES game.
		VSync:     false,                                                                            // 75 fps on my screen.
		Resizable: true,
		Maximized: false,
	}

	// Create window using configuration
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Save original window positions for when resetting window scale to 1.
	windowOriginalPosX = win.GetPos().X
	windowOriginalPosY = win.GetPos().Y

	setFPS(fpsLimit)

	canvasFbo = pixelgl.NewCanvas(pixel.R(0, 0, canvasFboWidth, canvasFboHeight))

	// Load assets
	bg01, _ := utils.LoadPicture("assets/testBg.png")
	tilesheet01, _ := utils.LoadPicture("assets/tilesheet01.png")
	block01, _ := utils.LoadPicture("assets/block.png")

	assetsMan = utils.AssetsManager{
		PicTilesheet01: tilesheet01,
		PicBlock01:     block01,
		Bg01:           bg01,
	}

	/*spriteBatch = pixel.Batch{}
	spriteBatch.MakePicture(tilesheet01)
	spriteBatch.MakeTriangles(&pixel.TrianglesData{})
	spriteBatch.SetMatrix(pixel.IM)
	spriteBatch.SetColorMask(color.Alpha{A: 1})*/

	spriteBatch = pixel.NewBatch(&pixel.TrianglesData{}, bg01)

	//fmt.Println("bathc: ", spriteBatch)

	PushScreen(menuScreen.NewMenuScreen(spriteBatch, &assetsMan, win))

	for !win.Closed() {
		dt = calculateDeltaTime()

		/*if dt > 1/60.0 {
			sdt = dt / (1/60.0)
		} else {
			sdt = 1.0
		}*/

		if windowRepositionToggle {
			windowRepositionTimer += dt

			if windowRepositionTimer > 1/24 { // wait 0.0416 ms.
				windowRepositionTimer = 0
				windowRepositionToggle = false
			}
		}

		// Input
		if win.JustPressed(pixelgl.KeyEscape) {
			exitGame(win)
			return // exit while loop
		}

		if win.JustPressed(pixelgl.KeyP) {
			fmt.Println("Pressed P")

			if !windowRepositionToggle {
				if !windowIsFullscreen {
					windowScaleOld = windowScale

					if windowScale != 5 {
						windowScale++
						resizeWindowedWindow(win)
					}
				}
			}
		}

		if win.JustPressed(pixelgl.KeyM) {
			fmt.Println("Pressed M")

			if !windowRepositionToggle {
				if !windowIsFullscreen {
					windowScaleOld = windowScale

					if windowScale != 1 {
						windowScale--
						resizeWindowedWindow(win)
					}
				}
			}
		}

		if win.JustPressed(pixelgl.KeyE) {
			fmt.Println("Pressed E")

			spriteBatch = pixel.NewBatch(&pixel.TrianglesData{}, tilesheet01)
			PushScreen(playScreen.NewPlayScreen(spriteBatch, &assetsMan, win))
		}

		if win.JustPressed(pixelgl.KeyQ) {
			fmt.Println("Pressed Q")

			PopScreen()
		}

		//fmt.Println(len(screens))

		if win.JustPressed(pixelgl.KeyF) {
			fmt.Println("Pressed F")

			if !windowIsFullscreen {
				monitorWidth, monitorHeight := pixelgl.PrimaryMonitor().Size()

				windowedWindowOldScale = windowScale

				win.SetBounds(pixel.R(0, 0, monitorWidth, monitorHeight))
				win.SetMonitor(pixelgl.PrimaryMonitor())
			} else {
				windowScale = windowedWindowOldScale

				win.SetMonitor(nil) // To use current monitor.
				win.SetBounds(pixel.R(0, 0, windowOriginalWidth*windowScale, windowOriginalHeight*windowScale))
			}

			windowIsFullscreen = !windowIsFullscreen
		}

		oneSecondTimer += dt
		if oneSecondTimer >= 1 { // Has one second passed?
			//fmt.Printf("fps: %d dt: %f ms\n", fps, dt)
			strDt := strconv.FormatFloat(dt, 'f', 6, 64)
			strFps := strconv.FormatInt(int64(fps), 10)
			var title = "GoGame32 fps: " + strFps + " dt: " + strDt + " ms"
			win.SetTitle(title)
			fps = 0
			oneSecondTimer = 0 // Reset timer
		}

		currentScreen.HandleInput(dt)

		fmt.Println("-new tick-")

		currentScreen.Tick(dt)

		for accumulator >= DT {
			//previousState = currentState;
			//integrate( currentState, t, dt )
			T += DT
			accumulator -= dt
		}

		//alpha := accumulator / DT

		var canvasBounds = pixel.Rect{
			Min: pixel.Vec{},
			Max: pixel.Vec{X: windowOriginalWidth, Y: windowOriginalHeight},
		}
		win.Canvas().SetBounds(canvasBounds) // Scale down window canvas
		canvasFbo.SetBounds(canvasBounds)    // Scale down FBO canvas

		currentScreen.Render(dt)

		if !windowIsFullscreen {
			canvasBounds = pixel.Rect{
				Min: pixel.Vec{},
				Max: pixel.Vec{X: windowOriginalWidth * windowScale, Y: windowOriginalHeight * windowScale},
			}
		} else {
			width, height := pixelgl.PrimaryMonitor().Size()
			canvasBounds = pixel.Rect{
				Min: pixel.Vec{},
				Max: pixel.Vec{X: width, Y: height},
			}
		}

		canvasFbo.SetPixels(win.Canvas().Pixels()) // Paint CanvasFBO with window's canvas pixels.
		win.Canvas().SetBounds(canvasBounds)       // Update window's canvas size. Else we can't paint CanvasFBO on top of it.

		//win.Clear(colornames.Brown) // test

		// Move canvasFBO to middle of the window's bounds.
		var canvasFboPoint = pixel.Vec{}

		if !windowIsFullscreen {
			// Render CanvasFBO in middle, scaled up to window's bounds.
			moveWidth := -win.Bounds().W()/2*(win.Bounds().W()/canvasFboWidth) + canvasFboWidth*windowScale/2
			moveHeight := -win.Bounds().H()/2*(win.Bounds().H()/canvasFboHeight) + canvasFboHeight*windowScale/2
			var canvasFboMove = pixel.Vec{X: moveWidth, Y: moveHeight}

			canvasFbo.Draw(win, pixel.IM.Moved(win.Canvas().Bounds().Center()).Scaled(canvasFboPoint, windowScale).Moved(canvasFboMove))
		} else {
			// Render CanvasFBO in bottom-left corner, scaled up to monitor's/window's bounds.
			monitorWidth, monitorHeight := pixelgl.PrimaryMonitor().Size()
			xPos := win.Canvas().Bounds().Min.X + canvasFboWidth/2
			yPos := win.Canvas().Bounds().Min.Y + canvasFboHeight/2
			pos := pixel.Vec{X: xPos, Y: yPos}

			canvasFbo.Draw(win, pixel.IM.Moved(pos).ScaledXY(canvasFboPoint, pixel.Vec{X: monitorWidth / canvasFboWidth, Y: monitorHeight / canvasFboHeight}))
		}

		win.Update() // Frame is finished, swap!

		if frameTick != nil {
			<-frameTick.C
		}

		fps++
	}
}

func PushScreen(screen screen.Screener) {
	screens = append(screens, screen)       // push screen on top
	currentScreen = screens[len(screens)-1] // currentScreen is top
}

func PopScreen() {
	if len(screens) > 1 {
		currentScreen.Destroy()

		top := len(screens) - 1
		screens = (screens)[:top]

		currentScreen = screens[len(screens)-1] // currentScreen is top
	}
}

func PopAllScreens() {
	fmt.Println("Popping all", len(screens), "screens")

	for len(screens) > 0 {
		currentScreen.Destroy()

		top := len(screens) - 1
		screens = (screens)[:top]

		if len(screens) != 0 {
			currentScreen = screens[len(screens)-1] // currentScreen is top
		} else {
			currentScreen = nil
		}
	}
}

func resizeWindowedWindow(window *pixelgl.Window) {
	newBounds := pixel.R(0, 0, windowOriginalWidth*windowScale, windowOriginalHeight*windowScale)
	window.SetBounds(newBounds)
	var newPos pixel.Vec
	if windowScale != 1 {
		if windowScaleOld < windowScale {
			newPos = pixel.Vec{X: window.GetPos().X - (newBounds.W()/windowScale)/2, Y: window.GetPos().Y - (windowOriginalHeight * windowScale / (windowScale * 2))}
		} else {
			newPos = pixel.Vec{X: window.GetPos().X + (newBounds.W()/windowScale)/2, Y: window.GetPos().Y + (windowOriginalHeight * windowScale / (windowScale * 2))}
		}
	} else {
		newPos = pixel.Vec{X: windowOriginalPosX, Y: windowOriginalPosY}
	}

	window.SetPos(newPos)

	windowRepositionToggle = true
}

func setFPS(fps int) {
	if fps <= 0 {
		frameTick = nil
	} else {
		frameTick = time.NewTicker(time.Second / time.Duration(fps))
	}
}

func calculateDeltaTime() float64 {
	last := now
	now = time.Now()
	deltaTime := now.Sub(last).Seconds()

	return deltaTime
}

func calculateAlpha() float64 {
	newTime := time.Now()
	frameTime := newTime.Sub(time.Now()).Seconds()
	if frameTime > 0.25 {
		frameTime = 0.25
	}
	currentTime = newTime

	accumulator += frameTime
	return 0
}

func exitGame(window *pixelgl.Window) {
	PopAllScreens()

	window.SetClosed(true)

	fmt.Println("Game exiting... Thanks for playing! <3")
}
