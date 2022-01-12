Собрать для MAC - go build calc.go
Собрать для Linux - CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build calc.go
Собрать для Windows - CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build calc.go 


Программа считает сложный процент от вклада для проекта SpaceBot. 
В этом проекте процент начисляется ежедневно и соответственно ежедневно можно реинвестировать.
