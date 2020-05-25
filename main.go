package main

import (
	"image/color"
	"log"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

//===== window description =====
type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// on Draw function, image and texts are draw in order of the written code. 
func (g *Game) Draw(screen *ebiten.Image) {
	// fill the back ground with color 
	// following parameter means red green blue alpha
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})

	// type message on screen instance. 
	// how to change the color of text? 
	ebitenutil.DebugPrint(screen, "Hello, User!")
	
	// draw image with option
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(1.5, 1)
	// draw image on screen 
	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
//=================================

func main() {
	// initialization for window
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("welcome Blokus")
	
	// start running instance
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}