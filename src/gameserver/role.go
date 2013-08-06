package main
import (
	"time"
	"sync"
)

type Role struct {
	name string
	id string
	x int
	y int
}

type RoleManager struct{

	roleList []*Role
	maxCount int
	nowCount int
	mutex	sync.Mutex
	cond    *sync.Cond

}

var idMaker int = 0

func NewRoleManager(maxPlayer int) *RoleManager{

	roleManager := new(RoleManager)
	roleManager.maxCount = maxPlayer;
	roleManager.nowCount = 0;
	roleManager.cond = sync.NewCond(&p.mutex)
	return roleManager

}

func NewRole(name string, id string) *Role{

	role := new(Role)
	role.x = 0;
	role.y = 0;
	role.name = name;
	role.id = id;

}

func (this *RoleManager) AddRole(name string) {

	this.mutex.Lock()
	{
		this.nowCount++
		if this.nowCount >= this.maxPlayer{
			return ;
		}
		i := 0
		for {
			if this.roleList[i] == nil {
				this.roleList[i] = NewRole(name, idMaker)
				break
			} else {
				i++
			}
		}
		idMaker++;	
	}
	this.mutex.Unlock()
	this.cond.Broadcast()

}

func (this *RoleManager) RemoveRole(id string) {

	this.mutex.Lock()
	{
		for i := 0; i < this.maxCount; i++ {
			if this.roleList[i] != nil && this.roleList[i].id == id {
				this.roleList[i] = nil
				this.nowCount--;
				break
			} 
		}
	}
	this.mutex.Unlock()

}
 
func (this *RoleManager) GetAllRole(){
	
}