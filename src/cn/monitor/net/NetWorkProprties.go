package net

import (

)


var msg_success =0;

var msg_fail =1;

var socket_pid_defult_nil ="-100";


type NetWorks struct{
	  
	  Pid int   `json:"pid"`
	  // key  localip
	  NetWork []NetWork  `json:"netWork"`
	  
}


type NetWork struct{
	
	 Proto  int  `json:"proto"`
	 
	 LocalAddress  string `json:"localAddress"`
	 
	 
}