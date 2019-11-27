package monitor

import (
	logger "cn/monitor/log"

//	"strconv"
//		"fmt"
	"io/ioutil"
	//	"os"
	//	"strconv"
	"sync"
)

func Scan() {
	logger.Logger("start sacn process \r\n")
	var sm sync.Map
	//	vv, ok := sm.Load(1)
	//	if !ok {
	//		vv = Process{Pid: 11}
	//		sm.Store(1, vv)
	//	}
	//	p, ok := (vv).(Process)

	list(&sm);
	
	listProcess(&sm);
}

func list(sm *sync.Map){
	
	files, err := ioutil.ReadDir("/proc")
	
	if err == nil {
		for _, f := range files {
			if f.IsDir() && IsDigit(f.Name()) {
				var process = Process{Pid: f.Name()};
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
