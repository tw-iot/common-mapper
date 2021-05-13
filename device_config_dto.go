package common_mapper

type DeviceConfig struct {
	TwDevices []TwDevice `json:"devices"`
	TwConfigs []TwConfig `json:"configs"`
}

type TwDevice struct {
	Id           string                 `json:"id"`
	TenantId     string                 `json:"tenant_id"`
	Ip           string                 `json:"ip"`
	Port         int                    `json:"port"`
	CollectCycle int64                  `json:"collect_cycle"`
	ReportCycle  int64                  `json:"report_cycle"`
	Matedata     string                 `json:"matedata"`
	Labels       map[string]interface{} `json:"labels"`
}

type TwConfig struct {
	DeviceGroupId          string `json:"device_group_id"`
	GroupNameEn            string `json:"group_name_en"`
	GatewayProgramConfigId string `json:"gateway_program_config_Id"`
	Version                string `json:"version"`
	Cycle                  int    `json:"cycle"`
	Items                  []Item `json:"items"`
}

type Item struct {
	AuthType        int                    `json:"auth_type"`
	Code            string                 `json:"code"`
	ConfigItemId    string                 `json:"config_item_id"`
	ConfigVersionId string                 `json:"config_version_id"`
	DataRange       string                 `json:"data_range"`
	DataType        string                 `json:"data_type"`
	OrderNo         int                    `json:"order_no"`
	SlaveId         string                 `json:"slave_id"`
	Title           string                 `json:"title"`
	UniqueKey       string                 `json:"unique_key"`
	QuerySql        string                 `json:"query_sql"`
	Field2          string                 `json:"field2"`
	Field3          string                 `json:"field3"`
	Field4          string                 `json:"field4"`
}
