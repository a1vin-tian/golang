# HOW TO USE
``` shell
git clone https://github.com/a1vin-tian/golang.git

make build 

export VERSION=1.0; ./bin/httpserver 

curl localhost:8080/ -i -H "A:B"

```
OUTPUT: 
``` shell
HTTP/1.1 200 OK
A: B
Accept: */*
User-Agent: curl/7.64.1
Date: Mon, 27 Sep 2021 01:06:31 GMT
Content-Length: 0
```

