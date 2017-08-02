//Copyright (c) 2017, Will Dixon. All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

package skywalker_test

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

func ExampleSkywalker() {
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
	// Output:
	// /subfolder/few.pdf
	// /the/few.pdf
}
