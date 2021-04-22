package common_mapper

/**
  初始化 mqtt 定时任务
 */
func CommonInit(ip, username, password string, port int, projectName string, configGet func([]DeviceConfig))  {
	globalInit(projectName)
	mqttInit(ip, username, password, port, configGet)
	startAskConfigTask(projectName)
}
