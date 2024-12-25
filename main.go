package main

import (
    "github.com/gen2brain/raylib-go/raylib"
)

const (
    SCREEN_WIDTH = 800
    SCREEN_HEIGHT = 640
)

func main() {
    initialize()

    scene := createScene()

	for !rl.WindowShouldClose() {
        handleInput(scene)
        update(scene)
        render(scene)
	}

    quit(scene)
}


func initialize() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Hello from window")
    rl.SetExitKey(rl.KeyDelete)
	rl.SetTargetFPS(60)
}

func handleInput(scene *Scene) {
    tempVelocity := rl.Vector2{X: 0, Y: 0}

    if rl.IsKeyDown(rl.KeyA) {
        tempVelocity.X -= 4;
    }
    if rl.IsKeyDown(rl.KeyD) {
        tempVelocity.X += 4;
    }

    if rl.IsKeyDown(rl.KeyW) {
        tempVelocity.Y -= 4;
    }
    if rl.IsKeyDown(rl.KeyS) {
        tempVelocity.Y += 4;
    }

    scene.player.vel = tempVelocity
}

func update(scene *Scene) {
    scene.player.update()

    for _, enemy := range(scene.enemies) {
        enemy.update()
    }
}

func render(scene *Scene) {
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    scene.draw()

    rl.EndDrawing()
}

func quit(scene *Scene) {
    rl.CloseWindow()

    rl.UnloadTexture(scene.player.sprite)

    for _, tile := range(scene.tiles) {
        rl.UnloadTexture(tile.sprite)
    }
    for _, enemy := range(scene.enemies) {
        rl.UnloadTexture(enemy.sprite)
    }
}


type Scene struct {
    player   *Player
    tiles   []Tile
    enemies []Enemy
}

func createScene() *Scene {
    player := Player{
        pos:     rl.Vector2{X: 100, Y: 100},
        vel:     rl.Vector2{X: 0,   Y: 0},
        hitbox:  rl.Rectangle{X: 100, Y: 100, Width: 100, Height: 150},
        sprite:  rl.LoadTexture("./res/placeholder-player.png"),
    }

    tiles := make([]Tile, 0)
    var groundY float32 = SCREEN_HEIGHT - 50;

    for i := 0; i < SCREEN_WIDTH; i += 50 {
        tile := Tile{
            hitbox: rl.NewRectangle(float32(i), groundY, 50, 50),
            sprite: rl.LoadTexture("./res/placeholder-tile.png"),
        }
        tiles = append(tiles, tile)
    }

    scene := Scene{
        player: &player,
        tiles: tiles,
        enemies: make([]Enemy, 0),
    }
    
    return &scene
}

func (scene *Scene) draw() {
    for _, tile := range(scene.tiles) {
        tile.draw()
    }

    scene.player.draw()

    for _, enemy := range(scene.enemies) {
        enemy.draw()
    }
}


type Player struct {
    pos    rl.Vector2
    vel    rl.Vector2
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (player *Player) update() {
    player.pos = rl.Vector2Add(player.pos, player.vel)
    player.hitbox.X = player.pos.X
    player.hitbox.Y = player.pos.Y
}

func (player *Player) draw() {
    rl.DrawTexture(player.sprite, int32(player.pos.X), int32(player.pos.Y), rl.White)
}


type Tile struct {
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (tile *Tile) draw() {
    rl.DrawTexture(tile.sprite, int32(tile.hitbox.X), int32(tile.hitbox.Y), rl.White)
}


type Enemy struct {
    pos    rl.Vector2
    vel    rl.Vector2
    hitbox rl.Rectangle
    sprite rl.Texture2D
}

func (enemy *Enemy) draw() {
    rl.DrawTexture(enemy.sprite, int32(enemy.hitbox.X), int32(enemy.hitbox.Y), rl.White)
}

func (enemy *Enemy) update() {

}
