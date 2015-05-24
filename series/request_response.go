package main

// The request and response types should be annotated sufficiently for all
// transports we intend to use.

type listrequest struct {
}

type listresponse struct {
	series []Series `json:"series"`
}

type getrequest struct {
	id string
}

type getresponse struct {
	Series `json:"series"`
}
