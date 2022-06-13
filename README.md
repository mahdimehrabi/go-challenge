## Please Concern 
I implement this project without using any framework
for best performance , scalability, and I wanted you to examine how much 
I'm good at golang standard library.<br>
for the purpose of this project I didn't implement all needed features 
so interfaces don't contain all needed methods ,
and I didn't implement JSON response generator,swagger,sentry and etc. <br>

## More Scalability
Used standard library for routing not any framework and no ORM and external packages that make the performance lower used.<br>
Used [fastjson](https://github.com/valyala/fastjson) library instead of json and encoding in standard library to
speed up parsing json **It's about 15x faster**.<br>
Used [pgx](https://github.com/jackc/pgx) as database interface because 
we only use PostgreSQL as database and its faster than `database/sql` package in standard library.

## Architecture , Design
Used interface for getting tools like logger,db , so using another tool for example another
logger or db don't force you to edit all codes of different layers.


## Solution 
I have 2 solutions for this and I implemented both of them
