package main

func filterGamesByIDs(games []Game, ids map[string]bool) []Game {
	var filtered []Game
	for _, game := range games {
		if ids[game.GameId] {
			filtered = append(filtered, game)
		}
	}
	return filtered
}