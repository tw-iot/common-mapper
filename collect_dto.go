package common_mapper

type DataCollect struct {
	Gp string                   `json:"gp"`
	Id string                   `json:"id"`
	T  int64                    `json:"t"`
	V  string                   `json:"v"`
	D  []map[string]interface{} `json:"d"`
}
