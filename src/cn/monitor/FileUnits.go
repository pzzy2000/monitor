package monitor

import (
	"fmt"
	"io/ioutil"
)
func readfile(file string,process *Process)(data string,err error){
	statPath := fmt.Sprintf("/proc/%s/%s", process.Pid,file);
	dataBytes,err := ioutil.ReadFile(statPath)
	if err != nil {
		return ;
	}else{
	data := string(dataBytes);
	return data,err;
	}
	
}

func readAll(process *Process){
	readComm(process);
	readCmdline(process);
}



func readCmdline(process *Process)(err error){
	 cmdline,err:=readfile("cmdline",process);
    if err==nil{
    	process.Cmdline = TrimUnicode(cmdline);
    }else{
    	process.Comm ="read cmdline file error"
    }
	return ;
}

func readComm(process *Process)(err error){
	
    comm,err:=readfile("comm",process);
    if err==nil{
    	process.Comm = Trim(comm);
    }else{
    	process.Comm ="read comm file error"
    }
	return ;
}