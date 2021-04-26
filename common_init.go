package common_mapper

/**
  初始化 全局变量 mqtt 定时任务
 */
func CommonInit(ip string, port int, username, password string, projectName string,
	configGet func([]DeviceConfig), subMap map[string]func(topic string, msg []byte))  {
	globalInit(projectName)
	mqttInit(ip, username, password, port, configGet, subMap)
	startAskConfigTask(projectName)
}
