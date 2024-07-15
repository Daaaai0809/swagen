package constant

const (
	QUERY_PARAM_ID = iota
	HEADER_PARAM_ID
	PATH_PARAM_ID
	COOKIE_PARAM_ID

	QUERY_PARAM  = "query"
	HEADER_PARAM = "header"
	PATH_PARAM   = "path"
	COOKIE_PARAM = "cookie"
)

type ParamLocation struct {
	ID   int
	Type string
}

type ParamLocations map[string]ParamLocation

var ParamLocationsMap = ParamLocations{
	QUERY_PARAM:  ParamLocation{ID: QUERY_PARAM_ID, Type: QUERY_PARAM},
	HEADER_PARAM: ParamLocation{ID: HEADER_PARAM_ID, Type: HEADER_PARAM},
	PATH_PARAM:   ParamLocation{ID: PATH_PARAM_ID, Type: PATH_PARAM},
	COOKIE_PARAM: ParamLocation{ID: COOKIE_PARAM_ID, Type: COOKIE_PARAM},
}

var ParamLocationsList = []string{
	QUERY_PARAM,
	HEADER_PARAM,
	PATH_PARAM,
	COOKIE_PARAM,
}
