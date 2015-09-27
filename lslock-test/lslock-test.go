package main

import (
  "os"
  "log"
  "github.com/theckman/go-flock"
  "strconv"
  "github.com/Pallinder/go-randomdata"
  //"flag"
  //"time"
)

const tmpLockDir string = "/tmp/lslock-test"
var lockPath string

func main() {
  // TODO add SIGINT functionality
  // And flag parsing
  // lCount := flag.Int("l", 10, "The number of flocks to create")
  // flag.Parse()
  CreateTmpDir()
  MakeLocks()
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
  return lockPath, nil
}

func MakeLocks() {
  fCount := 10
  // Create 10 locks
  for i := 0; i < fCount; i++ {
    Name := randomdata.SillyName()
    _, err := ExampleFlock_TryLock(tmpLockDir + "/" + Name + ".lock")
    if err != nil {
      log.Fatalf("[error] %s", err)
    } else {
      //log.Printf("[info] Locked file: %s\n", lp)
    }
  }
  log.Printf("[info] Locked " + strconv.Itoa(fCount) + " files")
  for {
   // I will run forever
  }
}
