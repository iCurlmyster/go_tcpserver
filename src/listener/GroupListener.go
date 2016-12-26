package listener

import (
	"bytes"
	"sync"
)

type WorldListener struct {
	Users []*UserListener
}

var instance *WorldListener
var once sync.Once
var mu sync.Mutex
var count int

func GetWorldInstance() *WorldListener {
	once.Do(func() {
		mu = &sync.Mutex{}
		count = 0
		instance = &WorldListener{make([]*UserListener, 0)}
	})
	return instance
}

func (this *WorldListener) AddPlayer(user *UserListener) {
	mu.Lock()
	user.ID = count
	count++
	this.Users = append(this.Users, user)
	mu.Unlock()
}

func (this *WorldListener) CurrentState() []byte {
	buf := bytes.NewBuffer("")
	str_fmt := ""
	// format buffer - ID X Y
	for val, _ := range this.Users {
		str_fmt = fmt.Sprintf("%d %d %d ", val.ID, val.User.X, val.User.Y)
		buf.WriteString(str_fmt)
	}
	return buf.Bytes()
}
