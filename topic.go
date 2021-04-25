package common_mapper

const (
	//----------发送--------------

	//发送 设备在线状态
	TopicSendOnlineStateUp = "$tw/device/%s/edge/online_state/up"

	//发送 采集程序询问采集配置
	TopicSendConfigAsk = "$tw/mapper/%s/config/ask"

	//发送 采集数据 上报消息
	TopicSendCollectDataMsg = "$tw/msg/pod/%s/tenant/%s/device/%s/twin/up"

	//发送 采集程序收到配置后反馈通知
	TopicSendConfigNotify = "$tw/mapper/%s/config/notify"

	//发送 采集程序离线遗言
	TopicWillOnlineStateUp = "$tw/will/pod/%s/online_state/up"

	//---------监听--------------

	//监听 采集程序获取采集配置
	TopicMonitorConfigGet = "$tw/mapper/%s/config/get"
)
