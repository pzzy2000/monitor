package monitor

import ()

// 用户名:密码:ID:用户组ID:注释:家目录:登录使用的shell

var err_msg_find_user_or_group =-100

var success_msg_defults = 0

type UserGroup struct{
	
	 Group []GroupBean  `json:"group"`
	 
	 User  []UserBean   `json:"user"`
	
}


type UserBean struct {
	Name string `json:"name"`

	Passwd bool `json:"pwd"`

	UID int `json:"uid"`

	GID int `json:"gid"`

	Home string `json:"home"`

	LoginShell string `json:"loginShell"`
}

//用户组root，x是密码段，表示没有设置密码，GID是0,root用户组下包括root、linuxsir以及GID为0的其它用户。
type GroupBean struct {
	Name   string `json:"name"`
//	Passwd bool   `json:"pwd"`
	GID    int    `json:"gid"`
}
