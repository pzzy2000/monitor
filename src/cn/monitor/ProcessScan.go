package monitor

import (
	logger "cn/monitor/log"
	
    "encoding/json"
//	"strconv"
		"fmt"
	"io/ioutil"
	//	"os"
	//	"strconv"
	"sync"
	"time"
)


func endProcess(sm *sync.Map){
	 time :=time.Now();
	sm.Range(func(k, v interface{}) bool {
		  process, ok := v.(Process);
		  if(ok){
		  	 fmt.Println(time.Format("2006-01-02 03:04:05 PM"))
		  	 process.EndTime =time.Format("2006-01-02 03:04:05");
		  }
		
		return true
	})
	
}

func ProcessToJsonString() string {
	var msg Message = Message{Action: action_type_process}
	msg.Status = success_msg_defults
	 sm := Scan();
	 var j[] interface{};
	 sm.Range(func(k, v interface{}) bool {
			j = append(j, v)
			return true;
		});
	 msg.Result = j;
	 buf, _ := json.Marshal(msg)
	return string(buf)

}

func Scan() sync.Map {
	logger.Logger("start sacn process \r\n")
	var sm sync.Map
	endProcess(&sm);
	list(&sm);
	return sm;
	
}


func list(sm *sync.Map){
	
	files, err := ioutil.ReadDir("/proc")
	
	if err == nil {
		for _, f := range files {
			if f.IsDir() && IsDigit(f.Name()) {
				v,ok := sm.Load(f.Name());
				if(ok){
					process := v.(Process);
					process.EndTime = "";
					readAll(&process);
//				    sm.Store(f.Name(),process)
				}else{
					var process = Process{Pid: f.Name(),StartTime:f.ModTime().Format("2006-01-02 03:04:05")};
				    readAll(&process);
				    sm.Store(f.Name(),process)
				}
				
			}

		}
	}
 
}
