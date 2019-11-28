package monitor

import (
	logger "cn/monitor/log"
	"sync"
	 "encoding/json"
)

func ListJson(sm sync.Map) {
	sm.Range(func(k, v interface{}) bool {
		
		 jsonBytes, err := json.Marshal(v)
        if err != nil {
               logger.Error(err.Error())
        }else{
        
        	logger.Logger(" Process -> "+string(jsonBytes));
        }
		return true
	})

}

func listProcess(sm *sync.Map) {
		logger.Logger("---------------------------------------- ");
		
	sm.Range(func(k, v interface{}) bool {
	
		
		 jsonBytes, err := json.Marshal(v)
        if err != nil {
               logger.Error(err.Error())
        }else{
        
        	logger.Logger(" Process -> "+string(jsonBytes));
        }

		return true
	})

}
