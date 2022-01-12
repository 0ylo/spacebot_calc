# spacebot_calc

---

Программа считает сложный процент от вклада для проекта SpaceBot.
В этом проекте процент начисляется ежедневно и соответственно ежедневно можно реинвестировать.

## Build
#### Mac
```bash
go build
```

#### Linux
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

#### Windows
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```


