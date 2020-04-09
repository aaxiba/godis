package godis

import "container/list"

type GodServer struct {
	Pid      int                    `json:"pid"`
	Port     int32                  `json:"port"`
	DbNum    DbNumT                 `json:"db_num"`
	Db       *GodDb                 `json:"db"`
	Clients  list.List              `json:"clients"`
	Commands map[string]*GodCommand `json:"commands"`
}
