package listener

import (
	//	"bufio"
	"bytes"
	"io"
	"log"
	"net"
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
	defer player_.Conn.Close()
	buff := make([]byte, 8)
	world.ManipulateUsers(player_, ADD_PLAYER)
	for {
		//tmp_buff, err := bufio.NewReader(player_.Conn).ReadString('\n')
		_, err := player_.Conn.Read(buff)
		//buff := []byte(tmp_buff)
		if err != nil {
			if err != io.EOF {
				log.Println("buff err", err)
			}
			log.Println("breaking out", player_)
			break
		}
		//		log.Println("buff:", buff)
		// disconnect if
		if bytes.Contains(buff, []byte("exit")) {
			log.Println("REMOVING")
			world.ManipulateUsers(player_, REMOVE_PLAYER)
			player_.Conn.Write([]byte("exit"))
			break
		}
		// buffer layout should be "x_int y_int"
		if len(buff) > 1 {
			var x_val int = int(buff[0]) | int(buff[1])<<8 | int(buff[2])<<16 | int(buff[3])<<24
			var y_val int = int(buff[4]) | int(buff[5])<<8 | int(buff[6])<<16 | int(buff[7])<<24
			player_.Player.X = x_val
			player_.Player.Y = y_val
		}
		state := world.CurrentState(player_.ID)
		if bytes.Equal(state, []byte("")) {
			player_.Conn.Write([]byte("null"))
		} else {
			player_.Conn.Write(state)
		}
	}
}
