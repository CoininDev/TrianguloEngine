package systems

import (
	"triango/comps"
	"triango/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderSystem struct{}

func (s *RenderSystem) Start(g *game.Game) {}
func (s *RenderSystem) End(g *game.Game) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Sprite"); i != -1 {
			s := e.Components[i].(*comps.Sprite)
			rl.UnloadTexture(s.Texture)
		}
	}
}
func (s *RenderSystem) Update(g *game.Game, delta float64) {
	for _, e := range g.Entities {
		if i := e.GetComponent("Sprite"); i != -1 {
			s := e.Components[i].(*comps.Sprite)
			t := e.Components[e.GetComponent("Transform")].(*comps.Transform)
			if !s.Active {
				continue
			}
			rl.DrawTexturePro(
				s.Texture,
				rl.NewRectangle(0, 0, float32(s.Texture.Width*int32(s.Fliph)), float32(s.Texture.Height*int32(s.Flipv))),
				rl.NewRectangle(0, 0, float32(s.Texture.Width)*t.Scale.X, float32(s.Texture.Height)*t.Scale.Y),
				rl.Vector2(t.Position.Negative()),
				t.Rotation,
				rl.White,
			)
		}
		if i := e.GetComponent("PlaceholderSprite"); i != -1 {
			s := e.Components[i].(*comps.PlaceholderSprite)
			t := e.Components[e.GetComponent("Transform")].(*comps.Transform)
			if !s.Active {
				continue
			}
			rl.DrawRectanglePro(
				rl.NewRectangle(0, 0, s.Size.X, s.Size.Y),
				rl.Vector2(t.Position.Negative()),
				t.Rotation,
				s.Color,
			)
		}
	}
}
