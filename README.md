# mini-project [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gauravICT/mini-project)](https://github.com/gauravICT/mini-project)

## Prerequisite

### 1. You should have installed postgresql for database if not than go through this link
```
link: https://www.postgresql.org/docs/9.3/tutorial-install.html

```
#### Make sure the version will be latest or alteast 12 or above

### 2. Pulsar
```
link: https://pulsar.apache.org/

```
#### If you want to deploye pulsar on Minikube please go through this
```
link: https://pulsar.apache.org/docs/en/kubernetes-helm/

```
### If you want run pulsar on your local machine go through this
```
link https://pulsar.apache.org/docs/en/standalone/

```
### When postgres and pulsar both will up and running
### hit the query for making table
```
CREATE TABLE IF NOT EXISTS "employee_data" (
  employee_id VARCHAR NOT NULL,
  first_name VARCHAR NOT NULL,
  last_name VARCHAR NOT NULL,
  department VARCHAR NOT NULL,
  address VARCHAR(200) NOT NULL,
  email VARCHAR NOT NULL,
  created_at VARCHAR NOT NULL,
  updated_at VARCHAR DEFAULT NULL,
  PRIMARY KEY (employee_id)
);

```

 
## Now you good to go :
```
Installation: `git clone github.com/gauravICT/mini-project`

```
### 1. Running Directly using go run after going to the directry
```
`cd mini-project`
Syntax: `go run main.go`

```

### 2. Build : 
Syntax: `go build -o executable_name path/to/main/directory`

```
go build -o application ./mini-project

or you can simply go to the directory and then type :-

go build -o application

```

## I am putting some JSON and endpoints for reference in help.txt file
## text.log file contains pulsar consumer msgs.
