package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"chat_v3.2/network"
)

func main() {
	connManager := &network.ConnMgr{
		LocalIP: "192.168.0.101:8989",
	}

	connManager.Start()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n') //读取器对象提供一个方法 ReadString(delim byte) ，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。
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
