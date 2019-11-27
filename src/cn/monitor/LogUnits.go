package monitor

import (
	logger "cn/monitor/log"
	"sync"
	 "encoding/json"
)

func listProcess(sm *sync.Map) {
		logger.Logger("---------------------------------------- ");
		
	sm.Range(func(k, v interface{}) bool {
		process, ok := v.(Process);
		if(ok){
			
			logger.Logger(" Process -> "+(process).Comm);
		 jsonBytes, err := json.Marshal(process)
        if err != nil {
               logger.Error(err.Error())
        }else{
        
        	logger.Logger(" Process -> "+string(jsonBytes));
        }


		}
		
		return true
	})

}
