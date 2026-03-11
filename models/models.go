package models

type Paddle struct {
	X float64
	Y float64

	Width  float64
	Height float64

	Speed float64
}

func (p *Paddle) MoveUp() {
	p.Y -= p.Speed
}

func (p *Paddle) MoveDown() {
	p.Y += p.Speed
}

type Ball struct {
	X  float64
	Y  float64
	DX float64
	DY float64

	Size float64
}

func (b *Ball) Update() {
	b.X += b.DX
	b.Y += b.DY
}
