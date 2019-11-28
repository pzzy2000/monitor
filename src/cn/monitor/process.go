package monitor

import ()

type Process struct {
	User string `json:"user"`

	Pid string `json:"pid"`

	Ppid string `json:"ppid"`
 
    //线程组号
    Pgid string `json:"pgid"`

    Sid string `json:"sid"`
	//task_state=R 任务的状态，
//	R:runnign, 
//	S:sleeping (TASK_INTERRUPTIBLE),
//	D:disk sleep (TASK_UNINTERRUPTIBLE), 
//	T: stopped, T:tracing
	State string `json:"state"`
	
	Comm string `json:"comm"`

	Cmdline string `json:"cmdline"`

	EndTime string `json:"endTime"`

	StartTime string `json:"startTime"`
}
