package main

import (
    "fmt"
    "flag"
    "log"
    "net/http"
    "os"
//    "io"
    "path"
    "strconv"
    "encoding/json"
)

var dir string
var port int
var staticHandler http.Handler


func init() {
    dir = path.Dir(os.Args[0])
  //  dir = "D:/Projects/local/src/test/fileserver"
    fmt.Println(dir)
    flag.IntVar(&port, "port", 10010, "port")
    flag.Parse()
    staticHandler = http.FileServer(http.Dir(dir))
}

func main() {
    http.HandleFunc("/", StaticServer)
    err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


func allFilesJSON( dir string ) []byte  {
    f, err := os.Open(dir)
    if err != nil {
        return nil
    }
    defer f.Close()

    names, err := f.Readdirnames(-1)
    if err != nil {
        return nil
    }

    data, err := json.Marshal(names)
    return data

}

func StaticServer(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        staticHandler.ServeHTTP(w, req)
        return
    }

    w.Write(allFilesJSON(dir))

}