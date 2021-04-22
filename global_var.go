package common_mapper

import "github.com/robfig/cron"

var (
	mapperName string

	//key=设备id,value=tenantId
	tenantIdMap map[string]string
	//key=设备id,value=labels
	labelMap map[string]map[string]interface{}

	//采集程序启动要配置 定时任务
	 cronAskConfig *cron.Cron
	//设备读取数据定时任务
	 cronDevices map[string]*cron.Cron
	//设备在线定时任务
	 cronOnlines map[string]*cron.Cron
)

func globalInit(projectName string)  {
	mapperName = projectName
	tenantIdMap = make(map[string]string)
	labelMap = make(map[string]map[string]interface{})
	cronDevices = make(map[string]*cron.Cron)
	cronOnlines = make(map[string]*cron.Cron)
}

func configCacheMap(deviceConfig []DeviceConfig) {
	//先清空
	tenantIdMap = make(map[string]string)
	labelMap = make(map[string]map[string]interface{})
	for _, dc := range deviceConfig {
		for _, dev := range dc.TwDevices {
			tenantIdMap[dev.Id] = dev.TenantId
			labelMap[dev.Id] = dev.Labels
		}
	}
}
