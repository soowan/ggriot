package ggriot

import (
	"github.com/tyhi/ggriot/cache"
	"github.com/tyhi/ggriot/models"
)

// ServerStatus returns the current service status for the requested region.
// This doesn't apply to rate limits, however we are still obeying the limit.
// This will change in the future possibly.
func ServerStatus(region string) (ss *models.ServerStatus, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseStatus, &ss, cp); err != nil {
		return nil, err
	}

	return ss, nil
}
