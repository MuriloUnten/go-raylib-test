package main

import (
    "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
    pos    rl.Vector2
    vel    rl.Vector2
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (enemy *Enemy) Draw() {
    rl.DrawTexture(enemy.sprite, int32(enemy.hitbox.X), int32(enemy.hitbox.Y), rl.White)
}

func (enemy *Enemy) Update(elapsedTime float32) {

}
