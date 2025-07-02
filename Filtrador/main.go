package main

import (
	"fmt"
)

func main() {
    ids, err := readGameIDsFromExcel("ids.xlsx")
    if err != nil {
        panic(err)
    }

    resp, err := loadGamesFromJSON("games.json")
    if err != nil {
        panic(err)
    }

    filtered := filterGamesByIDs(resp.Data, ids)

    fmt.Printf("🎯 Juegos encontrados: %d\n", len(filtered))
    for _, game := range filtered {
        fmt.Printf("GameID: %s | Título: %s\n", game.GameId, game.Title)
    }
}
