package main

import (
  "net"
  "log"
  "fmt"
  "bufio"
  "strings"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln(err)
    }
    go handle(conn)
  }
}

func handle(conn net.Conn) {
  defer conn.Close()

  //instructions
  fmt.Fprintln(conn, "SET: set a value, SET KEY VALUE\nGET: get value, GET KEY\nDEL: delete a value, DEL KEY")

  data := make(map[string]string)
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    ln := scanner.Text()
    fs := strings.Fields(ln)

    switch fs[0] {
    case "GET":
      key := fs[1]
      val := data[key]
      fmt.Fprintf(conn, "%s\n", val)
    case "SET":
      if len(fs) != 3 {
        fmt.Fprintln(conn, "EXPECTS A VALUE")
        continue
      }
      key := fs[1]
      val := fs[2]
      data[key] = val
    case "DEL":
      key := fs[1]
      delete(data, key)
    default:
      fmt.Fprintln(conn, "INVALID COMMAND")
    }
  }
}
