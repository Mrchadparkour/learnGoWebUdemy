package main

import (
  "bufio"
  "log"
  "fmt"
  "net"
  "strings"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Println(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
    }
    go handle(conn)
  }
}

func handle(conn net.Conn) {
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    ln := scanner.Text()
    encoded := rot13(ln)
    fmt.Fprintf(conn, "%s - %s\n\n", ln, encoded)
  }
  defer conn.Close()
}

func rot13(ln string) []byte {
  //turn ln into unicode array
  ln = strings.ToLower(ln)
  bs := []byte(ln)

  for i, v := range bs {
    if v <= 109 {
      bs[i] = v + 13
    } else {
      bs[i] = v - 13
    }
  }

  return bs;
}
