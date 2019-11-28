package monitor

import (
	logger "cn/monitor/log"

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


func Scan() {
	logger.Logger("start sacn process \r\n")
	var sm sync.Map
	endProcess(&sm);

	list(&sm);
	
//	listProcess(&sm);
}

func list(sm *sync.Map){
	
	files, err := ioutil.ReadDir("/proc")
	
	if err == nil {
		for _, f := range files {
			if f.IsDir() && IsDigit(f.Name()) {
				var process = Process{Pid: f.Name(),StartTime:f.ModTime().Format("2006-01-02 03:04:05")};
				readAll(&process);
				sm.Store(f.Name(),process)
//				 err :=readAll(&process);
//				  if(err!=nil){
//				  	logger.Error("read comm file error "+err.Error());
//				  }else{
//				  	sm.Store(f.Name(),process)
//				  }	
			}

		}
	}
 
}
