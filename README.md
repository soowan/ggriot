# ggriot
ggriot is a library that is easy to use and lets you connect to Riot's API.

# Caching
This is a pretty basic library but it does make caching call's extreamly easy. ggriot uses Postgres (gorm), and is super easy to setup

```
import (
	"github.com/soowan/ggriot/cache"
)

func init() {
	cache.UseCache("host=myhost port=myport user=gorm dbname=gorm password=mypassword")
}

```

and that's it. Everything else will be handled by ggriot.

# MMR
ggriot also uses whatsmymmr to return estimated MMR, please do note this isn't super accurate, and doesn't always work. This may be removed in the future.
