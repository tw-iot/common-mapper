package common_mapper

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ammario/temp.v2"
	"log"
	"strings"
	"time"
)

/*
 发送采集的数据
*/
func SendDataMessage(devieId, groupNameEn, version string, time int64, dataMap []map[string]interface{}) {
	//添加label
	lMap := labelMap[devieId]
	for _, dm := range dataMap {
		for k, v := range lMap {
			dm[k] = v
		}
	}

	tenantId := tenantIdMap[devieId]
	index := strings.Index(tenantId, ".")
	group := tenantId[:index]
	area := tenantId[index+1:]

	dataCollect := DataCollect{
		Gp: groupNameEn,
		G:  area,
		Id: devieId,
		T:  time,
		V:  version,
		D:  dataMap,
	}
	jsonBytes, err := json.Marshal(dataCollect)
	if err != nil {
		log.Println("dataCollect json Marshal", err)
		return
	}
	topic := fmt.Sprintf(TopicSendCollectDataMsg, mapperName, mapperName, group, devieId)
	MqttPublish(topic, string(jsonBytes))
}

/*
 发送设备在线离线数据
*/
func SendOnlineMessage(devieId, nodeId string, status bool) {
	s := 0 //设备不在线
	if status {
		s = 1 //在线
	}
	dataOnline := DataOnline{
		DId: devieId,
		GId: nodeId,
		G:   tenantIdMap[devieId],
		S:   s,
		T:   GetTimestamp(),
	}
	jsonBytes, err := json.Marshal(dataOnline)
	if err != nil {
		log.Println("dataOnline json Marshal", err)
		return
	}
	topic := fmt.Sprintf(TopicSendOnlineStateUp, devieId)
	if checkExpireMap(devieId, s) {
		MqttPublish(topic, string(jsonBytes))
	}
}

/*
  如果设备在线离线状态相同,则5分钟发一次
  true:过期,或状态不一致
  false: 没过期
*/
func checkExpireMap(deviceId string, status int) bool {
	flag := false
	//判断map里设备存在
	if od, ok := expireMap[deviceId]; ok {
		//true 已经过期
		if temp.Expired(&od) {
			flag = true
			//一定要从map里删除,不然过期后,每次进这里
			delete(expireMap, deviceId)
			//删除后在添加,不然下一秒进来,又发一次
			putExpireMap(deviceId, status)
		} else {
			//没有过期,比较状态是否一致
			if status != od.status {
				flag = true
				//从map里删除
				delete(expireMap, deviceId)
				putExpireMap(deviceId, status)
			}
		}
	} else {
		//设备不存在
		putExpireMap(deviceId, status)
		flag = true
	}
	return flag
}

/**
  设置过期时间5分钟,并保存到map
*/
func putExpireMap(deviceId string, status int) {
	onlineD := onlineData{status: status}
	//5分钟过期
	temp.ExpireAfter(&onlineD, time.Second*300)
	expireMap[deviceId] = onlineD
}
