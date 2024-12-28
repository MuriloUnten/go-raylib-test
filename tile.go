package main

import (
    "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (tile *Tile) Draw() {
    rl.DrawTexture(tile.sprite, int32(tile.hitbox.X), int32(tile.hitbox.Y), rl.White)
}
