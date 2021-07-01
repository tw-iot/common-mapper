package common_mapper

import mqtt "github.com/eclipse/paho.mqtt.golang"

/**
  初始化 全局变量 mqtt 定时任务
 */
func CommonInit(ip string, port int, username, password string, projectName string,
	configGet func([]DeviceConfig), subMap map[string]func(client mqtt.Client, msg mqtt.Message))  {
	globalInit(projectName)
	mqttInit(ip, username, password, port, configGet, subMap)
	startAskConfigTask(projectName)
}
