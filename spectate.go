package ggriot

import (
	"github.com/tyhi/ggriot/cache"
	"github.com/tyhi/ggriot/models"
)

// ActiveGame will get the active game from the supplied id.
func ActiveGame(region string, summonerID string) (ag *models.ActiveGame, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/active-games/by-summoner/"+summonerID, &ag, cp); err != nil {
		return nil, err
	}

	return ag, nil
}

// FeaturedGames will get Riot's featured games.
func FeaturedGames(region string) (fg *models.FeaturedGames, err error) {
	cp := cache.CachedParams{
		Cached: true,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/featured-games", &fg, cp); err != nil {
		return nil, err
	}

	return fg, nil
}
