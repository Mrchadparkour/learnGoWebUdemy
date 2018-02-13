package main

import (
  "fmt"
  "io"
  "log"
  "net"
)

func main() {
  //returns a listener and an err
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
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
    
    io.WriteString(conn, "\nHello from TCP server\n")
    fmt.Fprintln(conn, "How is yourt day?")
    fmt.Fprintf(conn, "%v", "Good, I hope :)")
    conn.Close()
  }
}
