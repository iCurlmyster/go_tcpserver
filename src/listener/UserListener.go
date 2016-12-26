package listener

import (
	"log"
	"net"
)

type UserListener struct {
	Conn   net.Conn
	player User
	ID     int
}

type User struct {
	X int
	Y int
}

func UserListenerLoop(conn net.Conn) {
	world := GetWorldInstance()
	player := &UserListener{Conn: conn}
	defer player.Conn.Close()
	world.AddPlayer(player)
	buff := make([]byte, 1024)
	for {
		readLen, err := player.Conn.Read(buff)
		if err != nil {
			log.Println(err)
		}
		// TODO Handle buffer logic
		player.Conn.Write(world.CurrentState())
	}
}
