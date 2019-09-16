# go-cmd-manager

Use golang manager shell cmd task, get shell cmd logging output realtime in golang manager.


### Demo

```bash
cd task-manager

go run main.go

2019/09/16 22:20:55 get: count: 0
2019/09/16 22:20:56 get: count: 1
2019/09/16 22:20:57 get: count: 2
2019/09/16 22:20:57 golang finish/end
```

### Project Struct

**python-worker**

A simple worker in python script.

The python worker could run:

```python
python ./main.py
```

> **NOTE**
>
> Logging in python script must set to stdout.


**task-manager**

The golang program manage the shell command task.

Golang use `exec.Command` to run shell script.


### Run This Project

```bash
# 1. cd go program dir
cd task-manager


# 2. run golang program
go run main.go
```


### Run Other Command

1. change `cmdName` in `task-manager/main.go`

2. Like: `ping -c 4 localhost`



