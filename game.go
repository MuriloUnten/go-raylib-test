package main

import (
    "github.com/gen2brain/raylib-go/raylib"
)

const (
    SCREEN_WIDTH = 800
    SCREEN_HEIGHT = 640
)

type Game struct {
    scenes []*Scene
    currentSceneIndex uint8
    currentScene *Scene

    player *Player
}

func (g *Game) Init() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Hello from window")
    rl.SetExitKey(rl.KeyDelete)
	rl.SetTargetFPS(60)

    g.initializeGameObjects()
}

func (g *Game) initializeGameObjects() {
    g.player = &Player{
        pos:     rl.Vector2{X: 100, Y: 100},
        vel:     rl.Vector2{X: 0,   Y: 0},
        hitbox:  rl.Rectangle{X: 100, Y: 100, Width: 100, Height: 150},
        sprite:  rl.LoadTexture("./res/placeholder-player.png"),
    }

    g.scenes = make([]*Scene, 0)
    g.scenes = append(g.scenes, CreateScene1(g.player))
    g.scenes = append(g.scenes, CreateScene2(g.player))
    g.currentSceneIndex = 0
    g.currentScene = g.scenes[g.currentSceneIndex]
}

func (g *Game) Run() {
    for !rl.WindowShouldClose() {
        g.handleInput()
        g.currentScene.Update()
        g.currentScene.Render()
    }
}

func (g *Game) swapScenes() {
    if g.currentSceneIndex == 0 {
        g.currentSceneIndex = 1
    } else {
        g.currentSceneIndex = 0
    }
    g.currentScene = g.scenes[g.currentSceneIndex]
}

func (g *Game) handleInput() {
    if rl.IsKeyPressed(rl.KeyR) {
        g.swapScenes()
    }

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

    g.currentScene.player.vel = tempVelocity
}

func (g *Game) Quit() {
    rl.CloseWindow()

    rl.UnloadTexture(g.player.sprite)

    for _, scene := range(g.scenes) {
        for _, tile := range(scene.tiles) {
            rl.UnloadTexture(tile.sprite)
        }
        for _, enemy := range(scene.enemies) {
            rl.UnloadTexture(enemy.sprite)
        }
    }
}
