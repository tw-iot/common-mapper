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
	Matedata               string `json:"matedata"`
	GatewayProgramConfigId string `json:"gateway_program_config_Id"`
	Version                string `json:"version"`
	Cycle                  int    `json:"cycle"`
	Items                  []Item `json:"items"`
}

type Item struct {
	AuthType        int    `json:"auth_type"`         //采集指标权限，可读，可写，可执行，按linux权限语法设置，r 4, w 2, x 1
	Code            string `json:"code"`              //采集指标编码code，采集属性英文描述
	ConfigItemId    string `json:"config_item_id"`    //采集指标id
	ConfigVersionId string `json:"config_version_id"` //配置版本id
	DataRange       string `json:"data_range"`        //数据取值范围，数值型最小值最大值描述为： 0-100，字符串类型为长度描述：50，表示该字符串最大长度为50
	DataType        string `json:"data_type"`         //数据类型定义，整形还是字符串
	OrderNo         int    `json:"order_no"`          //排序号
	SlaveId         string `json:"slave_id"`          //从站id，默认为1
	Title           string `json:"title"`             //采集属性，中文简述
	UniqueKey       string `json:"unique_key"`        //是否为唯一值标识，用来决定是否根据此值的变化来决定入关系库，0：默认为0，不为唯一标识，1：作值的唯一标识
	QuerySql        string `json:"query_sql"`         //查询sql字段
	Field1          string `json:"field1"`            //扩展字段1
	Field2          string `json:"field2"`            //扩展字段2
	Field3          string `json:"field3"`            //扩展字段3
	Field4          string `json:"field4"`            //扩展字段4
}
