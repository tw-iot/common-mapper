package common_mapper

type DataCollect struct {
	Gp string                   `json:"gp"`
	Id string                   `json:"id"`
	T  int64                    `json:"t"`
	V  string                   `json:"v"`
	D  []map[string]interface{} `json:"d"`
}

type DataOnline struct {
	DId string `json:"dId"`
	GId string `json:"gId"`
	G   string `json:"g"`
	S   int    `json:"s"`
	T   int64  `json:"t"`
}

type Collect struct {
	DevieId     string
	GroupNameEn string
	Version     string
	DataMap     []map[string]interface{}
}

type Online struct {
	DevieId string
	NodeId  string
	Status  bool
}
