//Copyright (c) 2017, Will Dixon. All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

package skywalker_test

import (
	"path/filepath"

	"github.com/dixonwille/skywalker"
)

var (
	testCases = []TestCase{
		{"Everything", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
		}},
		{"FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
		}},
		{"BlackList Directory", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub/folder"}, skywalker.LTBlacklist, nil, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
		}},
		{"BlackList Directory FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub/folder"}, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
		}},
		{"WhiteList Directory", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub/folder"}, skywalker.LTBlacklist, nil, false, false, []string{
			root,
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
		}},
		{"WhiteList Directory FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub/folder"}, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
		}},
		{"BlackList Extension", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
		}},
		{"BlackList Extension FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
		}},
		{"WhiteList Extension", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "sub/just.txt"),
		}},
		{"WhiteList Extension FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "sub/just.txt"),
		}},
		{"BlackList Glob", skywalker.LTBlacklist, []string{"**.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/files"),
		}},
		{"BlackList Glob FilesOnly", skywalker.LTBlacklist, []string{"**.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/files"),
		}},
		{"BlackList Glob FilesOnly Alt", skywalker.LTBlacklist, []string{"**/*.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/just.txt"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/files"),
			filepath.Join(root, "sub/just.txt"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/files"),
		}},
		{"WhiteList Glob", skywalker.LTWhitelist, []string{"**.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, false, false, []string{
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/few.pdf"),
		}},
		{"WhiteList Glob FilesOnly", skywalker.LTWhitelist, []string{"**.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/few.pdf"),
		}},
		{"WhiteList Glob FilesOnly Alt", skywalker.LTWhitelist, []string{"**/*.pdf"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, []string{
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/few.pdf"),
		}},
		{"BlackDir BlackExt", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub"}, skywalker.LTBlacklist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
		}},
		{"BlackDir BlackExt FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub"}, skywalker.LTBlacklist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "subfolder/a.log"),
			filepath.Join(root, "subfolder/few.pdf"),
			filepath.Join(root, "subfolder/files"),
			filepath.Join(root, "the/a.log"),
			filepath.Join(root, "the/few.pdf"),
			filepath.Join(root, "the/files"),
		}},
		{"BlackDir WhiteExt", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub"}, skywalker.LTWhitelist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "subfolder"),
			filepath.Join(root, "the"),
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "the/just.txt"),
		}},
		{"BlackDir WhiteExt FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"sub"}, skywalker.LTWhitelist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "subfolder/just.txt"),
			filepath.Join(root, "the/just.txt"),
		}},
		{"WhiteDir BlackExt", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub"}, skywalker.LTBlacklist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
		}},
		{"WhiteDir BlackExt FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub"}, skywalker.LTBlacklist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/a.log"),
			filepath.Join(root, "sub/folder/subfolder/few.pdf"),
			filepath.Join(root, "sub/folder/subfolder/files"),
			filepath.Join(root, "sub/a.log"),
			filepath.Join(root, "sub/few.pdf"),
			filepath.Join(root, "sub/files"),
		}},
		{"WhiteDir WhiteExt", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub"}, skywalker.LTWhitelist, []string{".txt"}, false, false, []string{
			root,
			filepath.Join(root, "sub"),
			filepath.Join(root, "sub/folder"),
			filepath.Join(root, "sub/folder/subfolder"),
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/just.txt"),
		}},
		{"WhiteDir WhiteExt FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"sub"}, skywalker.LTWhitelist, []string{".txt"}, true, false, []string{
			filepath.Join(root, "sub/folder/subfolder/just.txt"),
			filepath.Join(root, "sub/just.txt"),
		}},
	}
	benchCases = []TestCase{
		{"Everything", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, false, false, nil},
		{"FilesOnly", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, nil},
		{"BlackList Directory", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"0", "1", "2", "3", "4"}, skywalker.LTBlacklist, nil, true, false, nil},
		{"WhiteList Directory", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"5", "6", "7", "8", "9"}, skywalker.LTBlacklist, nil, true, false, nil},
		{"BlackList Extension", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{".0", ".1", ".2", ".3", ".4"}, true, false, nil},
		{"WhiteList Extension", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{".5", ".6", ".7", ".8", ".9"}, true, false, nil},
		{"BlackList Glob", skywalker.LTBlacklist, []string{"**/*.0", "**/*.1", "**/*.2", "**/*.3", "**/*.4"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, nil},
		{"WhiteList Glob", skywalker.LTWhitelist, []string{"**/*.5", "**/*.6", "**/*.7", "**/*.8", "**/*.9"}, skywalker.LTBlacklist, nil, skywalker.LTBlacklist, nil, true, false, nil},
		{"BlackDir BlackExt", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"0", "1", "2", "3", "4"}, skywalker.LTBlacklist, []string{".0", ".1", ".2", ".3", ".4"}, true, false, nil},
		{"BlackDir WhiteExt", skywalker.LTBlacklist, nil, skywalker.LTBlacklist, []string{"0", "1", "2", "3", "4"}, skywalker.LTWhitelist, []string{".5", ".6", ".7", ".8", ".9"}, true, false, nil},
		{"WhiteDir BlackExt", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"5", "6", "7", "8", "9"}, skywalker.LTBlacklist, []string{".0", ".1", ".2", ".3", ".4"}, true, false, nil},
		{"WhiteDir WhiteExt", skywalker.LTBlacklist, nil, skywalker.LTWhitelist, []string{"5", "6", "7", "8", "9"}, skywalker.LTWhitelist, []string{".5", ".6", ".7", ".8", ".9"}, true, false, nil},
	}
)

type TestCase struct {
	name        string
	listType    skywalker.ListType
	list        []string
	dirListType skywalker.ListType
	dirList     []string
	extListType skywalker.ListType
	extList     []string
	filesOnly   bool
	expectedErr bool
	expected    []string
}

func (tc TestCase) Test() (*skywalker.Skywalker, *TestWorker) {
	tw := NewTW()
	sw := skywalker.New(root, tw)
	sw.ListType = tc.listType
	sw.List = tc.list
	sw.DirListType = tc.dirListType
	sw.DirList = tc.dirList
	sw.ExtListType = tc.extListType
	sw.ExtList = tc.extList
	sw.FilesOnly = tc.filesOnly
	sw.NumWorkers = 50
	return sw, tw
}
