package ggriot

import (
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

// ChampionRotation is for getting the champion rotation for the week.
// This is the new api endpoint added. Subject to change.
func ChampionRotation(region string) (cr *models.ChampionRotation, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseChampionR, &cr, cp); err != nil {
		return nil, err
	}

	return cr, nil
}
