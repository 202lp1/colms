## GO GORM CRUD

Realiza el crud de una tabla de base de datos con gorm


### A. Runing local  

```
PS D:\dockr\lp1\colms\app> nodemon --exec go run main.go --signal SIGTERM

```


### B. Runing form Docker

Build docker project

```
PS D:\dockr\lp1\colms> docker-compose up --build -d
PS D:\dockr\lp1\colms> docker ps
CONTAINER ID        IMAGE                         COMMAND                  CREATED             STATUS              PORTS                    NAMES
36b836b4c783        colms_colms                   "bash"                   8 minutes ago       Up 7 minutes        0.0.0.0:8090->8080/tcp   colms-app


PS D:\dockr\lp1\colms> docker exec -it colms-app bash

or
PS D:\dockr\lp1\colms> docker exec -it colms-app sh

```

Running

```
PS D:\dockr\lp1\colms> docker exec -it colms-app bash

root@22be69ba019e:/app/server# make watch
```
