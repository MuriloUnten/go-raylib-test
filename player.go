package main

import (
    "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
    pos    rl.Vector2
    vel    rl.Vector2
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (player *Player) Update() {
    player.pos = rl.Vector2Add(player.pos, player.vel)
    player.hitbox.X = player.pos.X
    player.hitbox.Y = player.pos.Y
}

func (player *Player) Draw() {
    rl.DrawTexture(player.sprite, int32(player.pos.X), int32(player.pos.Y), rl.White)
}
