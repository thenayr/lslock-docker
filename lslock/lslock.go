package main

import (
  "os"
	"bufio"
  "os/exec"
  "strings"
  "fmt"
  "io/ioutil"
  "flag"
  "log"
)

// Map filenames to iNodes
var fNodes map[string]string
// and Pids to iNodes
var pNodes map[string]string
// For reading in '-d' var
var lsDir string

func main() {
  flag.StringVar(&lsDir, "d", "/tmp/lslock-test", "The directory to list locks in")
  flag.Parse()
  log.Printf("[info] Searching \"%v\" for file locks \n", lsDir)
	fNodes = make(map[string]string)
	pNodes = make(map[string]string)
	ReadDirs()	
	ReadLocks()	
}

// Get a listing of files and directories
func ReadDirs() {
  files, _ := ioutil.ReadDir(lsDir)
  for _, f := range files {
      if f.IsDir() {
      } else {
        cmdName := "ls"
        cmdArgs := []string{"-li",lsDir + "/" + f.Name()}
        out, err := exec.Command(cmdName, cmdArgs...).Output()
        if err != nil {
					fmt.Println(err)
        }
				cOut := string(out)
				iN := strings.Split(cOut, " ")[0]
        // Map filenames to iNodes (as strings)
        fNodes[iN] = f.Name()
      }
  }
}

func ReadLocks() {
  // Open the file and scan it.
	// '/proc/locks' is mounted to /go/locks
  f, _ := os.Open("/go/locks")
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    line := scanner.Text()
    // Split the line on commas.
    parts := strings.Split(line, "\n")
    // Loop over the parts from the string.
    for i := range parts {
			iNode := strings.Split(strings.Split(parts[i], ":")[3], " ")[0]
			pID := strings.Split(strings.Split(parts[i], ":")[2], " ")[0]
      pNodes[iNode] = pID
    }
  }
	FindLocks()
}

func FindLocks() {
	lCount := 0
  for key, value := range pNodes {
		i := fNodes[key]
		if i != "" {
      lCount++
      fmt.Println("Locked file - PID:" + value + " Dir: "+ lsDir + "/" + fNodes[key])
		}
	}
  if lCount == 0 {
    log.Println("[info] Found 0 locked files in directory " + lsDir)
  } else {
    log.Printf("[info] Found %d locked files\n", lCount)
  }
}
