# HOW TO USE
1. 构建
``` shell
git clone https://github.com/a1vin-tian/golang.git
git chekcout jike002
make release 
```

2. 通过 Docker 命令本地启动 httpserver. nsenter
```shell
$ docker run -it --rm -p8080:8080 -d docker.io/cncamp/httpserver:v1.0
$ 8ed89f5cb41ee5067becb341c497a7d00543ed3c06a8546e888a8cd8d2b0d6b7
$ curl localhost:8080/ -i -H "A:B"
$ HTTP/1.1 200 OK
A: B
Accept: */*
User-Agent: curl/7.64.1
Date: Mon, 27 Sep 2021 01:06:31 GMT
Content-Length: 0
# 获取pid
$ docker inspect --format "{{.State.Pid}}" 8ed89f5cb41ee5067becb341c497a7d00543ed3c06a8546e888a8cd8d2b0d6b7
18393
$ nsenter -t 18393 -n ip a
```