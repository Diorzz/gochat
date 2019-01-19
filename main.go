package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"gochat/network"
)

func main() {
	connManager := network.NewConnManager("192.168.0.101:8989")

	connManager.Start()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		cmd, content, err := network.ReadComman(input)
		if err != nil {
			log.Println(err)
			continue
		}

		current := time.Now()
		msg := &network.Message{
			MsgText:  content,
			SendTime: fmt.Sprintf("yyyy-mm-dd HH:mm:ss : ", current.Format("2006-01-02 15:04:05")),
		}

		switch cmd {
		case "chat":
			fmt.Println(content)
		case "bd":
			connManager.BroadCast(*msg)
		default:
			fmt.Println(content)
		}
	}
}
