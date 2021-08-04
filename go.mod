module g.tesamc.com/IT/zai

go 1.16

require (
	g.tesamc.com/IT/zaipkg v0.0.0
	g.tesamc.com/IT/keeper v0.0.0
	g.tesamc.com/IT/zproto v0.0.0
	github.com/jaypipes/ghw v0.8.0
)

// TODO GitLAB proxy issues
replace (
	g.tesamc.com/IT/keeper v0.0.0 => ../keeper
	g.tesamc.com/IT/zaipkg v0.0.0 => ../zaipkg
	g.tesamc.com/IT/zproto v0.0.0 => ../zproto
)