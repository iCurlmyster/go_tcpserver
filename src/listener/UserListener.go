package listener

import (
	"log"
	"net"
	"strconv"
)

type UserListener struct {
	Conn   net.Conn
	Player User
	ID     int
}

type User struct {
	X int
	Y int
}

func UserListenerLoop(conn net.Conn) {
	world := GetWorldInstance()
	player_ := &UserListener{Conn: conn}
	defer player.Conn.Close()
	world.AddPlayer(player)
	buff := make([]byte, 1024)
	for {
		readLen, err := player.Conn.Read(buff)
		if err != nil {
			log.Println(err)
		}
		// disconnect if
		if bytes.Contains(buff, []byte("exit")) {
			world.ManipulateUsers(player_, REMOVE_PLAYER)
			break
		}
		// buffer layout should be x_int, y_int
		values := bytes.Split(buff, []byte(" "))
		var count = 0
		for val, _ := range values {
			pos, err := strconv.Atoi(string(val))
			if err != nil {
				log.Println(err)
				break
			}
			if count == 0 {
				player_.Player.X = pos
			} else if count == 1 {
				player_.Player.Y = pos
			}
		}
		player_.Conn.Write(world.CurrentState())
	}
}
