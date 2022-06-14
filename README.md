## Please Concern 
I implement this project **without** using any framework and ORM
for best performance , scalability.<br>
for the purpose of this project I didn't implement all needed features 
so interfaces don't contain all needed methods ,
and I didn't implement JSON response generator,swagger,sentry etc.
and also I didn't implement docker-compose version for production
<br>

## More Scalability
Used standard library for routing not any framework and no ORM used. <br>
No external packages that make the performance lower used.<br>
Used [fastjson](https://github.com/valyala/fastjson) library instead of json and encoding in standard library to
speed up parsing json **It's about 15x faster**.
Used [pgx](https://github.com/jackc/pgx) as database interface because 
we only use PostgreSQL as database and its faster than `database/sql` package in standard library.

## Architecture , Design
The architecture of this project is clean architecture,
I created an image to help you understand architecture of this project better.<br>
![clean architecture](https://raw.githubusercontent.com/mahdimehrabi/go-challenge/main/clean.png)

Used interface for getting tools like logger, db, memoryDB so using another tool for example another
logger or db don't force you to edit all codes of different layers.
<br><br>
Used [uber fx](https://github.com/uber-go/fx) as dependency injection system
to increase readability and save more memory. 

## Solution 
My solution is using redis as in memory DB and storing segment user counts
every 24 hours on 00:00AM.<br>
so we only execute a postgres query every 24 hours and we have needed data for next day.



## Getting started
`git clone https://github.com/mahdimehrabi/go-challenge.git` <br>
no copy env file
`cd gin-gorm-boilerplate`<br>
`cp env.example .env` <br>

create docker volume and start
`docker volume create psql_data` <br>
`docker-compose up -d ` <br>
run migrations <br>
`make migrate-up` <br>

generate some users (seed), this command create between 0 and 1000 random users 
,so feel free to use this command as many times as you want to create more users.<br>
`make seed`

### endpoints 
POST `localhost:8000/users`  Create new user send 
example request data
```
{
    "ID":"fsas42aa3af",
    "segment":"1aa1x",
    "expiredSegment":1655209411
}
```

GET `localhost:8000/segments/count` get segments and it user counts

example response data
```
[
{"title":"RU","usersCount":1},
{"title":"z12XTowIA2","usersCount":1}
]
```