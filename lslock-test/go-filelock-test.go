package main

import (
  "os"
	"bytes"
	"bufio"
  "os/exec"
  "strings"
  "fmt"
  "io/ioutil"
	//"strconv"
)

const tmpLockDir string = "/tmp/lslock-test"
var lockPath string

func main() {

	ReadLocks()	
	ReadDirs()	
}


// Get a listing of files and directories
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
        //fmt.Printf("Locked files are \n %s\n", out)
        fmt.Printf("File is " + f.Name() + " - Inode is: %q\n", bytes.Split(out, []byte{' '})[2])
      }
  }
}

func ReadLocks() {
  // Open the file and scan it.
  f, _ := os.Open("/go/locks")
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    line := scanner.Text()
    // Split the line on commas.
    parts := strings.Split(line, "\n")
    // Loop over the parts from the string.
    for i := range parts {
			parts[i] = strings.Split(strings.Split(parts[i], ":")[3], " ")[0]
      fmt.Printf("inode: %q", parts[i])
    }
    // Write a newline.
    fmt.Println()
  }
}
