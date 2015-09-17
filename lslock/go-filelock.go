package main

import (
  "os"
	"bytes"
	"bufio"
  "os/exec"
  "log"
  "strings"
  "fmt"
  "io/ioutil"
	"github.com/theckman/go-flock"
	//"strconv"
	"github.com/Pallinder/go-randomdata"
  "time"
)

const tmpLockDir string = "/tmp/lslock-test"
var lockPath string

func main() {

  CreateTmpDir()

	// Create 10 locks
  for i := 0; i < 5; i++ {
		Name := randomdata.SillyName()
		_, err := ExampleFlock_TryLock(tmpLockDir + "/" + Name + ".lock")
		if err != nil {
			log.Fatalf("[error] %s", err)
		} else {
			//log.Printf("[info] Locked file: %s\n", lp)
		}
	}
	ReadLocks()	
	ReadDirs()	
	for {
   // I will run forever
	}
}

// Create a tmp dir for locking
func CreateTmpDir() {
  os.RemoveAll(tmpLockDir) 
  err := os.MkdirAll(tmpLockDir, 0644) 
  if err != nil{
    log.Fatal(err)
  } else {
    log.Println("[info] created " + tmpLockDir) 
  }
}

// Lock a file and return the path or error
func ExampleFlock_TryLock(p string) (string, error){
	// should probably put these in /var/lock
	fileLock := flock.NewFlock(p)

	locked, err := fileLock.TryLock()

	if err != nil {
		// handle locking error
		return "", err
	}

	if locked {
		log.Printf("[info] Locked path: %s; locked: %v\n", fileLock.Path(), fileLock.Locked())
		lockPath = fileLock.Path()
	}
	duration := time.Second
  time.Sleep(duration)
	return lockPath, nil
}

// Get a listing of files and directories
func ReadDirs() {
  files, _ := ioutil.ReadDir(tmpLockDir)
  for _, f := range files {
      if f.IsDir() {
      } else {
        fmt.Println(f.Name())
        cmdName := "ls"
        cmdArgs := []string{f.Name()}
        out, err := exec.Command(cmdName, cmdArgs...).Output()
        if err != nil {
					fmt.Println(err)
          log.Fatal(err)
        }
        //fmt.Printf("Locked files are \n %s\n", out)
        fmt.Printf("%q\n", bytes.Trim(out, " ")[0])
      }
  }
}

//func CatLocks() {
//	duration := time.Second
//  time.Sleep(duration)
//	cmdName := "cat"
//	cmdArgs := []string{"/go/locks"}
//	out, err := exec.Command(cmdName, cmdArgs...).Output()
//	if err != nil {
//		log.Fatal(err)
//	}
//	//fmt.Printf("Locked files are \n %s\n", out)
//	fmt.Printf("%q\n", strings.Split(out, " "))
//}

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
