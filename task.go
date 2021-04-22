package common_mapper

import (
	"fmt"
	"github.com/robfig/cron"
)

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
func stopAskConfigTask()  {
	cronAskConfig.Stop()
}

/**
  启动定时任务
 */
func StartCronJob(cronKey string, cycle int64, collect func(), online func()) {
	//设备数据采集
	cronDev := cron.New()
	// 添加定时任务 ms/1000=s
	s := cycle / 1000
	second := fmt.Sprintf("@every %ds", s)
	cronDev.AddFunc(second, collect)
	cronDev.Start()
	cronDevices[cronKey] = cronDev

	//设备在线离线
	cronOnline := cron.New()
	// 添加定时任务 1分钟执行一次
	cronOnline.AddFunc("@every 1m", online)
	cronOnline.Start()
	cronOnlines[cronKey] = cronOnline
}

/**
  停止定时任务
 */
func StopCronJob(key string) {
	c := cronDevices[key]
	if c != nil {
		c.Stop()
	}
	co := cronOnlines[key]
	if co != nil {
		co.Stop()
	}
}

/**
  停止所有定时任务
*/
func StopAllCronJob() {
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
}