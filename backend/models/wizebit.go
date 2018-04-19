package models

import (
	"time"
)

type Users struct {
	Id         int    `orm:"pk;column(id);auto"`
	PrivateKey string `orm:"column(private_key);unique"`
	PublicKey  string `orm:"column(public_key);unique"`
	Address    string `orm:"column(address);unique"`
	Password   string
	Status     bool
	Role       int //role of user - see const below
	Rate       int
	CreatedAt  time.Time  `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt  time.Time  `orm:"column(updated_at);type(timestamp);auto_now"`
	SessionKey string     `orm:"column(session_key)"`
	Servers    []*Servers `orm:"reverse(many)"`
}

const (
	USER_SUPERADMIN = 0
	USER_MANAGER    = 10
	USER_CUSTOMER   = 20
	USER_GUEST      = 30
)

type BugReports struct {
	Id        int `orm:"pk;column(id);auto"`
	UserId    int `orm:"column(user_id)"`
	Message   string
	Picture   string
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

type Servers struct {
	Id        int       `orm:"pk;column(id);auto"`
	User      *Users    `orm:"rel(fk);column(user_id)"`
	Name      string    //Unique id of server, maybe address of init wallet //TODO: придумать это
	Url       string    // Address of server maybe node1.wizebit.com for example
	Role      string    // Role of server - Blockchain, Raft, Storage
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(timestamp);auto_now"`
}

type ServerState struct {
	Id          int      `orm:"pk;column(id);auto"`
	ServerId    *Servers `orm:"rel(one)"`
	Ip          string   // IP of server, can be different, must monitoring this
	Status      bool     // up/down - true/false
	Latency     int      // in ms by ping?
	FreeStorage int      // in MB
	Uptime      int      // in sec from server goroutine
	TypeActive  string   // out/in for different type of monitoring -active/passive
	Rate        int      // calculated rate of server in moment
	// if status = false {Rate = 0}
	// else Rate = 0,2*FreeStorage/max.FreeStorage + 0,3*Uptime/max.Uptime +
	// + 0,1*min.Latency/Latency + TypeActive*0,4
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);auto_now_add"`
}

type ServerList struct {
	Id          int
	UserId      int
	SId         int
	Ip          string
	Status      bool
	Latency     int
	FreeStorage int
	Uptime      int
	Rate        int
	CreatedAt   time.Time
}

type ServerStateCount struct {
	Id                   int
	TotalBlockchainCount int
	TotalRaftCount       int
	TotalStorageCount    int
	TotalSuspiciosCount  int
}
