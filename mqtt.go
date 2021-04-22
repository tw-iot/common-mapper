package common_mapper

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/tw-iot/mqtt_tw"
	"log"
)

/**
  mqtt 初始化
 */
func mqttInit(ip, username, password string, port int, configGet func([]DeviceConfig)) {
	clientId := uuid.NewV4()
	mqttInfo := mqtt_tw.NewMqttInfo(ip, username,
		password, mapperName + fmt.Sprintf("%s", clientId), port)
	mqtt_tw.MqttInit(&mqttInfo)

	subscribeConfigGet(configGet)
}

/**
  监听 配置解析服务发来的采集配置
 */
func subscribeConfigGet(configGet func([]DeviceConfig)) {
	topic := fmt.Sprintf(TopicMonitorConfigGet, mapperName)
	token := mqtt_tw.MqttTw.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message)  {
		log.Printf("subscribeConfigGet message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		var deviceConfig []DeviceConfig
		err := json.Unmarshal(msg.Payload(), &deviceConfig)
		if err != nil {
			log.Println("json Unmarshal error", err.Error())
			return
		}
		if len(deviceConfig) > 0 {
			stopAskConfigTask()
			// 采集程序成功收到配置后反馈通知
			topicSend := fmt.Sprintf(TopicSendConfigNotify, mapperName)
			MqttPublish(topicSend, "1")
		}
		configCacheMap(deviceConfig)
		configGet(deviceConfig)
	})
	if token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

func MqttPublish(topic, payload string)  {
	mqtt_tw.MqttTw.Publish(topic, 0, false, payload)
}

/**
  关闭mqtt连接
 */
func MqttClose() {
	mqtt_tw.MqttDisconnect()
}
