package main

import (
    "github.com/gen2brain/raylib-go/raylib"
    "os"
    "encoding/csv"
    "log"
)

type Scene struct {
    player   *Player
    tiles   []Tile
    enemies []Enemy
    sprites map[string]rl.Texture2D
}

func CreateScene1(player *Player) *Scene {
    tiles := make([]Tile, 0)

    sprites := make(map[string]rl.Texture2D, 0)
    sprites["base-tile"] = rl.LoadTexture("./res/placeholder-tile-light.png")

    tileMap := readTileMap("./res/placeholder-universe-1.csv")
    for rowIndex, row := range(tileMap) {
        for colIndex, tile := range(row) {
            if tile == "-1" {
                continue
            }

            if tile == "0" {
                baseTile := Tile{
                    hitbox: rl.NewRectangle(float32(colIndex * 50), float32(rowIndex * 50), 50, 50),
                    sprite: sprites["base-tile"],
                }
                tiles = append(tiles, baseTile)
            }
        }
    }

    scene := Scene{
        player: player,
        tiles: tiles,
        enemies: make([]Enemy, 0),
    }
    
    return &scene
}

func CreateScene2(player *Player) *Scene {
    tiles := make([]Tile, 0)

    sprites := make(map[string]rl.Texture2D, 0)
    sprites["base-tile"] = rl.LoadTexture("./res/placeholder-tile-dark.png")

    tileMap := readTileMap("./res/placeholder-universe-2.csv")
    for rowIndex, row := range(tileMap) {
        for colIndex, tile := range(row) {
            if tile == "-1" {
                continue
            }

            if tile == "1" {
                baseTile := Tile{
                    hitbox: rl.NewRectangle(float32(colIndex * 50), float32(rowIndex * 50), 50, 50),
                    sprite: sprites["base-tile"],
                }
                tiles = append(tiles, baseTile)
            }
        }
    }

    scene := Scene{
        player: player,
        tiles: tiles,
        enemies: make([]Enemy, 0),
    }
    
    return &scene
}

func (s *Scene) Update(elapsedTime float32) {
    s.player.Update(elapsedTime)

    for _, enemy := range(s.enemies) {
        enemy.Update(elapsedTime)
    }
}

func (s *Scene) Render() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    s.Draw()

    rl.EndDrawing()
}

func (s *Scene) Draw() {
    for _, tile := range(s.tiles) {
        tile.Draw()
    }

    s.player.Draw()

    for _, enemy := range(s.enemies) {
        enemy.Draw()
    }
}

func readTileMap(fileMap string) [][]string {
    file, err := os.Open(fileMap)
    if err != nil {
        log.Fatal("Failed to load file: ", fileMap)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    tileMap, err := reader.ReadAll()
    if err != nil {
        log.Fatal("Failed to read from file: ", fileMap)
    }

    return tileMap
}
