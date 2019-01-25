package ggriot

import (
	"strings"
	"time"

	"github.com/tyhi/ggriot/cache"

	"github.com/tyhi/ggriot/models"
)

var (
	// ExpireSummoner is how long the summoner by ign cache is saved.
	ExpireSummoner = time.Duration(30 * time.Minute)
)

// SummonerByAccID will get summoner information using Account ID
func SummonerByAccID(region string, accountID string) (s *models.Summoner, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accountID, &s, cp); err != nil {
		return nil, err
	}

	return s, nil
}

// SummonerBySumID will get summoner information using Summoner ID
func SummonerBySumID(region string, summonerID string) (s *models.Summoner, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+summonerID, &s, cp); err != nil {
		return nil, err
	}

	return s, nil
}

// SummonerByIGN will get summoner information using IGN
func SummonerByIGN(region string, summonerIGN string) (s *models.Summoner, err error) {
	summonerIGN = strings.ToLower(summonerIGN)
	summonerIGN = strings.Replace(summonerIGN, " ", "", -1)

	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireSummoner,
		CallType:   "summoner_by_ign_" + region,
		CallKey:    summonerIGN,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s, cp); err != nil {
		return nil, err
	}

	return s, nil
}

// SummonerByPUUID will get summoner information using IGN
func SummonerByPUUID(region string, summonerPUUID string) (s *models.Summoner, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireSummoner,
		CallType:   "summoner_by_puuid_" + region,
		CallKey:    summonerPUUID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-puuid/"+summonerPUUID, &s, cp); err != nil {
		return nil, err
	}

	return s, nil
}
