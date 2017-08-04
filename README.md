# skywalker [![GoDoc](https://godoc.org/github.com/dixonwille/skywalker?status.svg)](https://godoc.org/github.com/dixonwille/skywalker) [![Build Status](https://travis-ci.org/dixonwille/skywalker.svg?branch=master)](https://travis-ci.org/dixonwille/skywalker) [![Build status](https://ci.appveyor.com/api/projects/status/d1h7lpf0pv546amh?svg=true)](https://ci.appveyor.com/project/dixonwille/skywalker) [![codecov](https://codecov.io/gh/dixonwille/skywalker/branch/master/graph/badge.svg)](https://codecov.io/gh/dixonwille/skywalker) [![Go Report Card](https://goreportcard.com/badge/github.com/dixonwille/skywalker)](https://goreportcard.com/report/github.com/dixonwille/skywalker)

Skywalker is a package to allow one to concurrently go through a filesystem with ease.

## Features

- Concurrency
- BlackList filtering
- WhiteList filtering
- Filter by Directory
- Filter by Extension
- Glob Filtering (provided by [gobwas/glob](https://github.com/gobwas/glob))

## Example

```go
package main

import (
    "fmt"
    "sort"
    "strings"
    "sync"

    "github.com/dixonwille/skywalker"
)

type ExampleWorker struct {
    *sync.Mutex
    found []string
}

func (ew *ExampleWorker) Work(path string) {
    //This is where the necessary work should be done.
    //This will get concurrently so make sure it is thread safe if you need info across threads.
    ew.Lock()
    defer ew.Unlock()
    ew.found = append(ew.found, path)
}

func main() {
    //Following two functions are only to create and destroy data for the example
    defer teardownData()
    standupData()

    ew := new(ExampleWorker)
    ew.Mutex = new(sync.Mutex)

    //root is the root directory of the data that was stood up above
    sw := skywalker.New(root, ew)
    sw.DirListType = skywalker.LTBlacklist
    sw.DirList = []string{"sub"}
    sw.ExtListType = skywalker.LTWhitelist
    sw.ExtList = []string{".pdf"}
    err := sw.Walk()
    if err != nil {
        fmt.Println(err)
        return
    }
    sort.Sort(sort.StringSlice(ew.found))
    for _, f := range ew.found {
        fmt.Println(strings.Replace(f, sw.Root, "", 1))
    }
}
```
