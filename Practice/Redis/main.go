package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// Command struct used to take in commands for Redis server
type Command struct {
	Fields []string
	Result chan string
}

func redisServer(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {
		if len(cmd.Fields) == 0 {
			continue
		}

		fmt.Println("GOT COMMAND", cmd)

		switch cmd.Fields[0] {

		case "GET":
			if len(cmd.Fields) != 2 {
				cmd.Result <- "Expected 2 Args"
				continue
			}
			k := cmd.Fields[1]
			v := data[k]
			if v == "" {
				cmd.Result <- "No Value"
				continue
			}
			cmd.Result <- v

		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "Expected 3 Args"
				continue
			}
			k := cmd.Fields[1]
			v := cmd.Fields[2]
			data[k] = v
			cmd.Result <- ""

		//case "ALL":
		//	for _, n := range data {
		//		cmd.Result <- n
		//	}

		case "DEL":
			k := cmd.Fields[1]
			delete(data, k)
			cmd.Result <- k + " DELETED"

		default:
			cmd.Result <- "INVALID COMMAND " + cmd.Fields[0]
		}
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		io.WriteString(conn, "=>")
		ln := scanner.Text()
		fs := strings.Fields(ln)

		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}
		io.WriteString(conn, <-result+"\n")
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println("Error Accepting: ", err)
	}
	defer li.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error Connecting", err)

		}
		io.WriteString(conn,
			"**************************************\n"+
				`Very basic redis database server.
		Commands: 
			SET <key> <value>
			GET <key>
			DEL <key>`+
				"\n**************************************\n")
		go handle(commands, conn)
	}
}
