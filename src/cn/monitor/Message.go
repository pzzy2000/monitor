package monitor

import (

)

var  action_type_group_user = 1000;

var  action_type_process = 1001;

type Message struct {
	
	  Status  int      `json:"status"`
	
	  Msg     string   `json:"msg"`
	  
	  Action   int     `json:"action"`
	  
	  Result  interface{} `json:"result"`
	  
}