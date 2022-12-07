package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const GT = "gt"
const LT = "lt"

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "1 || 2")
	flag.Parse()
	fmt.Println("Running part: ", part)

	if part == 1 {
		answer := part1(input)
		fmt.Println("Answer:", answer)
	} else {
		answer := part2(input)
		fmt.Println("Answer:", answer)
	}
}

func part1(input string) int {
	// build folder structure, including totalled directory sizes
	folder := calculateDirectorySize(createFolderStructure(input))

	// find all directories with size < 100000
	var totalSizeForDirectories int
	for _, dir := range findDirectoriesWithSize(folder, 100000, LT, false, make([]directory, 1)) {
		totalSizeForDirectories += dir.size
	}

	return totalSizeForDirectories
}

func part2(input string) int {
	// filesystem size = 7000000
	// needed space = 3000000
	// find smallest directory to remove to free up enough space for the update
	// build folder structure, including totalled directory sizes
	folder := calculateDirectorySize(createFolderStructure(input))

	maxSpace := 70000000
	neededSpace := 30000000

	freeSpace := maxSpace - folder.size

	minimumRemoval := neededSpace - freeSpace

	var smallestDirSize int
	for _, dir := range findDirectoriesWithSize(folder, minimumRemoval, GT, false, make([]directory, 1)) {
		if smallestDirSize == 0 || smallestDirSize > dir.size {
			smallestDirSize = dir.size
		}
	}

	return smallestDirSize
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}

func findDirectoriesWithSize(dir directory, size int, operator string, inclusive bool, directories []directory) []directory {
	for _, subDir := range dir.directories {
		directories = findDirectoriesWithSize(*subDir, size, operator, inclusive, directories)
		switch operator {
		case LT:
			if subDir.size < size {
				directories = append(directories, *subDir)
			}
		case GT:
			if subDir.size > size {
				directories = append(directories, *subDir)
			}
		}
	}

	return directories
}

func calculateDirectorySize(dir directory) directory {
	for _, subDir := range dir.directories {
		subDir.size = (calculateDirectorySize(*subDir)).size
		dir.size += subDir.size
	}
	for _, file := range dir.files {
		dir.size += file.size
	}

	return dir
}

func createFolderStructure(input string) (rootDir directory) {
	var currentDirectory *directory
	rootDir = directory{
		name:        "/",
		directories: make(map[string]*directory),
	}
	currentDirectory = &rootDir
	for _, line := range parseInput(input) {
		if line[0] == '$' && line[2] == 'c' {

			switch line[len(line)-1] {
			case '.':
				currentDirectory = currentDirectory.parent
				break
			case '/':
				continue
			default:
				currentDirectory = currentDirectory.directories[string(line[5:])]
				// fmt.Println("Current directory fetched from rootDir: ", currentDirectory)
			}
		} else {
			if line[0] == 'd' {
				dirName := string(line[4:])

				currentDirectory.directories[dirName] = &directory{
					parent:      currentDirectory,
					name:        dirName,
					directories: make(map[string]*directory),
				}
			} else if line[0] != '$' {
				splitFile := strings.Split(line, " ")
				size, _ := strconv.Atoi(splitFile[0])
				// fmt.Println("Current directory to add files to: ", currentDirectory)
				currentDirectory.files = append(currentDirectory.files, file{
					name: splitFile[1],
					size: size,
				})
			}
		}
	}
	return
}

type directory struct {
	parent      *directory
	name        string
	size        int
	files       []file
	directories map[string]*directory
}

type file struct {
	name string
	size int
}
