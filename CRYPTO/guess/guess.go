package main

import (
	"log"
	"net"
	"os"
	"time"
)

var flag []byte

func getFlag() {
	flag = func() []byte {
		flag, err := os.ReadFile("./flag")
		if err != nil {
			log.Fatal("read flag fail\n", err)
		}
		return flag
	}()
}

func cmpFlag(input []byte, inputLen int) bool {
	for i := 0; i < inputLen; i++ {
		if input[i] != flag[i] {
			return false
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
	return true
}

func game(conn net.Conn) {
	defer conn.Close()
	var input [80]byte

	conn.Write([]byte("我出题 你猜flag 没有中间商赚差价\n\n"))
	for {
		conn.Write([]byte("input flag:\n"))
		inputLen, err := conn.Read(input[:])
		if err != nil {
			return
		}

		if inputLen != 0 {
			if input[inputLen-1] == []byte("\n")[0] {
				input[inputLen-1] ^= []byte("\n")[0]
				inputLen--
			}
		} else {
			return
		}

		if inputLen != len(flag) {
			conn.Write([]byte("wrong length\n\n"))
		} else {
			ok := cmpFlag(input[:], inputLen)
			if ok {
				conn.Write([]byte("success,this is your flag\n"))
				return
			} else {
				conn.Write([]byte("error\n\n"))
				continue
			}
		}
	}
}

func main() {
	getFlag()

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen failed\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept failed\n", err)
			continue
		}

		go game(conn)
	}
}
