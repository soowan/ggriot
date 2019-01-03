package cache

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Cached is the model in which we store all our cached api calls.
// This is a low level cache so it only store the same key that you'd call the official riot API.
// Maybe in the future this could be changed to efficiently index when two calls return the same data.
type Cached struct {
	gorm.Model

	CallKey string `sql:"string"`
	JSON    string `sql:"type:jsonb"`
}

// CachedParams is a interface that is send with the call.
// This will make all caching happen the same for all calls, so that way any changes with the cache can be changed in one place.
type CachedParams struct {
	Cached     bool
	Expire     bool
	Expiration time.Duration
	CallKey    string
	CallType   string
}

// Summoner is a custom type struct.
// This was created because there are three different calls that all get the same data.
type Summoner struct {
	gorm.Model

	IGN        string `sql:"int"`
	AccountID  int    `sql:"int"`
	SummonerID int    `sql:"int"`
	JSON       string `sql:"jsonb"`
}
