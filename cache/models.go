package cache

import (
	"time"
)

// CachedParams is a interface that is send with the call.
// This will make all caching happen the same for all calls, so that way any changes with the cache can be changed in one place.
type CachedParams struct {
	Cached     bool
	Expire     bool
	Expiration time.Duration
	CallKey    string
	CallType   string
}
