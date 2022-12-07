package util

import (
	"fmt"
	"strings"
)

const (
	prefixCommand = "$"
	prefixDir     = "dir"
	cmdLS         = "ls"
	cmdCD         = "cd"
	dirUp         = ".."
	dirRoot       = "/"
)

func NewInterpreter() *interpreter {
	ip := interpreter{
		root: &dir{
			name: dirRoot,
		},
	}
	ip.current = ip.root
	return &ip
}

type interpreter struct {
	lsMode  bool
	current *dir
	root    *dir
}

func (ip *interpreter) ScanLine(s string) {
	if hasPrefix(s, prefixCommand) {
		ip.execute(trimPrefix(s, prefixCommand))
		return
	}
	if strings.HasPrefix(s, prefixDir) {
		ip.current.dirs = append(ip.current.dirs, &dir{
			name:   trimPrefix(s, prefixDir),
			parent: ip.current,
		})
		return
	}
	line := strings.Split(s, " ")
	ip.current.files = append(ip.current.files, file{
		name: line[1],
		size: ToInt(line[0]),
	})
}

func (ip *interpreter) execute(cmd string) {
	ip.lsMode = cmd == cmdLS
	if ip.lsMode {
		return
	}
	if hasPrefix(cmd, cmdCD) {
		ip.cd(trimPrefix(cmd, cmdCD))
		return
	}
	panic(fmt.Sprintf("unknown command: %s", cmd))
}

func (ip *interpreter) cd(dir string) {
	if dir == dirUp && ip.current != ip.root {
		ip.current = ip.current.parent
		return
	}
	for i, d := range ip.current.dirs {
		if d.name == dir {
			ip.current = ip.current.dirs[i]
			return
		}
	}
}

func (ip *interpreter) TotalSize() int {
	return ip.root.size()
}

func (ip *interpreter) DirSizes() map[string]int {
	ds := make(map[string]int)
	ip.root.dirSize(ds)
	delete(ds, dirRoot) // no root size
	return ds
}

func hasPrefix(s, p string) bool {
	return strings.HasPrefix(s, p)
}

func trimPrefix(s, p string) string {
	return strings.TrimPrefix(s, p+" ")
}

type dir struct {
	name   string
	parent *dir
	dirs   []*dir
	files  []file
}

type file struct {
	name string
	size int
}

func (d *dir) size() int {
	var size int
	for _, f := range d.files {
		size += f.size
	}
	for _, d := range d.dirs {
		size += d.size()
	}
	return size
}

func (d *dir) dirSize(ds map[string]int) {
	var names []string
	curr := d
	for {
		if curr.name == dirRoot {
			break
		}
		names = append([]string{curr.name}, names...)
		curr = curr.parent
	}
	path := dirRoot + strings.Join(names, "/")
	ds[path] = d.size()
	for _, d := range d.dirs {
		d.dirSize(ds)
	}
}
