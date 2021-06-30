package common_mapper

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	uuid "github.com/satori/go.uuid"
	"github.com/tw-iot/mqtt_tw"
	"log"
	"reflect"
	"regexp"
	"strings"
)

/**
  mqtt 初始化
 */
func mqttInit(ip, username, password string, port int, configGet func([]DeviceConfig),
	subMap map[string]func(topic string, msg []byte)) {
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
		subCustomize(subMap)
	}
	mqtt_tw.MqttInit(&mqttInfo, messagePubHandler, connectHandler, connectLostHandler)
}

/**
  自定义订阅
 */
func subCustomize(subMap map[string]func(topic string, msg []byte))  {
	if subMap == nil || len(subMap) == 0 {
		return
	}

	//保存特殊订阅的topic 正则表达式
	for topic, _ := range subMap {
		if strings.Index(topic, "$") == 0 {
			//如果topic是$开头
			topic = strings.Replace(topic, "$", "\\\\$", -1 )
		}
		if strings.Contains(topic, "#") {
			// topic= abc/#
			str := strings.Replace(topic, "#", "(.*)", -1 )
			// abc/(.*)
			//^abc/(.*) 正则表达式 ^以什么开头 $以什么结尾
			netStr := fmt.Sprintf("%s%s", "^", str)
			mqttHashTagTopicMap[netStr] = topic
		}
		if strings.Contains(topic, "+") {
			// topic= abc/+/123/+/456
			netStr := strings.Replace(topic, "+", "(\\\\w)+", -1 )
			if strings.Index(topic, "+") != 0 {
				//如果+号不是第一个字符
				netStr = fmt.Sprintf("%s%s", "^", netStr)
			}
			if strings.LastIndex(topic, "+") != len(topic) -1 {
				//如果+号不是最后一个字符
				netStr = fmt.Sprintf("%s%s", netStr, "$")
			}
			//^abc/(\w)+/123/(\w)+/456$ 正则表达式 ^以什么开头 $以什么结尾
			mqttPlusTagTopicMap[netStr] = topic
		}
	}

	for topic, subFun := range subMap {
		token := mqtt_tw.MqttTw.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message)  {
			funType := reflect.TypeOf(subFun)
			funName := funType.Name()
			log.Printf("subCustomize %s message: %s from topic: %s\n", funName, msg.Payload(), msg.Topic())
			// 经测试不能直接调用, 会覆盖上一个订阅topic的方法,导致所有订阅的topic,只会发到最后一个方法
			//subFun(msg.Topic(), msg.Payload())
			// 这种方式,可以准确调用对应方法
			//如果订阅的topic带aaa/#号,来的topic是aaa/123,则map找不到对应key,报空指针
			if _, ok := subMap[msg.Topic()]; ok {
				//全匹配
				subMap[msg.Topic()](msg.Topic(), msg.Payload())
			} else {
				//#号匹配
				for reg, topicHashtag := range mqttHashTagTopicMap {
					match, _ := regexp.MatchString(reg, msg.Topic())
					if match {
						subMap[topicHashtag](msg.Topic(), msg.Payload())
					}
				}
				//+号匹配
				for reg, topicPlustag := range mqttPlusTagTopicMap {
					match, _ := regexp.MatchString(reg, msg.Topic())
					if match {
						subMap[topicPlustag](msg.Topic(), msg.Payload())
					}
				}
			}
		})
		if token.Wait() && token.Error() != nil {
			log.Println("subCustomize token err: ", token.Error())
		}
	}
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
		log.Println("subscribeConfigGet token err: ", token.Error())
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
