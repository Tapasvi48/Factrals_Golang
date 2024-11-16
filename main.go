package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type particle struct {
	point  [2]float32 // Single 2D point
	colour color.Color
}

type Game struct {
	points    []particle // Points generated by the chaos algorithm
	vertices  []particle // Vertices of the triangle
	currPoint particle   // Current point in the chaos algorithm
	screenW   int
	screenH   int
}

// NewGame initializes the game state.
func NewGame() *Game {

	screenWidth, screenHeight := ebiten.WindowSize()

	vertices := []particle{
		{
			point:  [2]float32{float32(screenWidth) / 2, 100}, // Top vertex
			colour: color.RGBA{R: 255, G: 0, B: 0, A: 255},    // Red
		},
		{
			point:  [2]float32{float32(screenWidth) / 4, float32(screenHeight) - 100}, // Bottom-left vertex
			colour: color.RGBA{R: 0, G: 255, B: 0, A: 255},                            // Green
		},
		{
			point:  [2]float32{float32(screenWidth) * 3 / 4, float32(screenHeight) - 100}, // Bottom-right vertex
			colour: color.RGBA{R: 0, G: 0, B: 255, A: 255},                                // Blue
		},
	}

	// Initialize with no chaos points and a starting current point.
	points := []particle{}
	currPoint := particle{
		point:  [2]float32{400, 300},
		colour: color.White,
	}

	return &Game{
		points:    points,
		vertices:  vertices,
		currPoint: currPoint,
		screenW:   screenWidth,
		screenH:   screenHeight,
	}
}

// Update implements the chaos algorithm to add new points.
func (g *Game) Update() error {
	for i := 0; i < 100; i++ {
		var randId = rand.Intn(len(g.vertices))
		selectedVertex := g.vertices[randId]
		g.currPoint.point[0] = (g.currPoint.point[0] + selectedVertex.point[0]) / 2
		g.currPoint.point[1] = (g.currPoint.point[1] + selectedVertex.point[1]) / 2
		g.points = append(g.points, particle{
			point:  g.currPoint.point,
			colour: g.vertices[randId].colour,
		})
	}
	return nil
}

// Draw renders the triangle points on the screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen with a black background.
	screen.Fill(color.Black)

	// Draw the triangle vertices as larger points.
	pointThickness := float32(5.0)
	for _, vertex := range g.vertices {
		vector.DrawFilledRect(
			screen,
			vertex.point[0]-pointThickness/2, vertex.point[1]-pointThickness/2,
			pointThickness, pointThickness,
			vertex.colour,
			false,
		)
	}

	// Draw the points generated by the chaos algorithm.
	pointThickness = 2.0
	for _, p := range g.points {
		vector.DrawFilledRect(
			screen,
			p.point[0]-pointThickness/2, p.point[1]-pointThickness/2,
			pointThickness, pointThickness,
			p.colour,
			false,
		)
	}
}

// Layout implements the game layout dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenW, g.screenH
}

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Create the game instance.

	game := NewGame()

	// Set up the game window.
	ebiten.SetFullscreen(true)

	ebiten.SetWindowTitle("Sierpinski Triangle - Chaos Theory")

	// Run the game.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
