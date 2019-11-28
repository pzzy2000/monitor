package monitor

import (
	"bufio"
	logger "cn/monitor/log"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ScanGroupAndUserToString() (json string, err error ) {

	logger.Logger("start sacn process \r\n")

	group, grouperr = scanGroup()
	
	user, usererr = ScanUser()
	
	if(grouperr!=nil || usererr!=nil){
		return nil;
	}

	
	

}

func ScanGroupAndUser() (group ,user sync.Map, grouperr ,usererr error ) {

	logger.Logger("start sacn process \r\n")

	group, grouperr = scanGroup()
	
	user, usererr = ScanUser()

	return group, user ,grouperr ,usererr

}
//name:password:uid:gid:comment:home:shell
func ScanUser() (sync.Map, error) {
	var gp sync.Map
	f, err := os.Open("/etc/passwd")
	if err != nil {
		return gp, err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, _, err := bfRd.ReadLine()
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return gp, err
		}
		data := strings.Split(string(line), ":")	
//		logger.Logger(string(line))	
//		logger.Logger(string(line)+"  " + data[4])	
//		name:password:uid:gid:comment:home:shell
        uid, _ := strconv.Atoi(data[2])
        gid, _ := strconv.Atoi(data[3])
		gp.Store(gid, UserBean{Name:data[0],Passwd:true,UID:uid ,GID:gid , Home :data[5] , LoginShell :data[6] })
	}
	return gp, nil
}

func scanGroup() (sync.Map, error) {
	var gp sync.Map
	f, err := os.Open("/etc/group")
	if err != nil {
		return gp, err
	}
	defer f.Close()
	bfRd := bufio.NewReader(f)
	for {
		line, _, err := bfRd.ReadLine()
		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				break
			}
			return gp, err
		}

		data := strings.Split(string(line), ":")
		
		gid, _ := strconv.Atoi(data[2])
		
		gp.Store(gid, GroupBean{Name:data[0],GID:gid})
	}
	return gp, nil
}
