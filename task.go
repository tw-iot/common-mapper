package common_mapper

import (
	"fmt"
	"github.com/robfig/cron"
)

type FuncCollect func() Collect

func (f FuncCollect) collectDev() {
	//得到定时任务结果
	collect := f()
	//发送数据mqtt
	SendDataMessage(collect.DevieId, collect.GroupNameEn, collect.Version, GetTimestamp(), collect.DataMap)
}

type FuncOnline func() Online

func (f FuncOnline) onlineDev() {
	online := f()
	//发送在线状态mqtt
	SendOnlineMessage(online.DevieId, online.NodeId, online.Status)
}

/**
  发送 采集程序询问采集配置 定时任务
*/
func startAskConfigTask(projectName string) {
	topic := fmt.Sprintf(TopicSendConfigAsk, projectName)
	cronAskConfig = cron.New()
	// 添加定时任务, 5秒执行一次, 获取配置成功后,需要停止定时任务
	cronAskConfig.AddFunc("*/5 * * * * * ", func() {
		MqttPublish(topic, "1")
	})
	cronAskConfig.Start()
}

/**
  停止 采集程序询问采集配置 定时任务
*/
func stopAskConfigTask() {
	cronAskConfig.Stop()
}

/**
  启动定时任务 设备数据采集定时任务 设备在线离线定时任务
*/
func StartAllCron(cronKey string, cycle int64, collectF FuncCollect, onlineF FuncOnline) {
	//设备数据采集
	cronDev := cron.New()
	// 添加定时任务 ms/1000=s
	s := cycle / 1000
	second := fmt.Sprintf("@every %ds", s)
	cronDev.AddFunc(second, collectF.collectDev)
	cronDev.Start()
	cronDevices[cronKey] = cronDev

	//设备在线离线
	cronOnline := cron.New()
	// 添加定时任务 1分钟执行一次
	cronOnline.AddFunc("@every 1m", onlineF.onlineDev)
	cronOnline.Start()
	cronOnlines[cronKey] = cronOnline
}

/**
  停止定时任务
*/
func stopCron(cronKey string) {
	c := cronDevices[cronKey]
	if c != nil {
		c.Stop()
	}
	co := cronOnlines[cronKey]
	if co != nil {
		co.Stop()
	}
	//删除key
	delete(cronDevices, cronKey)
	delete(cronOnlines, cronKey)
}

/**
  停止所有定时任务
*/
func stopAllCron() {
	for _, cd := range cronDevices {
		if cd != nil {
			cd.Stop()
		}
	}
	for _, co := range cronOnlines {
		if co != nil {
			co.Stop()
		}
	}
	//清空map
	cronDevices = make(map[string]*cron.Cron)
	cronOnlines = make(map[string]*cron.Cron)
}
