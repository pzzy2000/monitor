package monitor

import (
//logger "cn/monitor/log"
	"sync"
	 "encoding/json"
)

var err_msg_find_user_or_group =-100

var success_msg_defults = 0

type UserGroup struct{
	
	 Status int   `json:"status"`
	 
	 Group []GroupBean  `json:"group"`
	 
	 User  []UserBean   `json:"user"`
	
}


func  UserGroupToString() string {
	
	var usergroup UserGroup;
	
	group ,  user , grouperr , usererr  :=ScanGroupAndUser();
	
	if(grouperr !=nil || usererr!=nil) {
		usergroup.Status =err_msg_find_user_or_group
	}else{
		 
		 usergroup.Status =success_msg_defults
		 {
		 var groups[] GroupBean;
	     group.Range(func(k, v interface{}) bool {
	     	 groupBean, _ := v.(GroupBean);
			groups = append(groups, groupBean)
			return true;
		});
	    usergroup.Group = groups;
		 }
		 {
		 	var users[] UserBean;
	     user.Range(func(k, v interface{}) bool {
	     		 userBean, _ := v.(UserBean);
			users = append(users, userBean)
			return true;
		});
	     usergroup.User = users;
		 }
		
	}
	
	    buf, _ := json.Marshal(usergroup)
        return string(buf);
}

func MapTostring(sm sync.Map) (string){
	   var j[] interface{};
	   sm.Range(func(k, v interface{}) bool {
			j = append(j, v)
			return true;
		});
	    buf, _ := json.Marshal(j)
        return string(buf);
	
}

