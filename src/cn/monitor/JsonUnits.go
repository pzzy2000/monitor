package monitor

import (
//logger "cn/monitor/log"
	"sync"
	 "encoding/json"
)


func SyncMapToJsonString(sm sync.Map) (string){
	   var j[] interface{};
	   sm.Range(func(k, v interface{}) bool {
			j = append(j, v)
			return true;
		});
	    buf, _ := json.Marshal(j)
        return string(buf);
	
}

