package monitor

import (

)

var  action_type_group_yser = 1000;

type Message struct {
	
	  Status  int      `json:"status"`
	
	  Msg     string   `json:"msg"`
	  
	  Action   int     `json:"action"`
	  
	  Result  interface{} `json:"result"`
	  
}