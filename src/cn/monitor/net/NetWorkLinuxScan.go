package net

import (
	logger "cn/monitor/log"
	toolsUnits "cn/monitor"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"io/ioutil"
	"strings"
)

func Sacn(sm *sync.Map) {
	Scan()
	//	sm.Range(func(k, v interface{}) bool {
	//		  process, _ := v.(processUnits.Process);
	//		 Scan(&process);
	//		 return true;
	//});
}

func Scan() {
	socketPid ,err:= scanSocket()
	
	if(err!=nil){
		
	}else{
	 var pidSocket map[string]NetWorks
	 
	 scanTcp(&pidSocket,&socketPid)	
	}
	
	
	
}

func scanSocket() (map[string]string,error) {
	 socketPid := make(map[string]string)
	files, err := ioutil.ReadDir("/proc")
	if err == nil {
		for _, f := range files {
			if f.IsDir() && toolsUnits.IsDigit(f.Name()) {
			statPath := fmt.Sprintf("/proc/%s/%s", f.Name(), "fd")
	        fdChilds, err_ := ioutil.ReadDir(statPath);
	        if(err_ ==nil){
	        	for _, fdChild := range fdChilds {
	        		path := fmt.Sprintf("/proc/%s/%s/%s", f.Name(), "fd",fdChild.Name())
	        		linkc ,err :=  os.Readlink(path);
	        		  if(err==nil && strings.HasPrefix(linkc,"socket:")){
	        		   linkc =	linkc[ 8:strings.LastIndex(linkc, "]")];
	        		   socketPid[linkc]=f.Name();
	        		  }
	        	}
	          }
			}

		}
	return socketPid,nil;
	}else{
		return socketPid,err
	}
	
	
}

func scanTcp(pidSocket *map[string]NetWorks ,socketPid *map[string]string ) (err error) {
    
	f, err := os.Open("/proc/net/tcp")
	if err != nil {
		return err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	sum := 0
	for {

		lineByte, _, err := bfRd.ReadLine()
		if sum == 0 {
			sum++
			continue
		}
		logger.Logger(string(lineByte))
		line := string(lineByte)
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return err
		}
		st, _ := strconv.ParseInt(line[34:36], 16, 8)

		switch st {
		case 10:
			{
				
				lines:=strings.Fields(line);
				
				int1, _ := strconv.ParseInt(line[6:8], 16, 8)
				//
				int2, _ := strconv.ParseInt(line[8:10], 16, 8)

				int3, _ := strconv.ParseInt(line[10:12], 16, 8)

				int4, _ := strconv.ParseInt(line[12:14], 16, 8)

				port, _ := strconv.ParseInt(line[15:19], 16, 16)

				localAddress := fmt.Sprintf("%d.%d.%d.%d:%d", int1, int2, int3, int4, port)
				
//				fmt.Printf("Fields are: %q", lines)
				
				logger.Logger(localAddress+"    >   " +lines[9]  );
//                tcp = append(tcp, localAddress)
//								logger.Logger(localAddress)
			}
		}

	}
	return nil
}
