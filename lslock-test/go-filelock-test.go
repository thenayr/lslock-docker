package main

import (
  "os"
	//"bytes"
	"bufio"
  "os/exec"
  "strings"
  "fmt"
  "io/ioutil"
	//"strconv"
)

const tmpLockDir string = "/tmp/lslock-test"
var lockPath string
var mLocks map[string]string

func main() {
	mLocks = make(map[string]string)
	ReadDirs()	
	ReadLocks()	
	fmt.Println("Printing mLocks:")
	fmt.Println(mLocks)
}


// Get a listing of files and directories
// Map filenames to iNodes (as strings)
func ReadDirs() {
  files, _ := ioutil.ReadDir(tmpLockDir)
  for _, f := range files {
      if f.IsDir() {
      } else {
        cmdName := "ls"
        cmdArgs := []string{"-li",tmpLockDir + "/" + f.Name()}
        out, err := exec.Command(cmdName, cmdArgs...).Output()
        if err != nil {
					fmt.Println(err)
        }
				cOut := string(out)
				iN := strings.Split(cOut, " ")[0]
        mLocks[f.Name()] = iN
				fmt.Printf("iN is equal to : %q\n", iN)
      }
  }
}

func ReadLocks() {
  // Open the file and scan it.
	// '/proc/locks' is mounted to /go/locks
  f, _ := os.Open("/go/locks")
  scanner := bufio.NewScanner(f)
	fileCount := 0

  for scanner.Scan() {
    line := scanner.Text()
    // Split the line on commas.
    parts := strings.Split(line, "\n")
		fileCount++
    // Loop over the parts from the string.
    for i := range parts {
			parts[i] = strings.Split(strings.Split(parts[i], ":")[3], " ")[0]
      fmt.Printf("inode: %q", parts[i])
    }
    // Write a newline.
    fmt.Println()
  }
	fmt.Printf("Total file count: %d", fileCount)
}
