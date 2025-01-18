package lib

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	IsActive bool
	e        *Entity
	Texture  rl.Texture2D
}

func NewSprite(tex rl.Texture2D) *Sprite {
	return &Sprite{true, nil, tex}
}
func (s *Sprite) Start(e *Entity) { s.e = e }
func (s *Sprite) Tick() {
	rl.DrawTexture(s.Texture, s.e.Position.X, s.e.Position.Y, rl.White)
}
func (s *Sprite) End() {
	rl.UnloadTexture(s.Texture)
}
func (s *Sprite) GetType() string { return "Sprite" }

type PlaceholderSprite struct {
	IsActive bool
	e        *Entity
	Size     Vector2
	Color    rl.Color
}

func NewPlaceholderSprite(size Vector2, color rl.Color) *PlaceholderSprite {
	return &PlaceholderSprite{true, nil, size, color}
}

func (s *PlaceholderSprite) Start(e *Entity) { s.e = e }
func (s *PlaceholderSprite) Tick() {
	rl.DrawRectangle(s.e.Position.X, s.e.Position.Y, s.Size.X, s.Size.Y, s.Color)
}
func (s *PlaceholderSprite) End()            {}
func (s *PlaceholderSprite) GetType() string { return "PlaceholderSprite" }
