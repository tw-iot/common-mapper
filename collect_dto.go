package common_mapper

type DataCollect struct {
	G  string                   `json:"g"`
	Id string                   `json:"id"`
	T  int64                    `json:"t"`
	V  int                      `json:"v"`
	D  []map[string]interface{} `json:"d"`
}
