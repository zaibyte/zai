module g.tesamc.com/IT/zai

go 1.16

require (
	g.tesamc.com/IT/keeper v0.0.0
	g.tesamc.com/IT/zaipkg v0.0.0
	g.tesamc.com/IT/zproto v0.0.0
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/go-redis/redis/v8 v8.11.2
	github.com/jaypipes/ghw v0.8.0
	google.golang.org/appengine v1.4.0
)

// TODO GitLAB proxy issues
replace (
	g.tesamc.com/IT/keeper v0.0.0 => ../keeper
	g.tesamc.com/IT/zaipkg v0.0.0 => ../zaipkg
	g.tesamc.com/IT/zproto v0.0.0 => ../zproto
)
