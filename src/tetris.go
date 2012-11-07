package main

// TODO: keydown

import (
	"fmt"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"math/rand"
//	"os"
//	"strings"
	"time"
)

type Coordinate struct{ x, y int }
type TetPosition [4]Coordinate
type Tetromino struct {
	position [4]TetPosition
	color int
}

func colorToUint32(color sdl.Color) (uint32) {
	return uint32(color.R) << 16 | 
		uint32(color.G) << 8 |
		uint32(color.B)
}

var colorTable = [8]sdl.Color{
	sdl.Color{0,0,0,0},
	sdl.Color{0xaa,0xaa,0xaa,0},
	sdl.Color{0xee,0x0,0x0,0},
	sdl.Color{0x0,0xee,0x0,0},
	sdl.Color{0x0,0x0,0xee,0},
	sdl.Color{0x0,0x99,0x99,0},
	sdl.Color{0x99,0x0,0x99,0},
	sdl.Color{0x99,0x99,0x0,0},
}

type Playfield struct {
	// Use a single underlying array--
	// otherwise we have to muddle with
	// slices of slices.
	width int
	height int
	field []int
}

var tSquare = Tetromino{
	[4]TetPosition{ 
		{ { 1, 1 }, { 0, 0 }, { 0, 1 }, { 1, 0 } },
		{ { 1, 1 }, { 0, 0 }, { 0, 1 }, { 1, 0 } },
		{ { 1, 1 }, { 0, 0 }, { 0, 1 }, { 1, 0 } },
		{ { 1, 1 }, { 0, 0 }, { 0, 1 }, { 1, 0 } }, 
	}, 1, }

var tJog1 = Tetromino{
	[4]TetPosition{ 
		{ { 0, 0 }, { 1, 0 }, { 1, 1 }, { 2, 1 } },
		{ { 1, 0 }, { 0, 1 }, { 1, 1 }, { 0, 2 } },
		{ { 0, 0 }, { 1, 0 }, { 1, 1 }, { 2, 1 } },
		{ { 1, 0 }, { 0, 1 }, { 1, 1 }, { 0, 2 } },
	}, 2 }

var tJog2 = Tetromino{ 
	[4]TetPosition{ 
		{ { 1, 0 }, { 2, 0 }, { 0, 1 }, { 1, 1 } },
		{ { 0, 0 }, { 0, 1 }, { 1, 1 }, { 1, 2 } },
		{ { 1, 0 }, { 2, 0 }, { 0, 1 }, { 1, 1 } },
		{ { 0, 0 }, { 0, 1 }, { 1, 1 }, { 1, 2 } },
	}, 3 }

var tTee = Tetromino{
	[4]TetPosition{ 
		{ { 1, 1 }, { 0, 1 }, { 2, 1 }, { 1, 0 } },
		{ { 1, 1 }, { 1, 2 }, { 2, 1 }, { 1, 0 } },
		{ { 1, 1 }, { 0, 1 }, { 2, 1 }, { 1, 2 } },
		{ { 1, 1 }, { 0, 1 }, { 1, 2 }, { 1, 0 } },
	}, 4 }

var tEl1 = Tetromino{
	[4]TetPosition{ 
		{ {1,1},{1,0},{2,0},{1,2} },
		{ {1,1},{0,1},{2,1},{2,2} },
		{ {1,1},{1,0},{0,2},{1,2} },
		{ {1,1},{0,0},{0,1},{2,1} },
	}, 5 }

var tEl2 = Tetromino{
	[4]TetPosition{ 
		{ {1,1},{0,0},{1,0},{1,2} },
		{ {1,1},{2,0},{0,1},{2,1} },
		{ {1,1},{1,0},{1,2},{2,2} },
		{ {1,1},{0,1},{0,2},{2,1} },
	}, 6 }

var tLong = Tetromino{
	[4]TetPosition{ 
		{ {1,0},{1,1},{1,2},{1,3} },
		{ {0,1},{1,1},{2,1},{3,1} },
		{ {1,0},{1,1},{1,2},{1,3} },
		{ {0,1},{1,1},{2,1},{3,1} },
	}, 7 }

var Tetrominos = []Tetromino{
	tSquare,
	tJog1,
	tJog2,
	tTee,
	tEl1,
	tEl2,
	tLong,
}

func NewPlayfield(width, height int) (Playfield) {
	var p Playfield
	p.width = width
	p.height = height
	p.field = make([]int, width*height)
	return p
}

// -1 is out-of-bounds
func (p *Playfield) at(x,y int) (int) {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return -1
	}
	return p.field[x+y*p.width]
}

func (p *Playfield) isLineComplete(y int) (bool) {
	for i := 0; i < p.width; i++ {
		if p.at(i,y) == 0 { return false }
	}
	return true
}

func (p *Playfield) removeLine(y int) {
	if y > 0 {
		for i := y; i > 0; i-- {
			for j := 0; j < p.width; j++ {
				p.set(j,i,p.at(j,i-1))
			}
		}
	}
	for j := 0; j < p.width; j++ {
		p.field[0] = 0
	}
}

func (p *Playfield) removeCompletedLines() (int) {
	count := 0
	for i := 0; i < p.height; i++ {
		if p.isLineComplete(i) {
			p.removeLine(i)
			count++
		}
	}
	return count
}

func (p *Playfield) set(x,y,val int) {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		panic("Accessing cell outside of playfield")
	}
	p.field[x+y*p.width] = val
}

func (p *Playfield) print() {
	for i := 0; i < p.height; i++ {
		for j := 0; j < p.width; j++ {
			v := p.at(j,i)
			if v == 0 { 
				fmt.Print(" ") 
			} else { 
				fmt.Print(p.at(j,i)) 
			}
		}
		fmt.Println()
	}
}

func (p *Playfield) place(piece *Piece) {
	for i := 0; i < 4; i++ {
		r := piece.rot
		p.set(
			piece.pos.x+piece.tet.position[r][i].x,
			piece.pos.y+piece.tet.position[r][i].y,
			piece.tet.color)
	}
}

func (p *Playfield) render(screen *sdl.Surface) {
	screen.FillRect(nil,0x000000)
	for j := 0; j < p.height; j++ {
		for i := 0; i < p.width; i++ {
			color := p.at(i,j)
			if color != 0 {
				rect := sdl.Rect{
					int16(i*PixPerBrick),
					int16(j*PixPerBrick),
					uint16(PixPerBrick),
					uint16(PixPerBrick) }
				screen.FillRect(&rect, colorToUint32(colorTable[color]))
			}
		}
	}
}

type Piece struct { 
	tet *Tetromino
	pos Coordinate
	rot int }

func (p *Piece) render(screen *sdl.Surface) {
	if p == nil { return }
	for i := 0; i < 4; i++ {
		x := p.pos.x + p.tet.position[p.rot][i].x
		y := p.pos.y + p.tet.position[p.rot][i].y
		rect := sdl.Rect{
			int16(x*PixPerBrick),
			int16(y*PixPerBrick),
			uint16(PixPerBrick),
			uint16(PixPerBrick) }
		screen.FillRect(&rect, colorToUint32(colorTable[p.tet.color]))
	}
}

func (p *Piece) move(field *Playfield, xi, yi int) (bool) {
	if p == nil { return false }
	tx := p.pos.x + xi
	ty := p.pos.y + yi
	for i := 0; i < 4; i++ {
		x := tx + p.tet.position[p.rot][i].x
		y := ty + p.tet.position[p.rot][i].y
		if field.at(x,y) != 0 { 
			return false 
		}
	}
	p.pos.x = tx
	p.pos.y = ty
	return true
}

func (p *Piece) rotate(field *Playfield, ri int) (bool) {
	if p == nil { return false }
	tr := int(uint(p.rot + ri) % 4)
	for i := 0; i < 4; i++ {
		x := p.pos.x + p.tet.position[tr][i].x
		y := p.pos.y + p.tet.position[tr][i].y
		if field.at(x,y) != 0 { 
			return false 
		}
	}
	p.rot = tr
	return true
}

func makeNewTet() (*Piece) {
	tet := &Tetrominos[rand.Intn(len(Tetrominos))]
	return &Piece{ tet, Coordinate{0,0}, 0 }
}

const StandardWidth int = 10
const StandardHeight int = 20
const PixPerBrick int = 20
func main() {
	var p Playfield = NewPlayfield(StandardWidth, StandardHeight)

	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		panic(sdl.GetError())
	}


	var screen = sdl.SetVideoMode(p.width * PixPerBrick,
		p.height * PixPerBrick, 32, sdl.RESIZABLE)

	if screen == nil {
		panic(sdl.GetError())
	}

	//var video_info = sdl.GetVideoInfo()

	sdl.EnableUNICODE(1)

	sdl.WM_SetCaption("YATI:Go", "")

	
	dropTicker := time.NewTicker(time.Second / 2.0)
	var curPiece *Piece = nil
	score := 0
mainloop:
	for {
		select {
		case <-dropTicker.C:
			if curPiece == nil {
				curPiece = makeNewTet()
				curPiece.pos.x = p.width/2 - 2
				if !curPiece.move(&p,0,0) {
					fmt.Println("Failed placement test")
					break mainloop
				}
			} else {
				// advance piece
				if !curPiece.move(&p,0,1) {
					p.place(curPiece)
					// remove lines
					score += p.removeCompletedLines()
					fmt.Println(score)
					timeDiv := 2.0 + float64(score/10)/4.0
					duration := float64(time.Second) / timeDiv
					dropTicker = time.NewTicker(time.Duration(duration))
					curPiece = nil
				}
			}
			
		case _event := <-sdl.Events:
			switch e := _event.(type) {
			case sdl.QuitEvent:
				break mainloop
			case sdl.KeyboardEvent:
				//key := sdl.GetKeyName(sdl.Key(e.Keysym.Sym))
				//println(e.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keysym.Sym)))
				if (e.Type == sdl.KEYDOWN) {
					switch e.Keysym.Sym {
					case sdl.K_ESCAPE:
						break mainloop
					case sdl.K_w:
						curPiece.rotate(&p,-1)
					case sdl.K_a:
						curPiece.move(&p,-1,0)
					case sdl.K_d:
						curPiece.move(&p,1,0)
					case sdl.K_s:
						//curPiece.rotate(&p,1)
						fallthrough
					case sdl.K_SPACE:
						// turn on drop mode
						dropTicker = time.NewTicker(time.Second / 60.0)
					}

				}
			}
		}
		p.render(screen)
		curPiece.render(screen)
		screen.Flip()
	}
	//fmt.Println("----")
	//p.removeLine(5)
	//p.print()
}