//buf.io stuff
package main

import (
  "fmt"
  "bufio"
  "log"
  "net"
)

func main() {
  //returns a listener and an err
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Panic(err)
  }
  defer li.Close()

  //loop forever
  for {
    //if a request comes in we accept
    //conn is a readr and a write polymorphism
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
    fmt.Println(ln)
  }
  defer conn.Close()

  //we never get here
  //we have an open stream connection
  //how does the above reader know when its done
  fmt.Println("Code got here")
}
