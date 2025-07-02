package main

import (
	"encoding/json"
	"os"
)

type Game struct {
    GameId string `json:"gameId"`
    Title  string `json:"title"`
}

type Response struct {
    Count int    `json:"count"`
    Data  []Game `json:"data"`
}

func loadGamesFromJSON(path string) (*Response, error) {
    file, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var resp Response
    err = json.Unmarshal(file, &resp)
    return &resp, err
}
