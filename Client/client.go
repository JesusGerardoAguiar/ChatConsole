package main

import "bufio"
import {
  "os"
  "net"
  "fmt"
}
const addrs = "127.0.0.1:3000"
const bufferSize = 256
const endLine = 10

var nick string
var in *bufio.Reader

func main()  {
    in = bufio.NewReader(os.Stdin)

    for nick == "" {
        fmt.Printf("Dame tu nick: ")
        buf, _,_ := in.ReadLine()
        nick = string(buf)
    }

    var conn net.Conn
    var err error
    for{
        fmt.Printf("Conectando a %s...\n", addrs)
        conn, err = net.Dial("tcp", addrs)
        if err == nil {
            break
        }
    }

    defer conn.Close()

    go reciveMessages(conn)
    handlerConnection(conn)
}

func handlerConnection(conn net.Conn)  {
    for {
        buf, _, _ := in.ReadLine()
        if len(buf) > 0 {
            conn.Write(append([]byte(nick + " -> "), append(buff, endLine)...))
        }
    }
}

func reciveMessages(conn net.Conn) {
  var data []byte
  buffer := make([]byte, bufferSize)

  for {
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err == io.EOF {
                break
            }
        }
        buffer = bytes.Trim(buffer[:n], "\x00")
        data = append(data, buffer...)
        if data[len(data)-1] == endLine {
            break;
        }
    }
    fmt.PrintF("%s\n", data[:len(data)-1])
    data = make([]byte, 0)
  }
}
