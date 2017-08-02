//Copyright (c) 2017, Will Dixon. All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

//nolint
package skywalker_test

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	root       = "testingFolder"
	benchRoot  = "benchmarking"
	subFolders = []string{"the", "subfolder", "sub/subfolder"}
	subFiles   = []string{"just.txt", "a.log", "few.pdf", "files"}
)

type TestWorker struct {
	*sync.Mutex
	found map[string]struct{}
}

func NewTW() *TestWorker {
	found := make(map[string]struct{})
	mutext := new(sync.Mutex)
	return &TestWorker{
		Mutex: mutext,
		found: found,
	}
}

func (tw *TestWorker) Work(path string) {
	tw.Lock()
	defer tw.Unlock()
	tw.found[path] = struct{}{}
}

func TestMain(m *testing.M) {
	if err := standupData(); err != nil {
		log.Fatal(err)
	}
	retCode := m.Run()
	teardownData()
	os.Exit(retCode)
}

func BenchmarkFilepathWalk(b *testing.B) {
	defer teardownBenchmark()
	standupBenchmark(b.N)
	tw := NewTW()
	b.ResetTimer()
	filepath.Walk(filepath.Join(root, benchRoot), func(path string, _ os.FileInfo, _ error) error {
		tw.Work(path)
		return nil
	})
	b.ReportAllocs()
}

func BenchmarkSkyWalker(b *testing.B) {
	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			defer teardownBenchmark()
			standupBenchmark(b.N)
			sw, _ := bc.Test()
			sw.Root = filepath.Join(root, benchRoot)
			b.ResetTimer()
			sw.Walk()
			b.ReportAllocs()
		})
	}
}

func TestWalk(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			sw, tw := tc.Test()
			err := sw.Walk()
			assert.Equal(tc.expectedErr, (err != nil), "Was/Wasn't expecting an error")
			assert.Equal(len(tc.expected), len(tw.found), "Not the expected number of results")
			for _, e := range tc.expected {
				path, _ := filepath.Abs(e)
				_, ok := tw.found[path]
				assert.True(ok, "Could not find %s", e)
			}
		})
	}
}

func standupBenchmark(num int) error {
	for i := 0; i < num; i++ {
		name := strconv.Itoa(i)
		folders := strings.Split(name, "")
		if len(name) > 1 {
			path := filepath.Join(folders[:len(folders)-1]...)
			fileName := folders[len(folders)-1]
			fileName = fileName + "." + fileName
			os.MkdirAll(filepath.Join(root, benchRoot, path), 0777)
			file, _ := os.OpenFile(filepath.Join(root, benchRoot, path, fileName), os.O_CREATE|os.O_RDONLY, 0666)
			defer file.Close()
		} else {
			os.MkdirAll(filepath.Join(root, benchRoot), 0777)
			file, _ := os.OpenFile(filepath.Join(root, benchRoot, name+"."+name), os.O_CREATE|os.O_RDONLY, 0666)
			defer file.Close()
		}
	}
	return nil
}

func teardownBenchmark() error {
	return os.RemoveAll(filepath.Join(root, benchRoot))
}

func standupData() error {
	for _, sf := range subFolders {
		if err := os.MkdirAll(filepath.Join(root, sf), 0777); err != nil {
			return err
		}
		for _, f := range subFiles {
			file, err := os.OpenFile(filepath.Join(root, sf, f), os.O_RDONLY|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()
		}
	}
	return nil
}

func teardownData() error {
	return os.RemoveAll(root)
}
