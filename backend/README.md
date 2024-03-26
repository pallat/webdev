# Go API Template

## Cross Functional Requirement

- SLA สำหรับตั้งค่า timeout ต่างๆ เช่น Read/Write database
  - มีค่า default ในกรณีที่ไม่ได้ตั้งค่า config
- Log
  - level อะไรบ้าง และแต่ละ level ต้องการข้อมูลอะไร
- Gracefully shutting down
- tracing ID คือค่าอะไร มี span หรือไม่
-

## private repo

```sh
go install golang.org/x/tools/cmd/gonew@latest
```

```git
[url "ssh://git@gitdev.devops.krungthai.com:2222/"]
        insteadOf = https://gitdev.devops.krungthai.com/
```

```sh
go env -w GOPRIVATE=gitdev.devops.krungthai.com
```

```sh
gonew github.com/pallat/webdev github.com/yourname/yourproject
```

<!-- ```sh
machine gitdev.devops.krungthai.com login pallat.a password [Pernonal Token]
``` -->

## Usage

### single main.go

ในกรณีที่มี main เดียว สามารถวาง `main.go` ไว้ที่ root project (ที่เดียวกับ go.mod)
หรือใน `cmd/main.go`

### multiple main.go

ในกรณีที่มีหลาย main.go โดยปกติจะสร้าง directory ใน cmd เช่น `cmd/batch`
และวาง `main.go` ไว้ในนั้น

### /app

path นี้ตั้งใจใช้เป็นที่รวม business logic ไว้ด้วยกัน โดยแยกตาม module package
เช่น `app/register` หรือ `app/booking`

### /database

สำหรับเป็น database connector ทำแค่ connect database

### /cache

สำหรับเป็น cache connector เช่น redis

### /mq

สำหรับเป็น messaging queue connector เช่น Kafka

### /logger

ใช้จัดการ log instance และ log function ที่สามารถเรียกใช้ได้เลย

### /githook

สรับวาง script ทำ git hook เช่น pre-commit หรือ husky

### /test

ที่นี่ไม่ใช่ที่สำหรับ unit test แต่มีไว้วาง API test เช่น .http file หรือ
Robotframework

### /docs

สำหรับวาง API documents เช่น swagger

## log.Fatal/log.Panic

ใช้เฉพาะก่อนที่ API จะ Ready เท่านั้น

## Gracefully shutting down

ทำ graceful shutdown ทุกครั้ง

## Hexagonal Architecture

Business Logic จะอยู่ใน `app/` และไม่ใช้ adapter ตรงๆ
ต้องใช้ผ่าน interface เท่านั้น
adapter เช่น database, redis, kafka, external API

unit test และ mock ต่างๆ วางไว้ใน package เดียวกันกับ code
