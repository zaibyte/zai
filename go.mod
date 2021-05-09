module g.tesamc.com/IT/zai

go 1.16

require (
	g.tesamc.com/IT/zaipkg v0.0.0-20210221102003-7de4fb06bd59
	g.tesamc.com/IT/zproto v0.0.0-20210114071438-1018be4c55d3
)

// TODO GitLAB proxy issues
replace (
	g.tesamc.com/IT/zaipkg v0.0.0-20210221102003-7de4fb06bd59 => ../zaipkg
	g.tesamc.com/IT/zproto v0.0.0-20210114071438-1018be4c55d3 => ../zproto
)