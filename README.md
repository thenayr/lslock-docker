Description
---

lslock is golang utility to lock files and lslock-test is a golang utility to find the locked files and PID's

Run it
---

Start lslock to lock files inside your Docker VM
```docker-compose up -d lslock```

Now run lslock-test to get a list of PID and file-names that are currently locked
```docker-compose up lslock-test```
