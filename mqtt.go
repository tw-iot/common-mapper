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
	willTopic := fmt.Sprintf(TopicWillOnlineStateUp, mapperName)
	willMsg := "0" //遗言为离线状态
	mqttInfo := mqtt_tw.NewMqttInfoWill(ip, username,
		password, mapperName + fmt.Sprintf("%s", clientId), willTopic, willMsg, port)

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		//连接成功的回调
		log.Println("mqtt connected success")
		//订阅放在这里,断开后重新连接时,重新订阅
		subscribeConfigGet(configGet)
	}
	mqtt_tw.MqttInit(&mqttInfo, messagePubHandler, connectHandler, connectLostHandler)
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
			log.Println("json Unmarshal error:", err.Error())
			return
		}
		if len(deviceConfig) > 0 {
			stopAskConfigTask()
			// 采集程序成功收到配置后反馈通知
			topicSend := fmt.Sprintf(TopicSendConfigNotify, mapperName)
			MqttPublish(topicSend, "1")

			configCacheMap(deviceConfig)
			stopAllCron()
			configGet(deviceConfig)
		}
	})
	if token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

func MqttPublish(topic, payload string)  {
	mqtt_tw.MqttTw.Publish(topic, 0, false, payload)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//全局 MQTT pub 消息处理
	//log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	//连接丢失的回调
	log.Println("mqtt connect lost err: ", err)
}

/**
  关闭mqtt连接
 */
func MqttClose() {
	mqtt_tw.MqttDisconnect()
}
