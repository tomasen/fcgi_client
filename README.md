Go fastcgi client with fcgi params support.

Based on https://code.google.com/p/go-fastcgi-client/

<code language="golang">
func main() {
        reqParams := "name=value"

        env := make(map[string]string)
        env["REQUEST_METHOD"] = "GET"
        env["SCRIPT_FILENAME"] = "/Users/ivan/work/test/fcgi/test.php"
        env["SERVER_SOFTWARE"] = "go / fcgiclient "
        env["REMOTE_ADDR"] = "127.0.0.1"
        env["SERVER_PROTOCOL"] = "HTTP/1.1"
        env["QUERY_STRING"] = reqParams

        fcgi, err := fcgiclient.New("unix", "/tmp/php-fpm.sock")
        if err != nil {
                fmt.Printf("err: %v", err)
        }

        content, err := fcgi.Request(nil, env, "")
        if err != nil {
                fmt.Printf("err: %v", err)
        }

        fmt.Printf("content: %s", content)
}
</code>

or:
<code language="golang">

package main

import (
 "net"
 "net/http"
 "net/http/fcgi"
 "fmt"
 "time"
 "code.google.com/r/tomasen-go-fastcgi-client/"
)

const (
  ip_port = "127.0.0.1:3647"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
    
  fmt.Println("Server:" , req)
}


func sendFcgiRequest(addr string, token string) (content []byte, err error) {

  fcgi, err = fcgiclient.New("tcp", addr)
  if err != nil {
    fmt.Printf("err: %v\n", err)
    return
  }

  fcgi_params := make(map[string]string,1)
  fcgi_params["key"] = token
  content, err = fcgi.Request(nil, fcgi_params, "")
  if err != nil {
    return
  }
  
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
  sendFcgiRequest(ip_port, "blahblah")
 
}
</code>