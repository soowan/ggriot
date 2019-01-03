package ggriot

import (
	"strconv"
	"strings"
	"time"

	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

var (
	// GetMasteryListExpire sets the time it takes for this cached call to be considered "expired"
	GetMasteryListExpire = time.Duration(240 * time.Minute)

	// GetTotalMasteryLevelExpire sets the time it takes for this cached call to be considered "expired"
	GetTotalMasteryLevelExpire = time.Duration(240 * time.Minute)
)

// MasteryList will return a struct with all the summoners champions and mastery exp/level.
func MasteryList(region string, summonerID string) (ml *models.MasteryList, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: GetMasteryListExpire,
		CallType:   strings.ToLower("mastery_by_summoner_" + region),
		CallKey:    summonerID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summonerID, &ml, cp); err != nil {
		return nil, err
	}

	return ml, nil
}

// ChampionMastery will return a single champion mastery struct
// TODO: Add special case for this, as it uses two inputs.
func ChampionMastery(region string, summonerID string, championID int) (ml *models.MasteryList, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summonerID+"/by-champion"+strconv.Itoa(championID), &ml, cp); err != nil {
		return nil, err
	}

	return ml, nil
}

// TotalMasteryLevel gets the total mastery level.
func TotalMasteryLevel(region string, summonerID string) (ml int, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: GetTotalMasteryLevelExpire,
		CallType:   strings.ToLower("mastery_level_" + region),
		CallKey:    summonerID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summonerID, &ml, cp); err != nil {
		return 0, err
	}

	return ml, nil
}
