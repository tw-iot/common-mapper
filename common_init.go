package common_mapper

/**
  初始化 全局变量 mqtt 定时任务
 */
func CommonInit(ip string, port int, username, password string, projectName string,
	configGet func([]DeviceConfig), subTopics []string, subFuns []func(topic string, msg []byte))  {
	globalInit(projectName)
	mqttInit(ip, username, password, port, configGet, subTopics, subFuns)
	startAskConfigTask(projectName)
}
