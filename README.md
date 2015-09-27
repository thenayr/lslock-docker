Description
---
lslock is a golang utility to list the PID and path of a locked file inside of a directory. lslock-test is a helper utility that creates random locked files inside of the directory '/tmp/lslock-test'

Run it
---
### lslock
Run lslock container to have it output the PID and path of locked files in the specified directory (/tmp/lslock-test by default)

```docker-compose up lslock```

Pass in a custom path to lslock

```docker-compose run lslock -d /path/to/search```

### lslock-test
lslock-test will simply create 10 random locks in /tmp/lslock-test and wait forever. It doesn't take any arguments.

```docker-compose up lslock-test```

Notes
---
* Because we are locking files using Docker, PID is always going to report "01", if you run outside of Docker, this shouldn't be the case ;)
* /proc/locks is mounted inside of both our app containers at "/go/locks", /tmp is also mounted directly to /tmp
* I'm sure there are many other nuances to running the flock system call from inside of a container...
* This is my first project written in golang

TODO
---
Add more CLI flags for 
