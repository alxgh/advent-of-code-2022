package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type directory struct {
	parent    *directory
	subdirs   map[string]*directory
	fileSizes uint64
}

func newdir(p *directory) *directory {
	return &directory{
		parent:  p,
		subdirs: make(map[string]*directory),
	}
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	rootdir := newdir(nil)
	cdir := rootdir
	for sc.Scan() {
		n := sc.Text()
		if n[0] == '$' {
			cmd := strings.TrimSpace(n[1:])
			if cmd == "ls" {
				// do nothing?
			} else if cmd[0:2] == "cd" {
				dir := strings.TrimSpace(cmd[2:])
				if dir == "/" {
					cdir = rootdir
				} else if dir == ".." {
					cdir = cdir.parent
					if cdir == nil {
						cdir = rootdir
					}
				} else {
					if _, ok := cdir.subdirs[dir]; !ok {
						cdir.subdirs[dir] = newdir(cdir)
					}
					cdir = cdir.subdirs[dir]
				}
			} else {
				log.Println("what?", n)
			}
			continue
		}

		// ls output
		p := strings.Split(n, " ")
		if p[0] == "dir" {
			continue
		}
		size, _ := strconv.Atoi(p[0])

		iterdir := cdir
		for iterdir != nil {
			iterdir.fileSizes += uint64(size)
			iterdir = iterdir.parent
		}
	}

	dirs := []*directory{rootdir}
	var totalSize uint64
	for len(dirs) > 0 {
		dir := dirs[0]
		if len(dirs) == 0 {
			dirs = dirs[0:0]
		} else {
			dirs = dirs[1:]
		}

		if dir.fileSizes < 100000 {
			totalSize += dir.fileSizes
		}

		for _, d := range dir.subdirs {
			dirs = append(dirs, d)
		}
	}
	fmt.Println(totalSize)
}
