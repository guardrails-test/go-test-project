package main
import (
    "crypto/tls"
    "fmt"
    "net/http"
)
func main() {
    tr := &http.Transport{
        // This is the insecure setting, it should be set to false.
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    _, err := client.Get("https://golang.org/")
    if err != nil {
        fmt.Println(err)
    }
}
