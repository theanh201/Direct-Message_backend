create database Direct_Backend_DB;
use Direct_Backend_DB;

CREATE TABLE USER_ACCOUNT (
	USER_ACCOUNT_USERNAME char(20) unique,
    USER_ACCOUNT_PASSWORD char(32),
    USER_ACCOUNT_ID int,
    primary key (USER_ACCOUNT_ID)
);
