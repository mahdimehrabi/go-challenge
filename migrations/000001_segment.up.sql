/*
For the purpose of this project I dont create segment table and users table is enough
*/
CREATE TABLE "users" (
    ID      VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    segment VARCHAR(50) UNIQUE NOT NULL,
    expired_segment  TIMESTAMP,
)