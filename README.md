Go fastcgi client with fcgi params support.

####Status: In Development (future versions may change API)

Forked from https://code.google.com/p/go-fastcgi-client/


    func main() {
        reqParams := "name=value"

        env := make(map[string]string)
        env["SCRIPT_FILENAME"] = "/Users/ivan/work/test/fcgi/test.php"
        env["SERVER_SOFTWARE"] = "go / fcgiclient "
        env["REMOTE_ADDR"] = "127.0.0.1"
        env["QUERY_STRING"] = reqParams

        fcgi, err := fcgiclient.New("unix", "/tmp/php-fpm.sock")
        if err != nil {
                log.Println("err:", err)
        }

        resp, err := fcgi.Get(env)
        if err != nil {
                log.Println("err:", err)
        }
        
        content, err = ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Println("err:", err)
        }
        log.Println("content:", string(content))
    }


More example can be found in [fcgiclient_test.go](./src/fcgiclient_test.go)
