package main

import(
    "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Hello from window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.SkyBlue)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.RayWhite)

		rl.EndDrawing()
	}
}
