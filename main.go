package main

func main() {
    game := new(Game)
    game.Init()

    game.Run()
    game.Quit()
}
