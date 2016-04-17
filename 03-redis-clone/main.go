package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//Command is a channel and the fields
type Command struct {
	Fields []string
	Result chan string
}

var data = make(map[string]string)

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(commands, conn)
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}

		fmt.Fprintln(conn, <-result)
	}
}

func redisServer(commands chan Command) {
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 arguments"
			continue
		}
		switch cmd.Fields[0] {
		// get <KEY> <VALUE>
		case "GET":
			key := cmd.Fields[1]
			value := data[key]
			cmd.Result <- "Found " + key + " : " + value
		// set <KEY> <VALUE>
		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "EXPECTED VALUE\n"
				continue
			}
			key := cmd.Fields[1]
			value := cmd.Fields[2]
			data[key] = value
			cmd.Result <- "added " + key + " : " + value
		case "DEL":
			key := cmd.Fields[1]
			delete(data, key)
			cmd.Result <- "deleted " + key
		default:
			cmd.Result <- "INVALID COMMAND " + cmd.Fields[0] + "\n"
		}
	}
}
