package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
    i := 1
    for {
        buf, e := r.ReadBytes('\n')
        if *numberFlag {
            fmt.Fprintf(os.Stdout, "[%5d] %s", i, buf)
            i++
        } else {
            fmt.Fprintf(os.Stdout, "%s", buf)
        }
        if e == io.EOF {
            fmt.Fprintf(os.Stdout, "\n")
            break
        }
    }

    return
}

func openFileReader(name string) (r *bufio.Reader, e error) {
    f, e := os.OpenFile(name, os.O_RDONLY, 0)
    if e != nil {
        r = nil
    } else {
        r = bufio.NewReader(f)
    }
    return r, e
}

func openHttpReader(url string) (r *bufio.Reader, e error) {
    res, e := http.Get(url)
    if e != nil {
        res = nil
    } else {
        r = bufio.NewReader(res.Body)
    }
    return
}

func openPath(path string) (r *bufio.Reader, e error) {
    if strings.Index(path, "http://") == 0 || strings.Index(path, "https://") == 0 {
        return openHttpReader(path)
    } else {
        return openFileReader(path)
    }
}

func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        cat(bufio.NewReader(os.Stdin))
    } else {
        for i := 0; i < flag.NArg(); i++ {
            r, e := openPath(flag.Arg(i))
            if e != nil {
                fmt.Fprintf(os.Stderr, "Gocat: error reading from %s: %+v\n", flag.Arg(i), e)
                continue
            }
            cat(r)
        }
    }
}
