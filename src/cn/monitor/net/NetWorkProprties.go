package net

import (

)

var tcp_listen = 10

var msg_success =0;

var msg_fail =1;


type NetWorks struct{
	  
	  msg int `json:"msg"`
	 
	  Pid int   `json:"pid"`
	 
	  Tcp  []NetWork `json:"tcp"`
	  
	  Udp  []NetWork `json:"udp"`
}


type NetWork struct{
	
	 LocalAddress []string  `json:"local"`
	 
}