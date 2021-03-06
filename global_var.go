package common_mapper

import (
	"github.com/robfig/cron"
	"time"
)

var (
	mapperName string

	//key=设备id,value=tenantId
	tenantIdMap map[string]string
	//key=设备id,value=labels
	labelMap map[string]map[string]interface{}
	//过期map,key=设备id,value=设备在线离线状态
	expireMap map[string]onlineData

	//mqtt监听的topic 带#号
	mqttHashTagTopicMap map[string]string
	//mqtt监听的topic 带+号
	mqttPlusTagTopicMap map[string]string

	//采集程序启动要配置 定时任务
	cronAskConfig *cron.Cron
	//设备读取数据定时任务
	cronDevices map[string]*cron.Cron
	//设备在线定时任务
	cronOnlines map[string]*cron.Cron
	//设备读取数据定时任务 毫秒级定时
	tickerDevices map[string]*time.Ticker
)

func globalInit(projectName string)  {
	mapperName = projectName

	tenantIdMap = make(map[string]string)
	labelMap = make(map[string]map[string]interface{})
	expireMap = make(map[string]onlineData)

	mqttHashTagTopicMap = make(map[string]string)
	mqttPlusTagTopicMap = make(map[string]string)

	cronDevices = make(map[string]*cron.Cron)
	cronOnlines = make(map[string]*cron.Cron)
	tickerDevices = make(map[string]*time.Ticker)
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
