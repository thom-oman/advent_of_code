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
	name string
	// size     int
	files    map[string]int
	parent   *directory
	children []*directory
}

func NewDirectory(name string) *directory {
	return &directory{
		name:     name,
		files:    make(map[string]int, 0),
		children: make([]*directory, 0),
	}
}

func (d *directory) fullPath() string {
	if d.parent == nil {
		return "/"
	}
	cur := d
	parents := []*directory{cur}
	for cur.parent != nil {
		parents = append(parents, cur.parent)
		cur = cur.parent
	}
	ordered := make([]string, len(parents))
	for i := 0; i < len(parents)-1; i++ {
		ordered[len(parents)-1-i] = parents[i].name
	}
	return strings.Join(ordered, "/")
}

func (d *directory) addFile(name string, size int) {
	_, ok := d.files[name]
	if !ok {
		d.files[name] = size
	}
}

func (d *directory) addChild(child *directory) {
	child.parent = d
	d.children = append(d.children, child)
}

func (d *directory) Size() int {
	var size int
	for _, child := range d.children {
		size += child.Size()
	}

	return size + d.SizeOfFiles()
}

func (d *directory) SizeOfFiles() int {
	var size int
	for _, filesize := range d.files {
		size += filesize
	}
	return size
}

func (d *directory) FindChild(name string) (*directory, bool) {
	for _, child := range d.children {
		if child.name == name {
			return child, true
		}
	}
	return nil, false
}

func Print(d directory, indent int) {
	fmt.Print(strings.Repeat("\t", indent)[:])
	fmt.Printf(
		"%v - (subdirs: %v, NoOfFiles: %v, SizeOfFiles: %v, Size: %v, Path: %v) \n",
		d.name,
		len(d.children),
		len(d.files),
		d.SizeOfFiles(), d.Size(),
		d.fullPath(),
	)
	for _, child := range d.children {
		Print(*child, indent+1)
	}
}

var currentDir *directory
var rootDir *directory
var dirs map[string]*directory
var sizeLimit, requiredFreeSpace, totalDiskSize int

func init() {
	rootDir = NewDirectory("/")
	dirs = make(map[string]*directory)
	sizeLimit = 100000
	totalDiskSize = 70000000
	requiredFreeSpace = 30000000
}

func main() {
	parseInput()

	// Not sure why didnthis did work
	// total := 0
	// count := 0
	// for _, dir := range dirs {
	// 	count++
	// 	s := dir.Size()
	// 	if s <= sizeLimit {
	// 		total += s
	// 	}
	// }
	// fmt.Println(count)
	// fmt.Println(total)

	// Part 1
	// total := 0
	// count := 0
	// bfs(func(d *directory) {
	// 	count++
	// 	s := d.Size()
	// 	if s <= sizeLimit {
	// 		total += s
	// 	}
	// })
	// fmt.Println(count)
	// fmt.Println(total)

	// Part 2
	//
	currentlyFree := totalDiskSize - rootDir.Size()
	toFree := requiredFreeSpace - currentlyFree
	fmt.Printf(
		"totalDiskSize: %v, rootDir.Size(): %v, currentlyFree: %v\n",
		totalDiskSize, rootDir.Size(), currentlyFree,
	)
	fmt.Printf(
		"requiredFreeSpace: %v, toFree: %v\n",
		requiredFreeSpace, toFree,
	)
	smallestDisk := totalDiskSize
	bfs(func(d *directory) {
		s := d.Size()
		if s >= toFree {
			if s < smallestDisk {
				smallestDisk = s
			}
		}
	})

	fmt.Println(smallestDisk)
}

func parseInput() {
	file, err := os.Open("inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0

	for scanner.Scan() {
		if i == 50 {
			break
		}
		line := scanner.Text()
		outputs := strings.Split(line, " ")

		if outputs[0] == "$" { // is a command
			arg := ""
			if len(outputs) > 2 {
				arg = outputs[2]
			}
			handleCommand(outputs[1], arg)
		} else { // else the result of an `ls`
			handleOutput(outputs...)
		}
	}
}

func handleCommand(command, arg string) {
	switch command {
	case "cd":
		newDirName := arg
		switch newDirName {
		case "/":
			currentDir = rootDir
		case "..":
			currentDir = currentDir.parent
		default:
			newDir, ok := currentDir.FindChild(newDirName)
			if !ok {
				log.Fatal("sfjfdjgfdkg")
			}
			currentDir = newDir
		}
	}
}

func handleOutput(outputs ...string) {
	if outputs[0] == "dir" {
		newDirName := outputs[1]
		// Check if directory already exists within current directory
		_, existing := currentDir.FindChild(newDirName)

		if !existing {
			// If it doesn't then add as a child to current directory
			newDir := NewDirectory(newDirName)
			currentDir.addChild(newDir)
			//
			dirs[currentDir.fullPath()] = currentDir
		} else {
			log.Fatalf("%v already exists", newDirName)
		}
	} else { // file
		filesize, _ := strconv.ParseInt(outputs[0], 10, 0)
		filename := outputs[1]
		currentDir.addFile(filename, int(filesize))
	}
}

func bfs(f func(d *directory)) {
	visited := map[string]bool{}
	queue := []*directory{}
	queue = append(queue, rootDir.children...)
	// fmt.Printf("Starting queue: %v\n", len(queue))

	for len(queue) > 0 {
		d := queue[0]
		queue = queue[1:]
		_, ok := visited[d.fullPath()]
		if !ok {
			f(d)
			visited[d.fullPath()] = true
			if len(d.children) > 0 {
				queue = append(queue, d.children...)
			}
		}
	}
}
