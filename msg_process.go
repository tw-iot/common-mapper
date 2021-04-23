package common_mapper

import (
	"encoding/json"
	"fmt"
	"log"
)

func SendDataMessage(devieId, groupNameEn, version string, time int64, dataMap []map[string]interface{}) {
	//添加label
	lMap := labelMap[devieId]
	for _, dm := range dataMap {
		for k, v := range lMap {
			dm[k] = v
		}
	}
	dataCollect := DataCollect{
		Gp:  groupNameEn,
		Id: devieId,
		T:  time,
		V:  version,
		D:  dataMap,
	}
	jsonBytes, err := json.Marshal(dataCollect)
	if err != nil {
		log.Println(err)
		return
	}
	topic := fmt.Sprintf(TopicSendCollectDataMsg, mapperName, tenantIdMap[devieId], devieId)
	MqttPublish(topic, string(jsonBytes))
}

func SendOnlineMessage(devieId, msg string) {
	topic := fmt.Sprintf(TopicSendOnlineStateUp, devieId)
	MqttPublish(topic, msg)
}
