package listener

import (
	"bytes"
	"fmt"
	"sync"
)

type WorldListener struct {
	Users []*UserListener
}

const (
	ADD_PLAYER int = iota
	REMOVE_PLAYER
)

var instance *WorldListener
var once sync.Once
var mu *sync.Mutex
var count int

func GetWorldInstance() *WorldListener {
	once.Do(func() {
		mu = &sync.Mutex{}
		count = 0
		instance = &WorldListener{make([]*UserListener, 0)}
	})
	return instance
}

func (this *WorldListener) ManipulateUsers(user *UserListener, mode int) {
	mu.Lock()
	defer mu.Unlock()
	if mode == ADD_PLAYER {
		this.addPlayer(user)
	} else if mode == REMOVE_PLAYER {
		this.deletePlayer(user)
	}

}

func (this *WorldListener) addPlayer(user *UserListener) {
	user.ID = count
	count++
	this.Users = append(this.Users, user)
}

func (this *WorldListener) deletePlayer(user *UserListener) {
	for i, val := range this.Users {
		if val == user {
			this.Users = append(this.Users[:i], this.Users[i+1:]...)
			break
		}
	}
}

func (this *WorldListener) CurrentState() []byte {
	buf := bytes.NewBuffer([]byte(""))
	str_fmt := ""
	// format buffer - ID X Y
	var last_index = len(this.Users) - 1
	// this should be reworked later to not be raw values
	for i, val := range this.Users {
		if i != last_index {
			str_fmt = fmt.Sprintf("%d %d %d,", val.ID, val.Player.X, val.Player.Y)
		} else {
			str_fmt = fmt.Sprintf("%d %d %d", val.ID, val.Player.X, val.Player.Y)
		}
		buf.WriteString(str_fmt)
	}
	return buf.Bytes()
}
