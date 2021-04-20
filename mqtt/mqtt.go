package mqtt

import (
	"common-mapper/topic"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/tw-iot/mqtt_tw"
	"log"
)

func MqttInit(projectName string) {
	clientId := uuid.NewV4()
	mqttInfo := mqtt_tw.NewMqttInfo("127.0.0.1", "",
		"", projectName + fmt.Sprintf("%s", clientId), 1883)
	mqtt_tw.MqttInit(&mqttInfo)

	subscribeConfigGet(projectName)
}

/**
  监听 配置解析服务发来的采集配置
 */
func subscribeConfigGet(projectName string) {
	topic := fmt.Sprintf(topic.TopicMonitorConfigGet, projectName)
	token := mqtt_tw.MqttTw.Subscribe(topic, 0, receiveConfigGet)
	if token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

func receiveConfigGet(client mqtt.Client, msg mqtt.Message)  {
	log.Printf("receiveConfigGet message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func MqttClose() {
	mqtt_tw.MqttDisconnect()
}
