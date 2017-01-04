package listener

import (
	"bytes"
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

func (this *WorldListener) CurrentState(id int) []byte {
	buf := bytes.NewBuffer([]byte(""))
	// format buffer - ID X Y
	for _, val := range this.Users {
		if id != val.ID {
			buf.WriteByte(byte(val.ID))
			buf.WriteByte(byte(val.ID >> 8))
			buf.WriteByte(byte(val.ID >> 16))
			buf.WriteByte(byte(val.ID >> 24))
			buf.WriteByte(byte(val.Player.X))
			buf.WriteByte(byte(val.Player.X >> 8))
			buf.WriteByte(byte(val.Player.X >> 16))
			buf.WriteByte(byte(val.Player.X >> 24))
			buf.WriteByte(byte(val.Player.Y))
			buf.WriteByte(byte(val.Player.Y >> 8))
			buf.WriteByte(byte(val.Player.Y >> 16))
			buf.WriteByte(byte(val.Player.Y >> 24))
		}
	}
	return buf.Bytes()
}
