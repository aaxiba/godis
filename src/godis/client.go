package godis

type GodClient struct {
	Fd       int          `json:"fd"`
	Name     string       `json:"name"`
	QueryBuf string       `json:"query_buf"`
	Argv     []*GodObject `json:"argv"`
	Argc     int          `json:"argc"`
	Flags    int          `json:"flags"`
	Buf      string       `json:"buf"`
	Reply    string       `json:"reply"`
}

