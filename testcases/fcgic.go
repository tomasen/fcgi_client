package main

import (
 "net"
 "net/http"
 "net/http/fcgi"
 "fmt"
 "time"
 fcgiclient "bitbucket.org/PinIdea/fcgi_client"
)

const (
  ip_port = "127.0.0.1:3647"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    
  fmt.Println("Server:")
  fmt.Println(resp)
  fmt.Println(req)
}


func sendFcgiRequest(addr string, token string) (content []byte) {

  fcgi, err := fcgiclient.New("tcp", addr)
  if err != nil {
    fmt.Printf("err: %v\n", err)
    return
  }
  fmt.Println("c: conn")
  fcgi_params := make(map[string]string,1)
  fcgi_params["test"] = token
  fmt.Println("c: sending")
  content, err = fcgi.Request(nil, fcgi_params, "")
  if err != nil {
    fmt.Printf("err: %v\n", err)
    return
  }
  fmt.Println("c: sent")
  
  fcgi.Close()
  return 
}


func main() {
  
  // server
  go func() {
    listener, err := net.Listen("tcp", ip_port)
    if err != nil {
      // handle error
      fmt.Println("listener creatation failed: ", err)
    }

    srv := new(FastCGIServer)
    fcgi.Serve(listener, srv)
  }()
  
  time.Sleep(1 * time.Second)
  fmt.Println("fcgi1: start")
  sendFcgiRequest(ip_port, "1")
  fmt.Println("fcgi1: stop")
  
  time.Sleep(20 * time.Second)
  
  fmt.Println("fcgi2: start")
  sendFcgiRequest(ip_port, "2")
  fmt.Println("fcgi2: stop")
}
