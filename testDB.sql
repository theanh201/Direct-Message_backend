create database Direct_Backend_DB;
use Direct_Backend_DB;

select * from USER;

CREATE TABLE USER (
    USER_ID INT AUTO_INCREMENT,
    USER_EMAIL CHAR(64) unique,
    USER_PHONE_NUMB CHAR(12) unique,
    USER_PASSWORD CHAR(64),
    USER_NAME CHAR(64),
    USER_AVATAR BLOB,
    USER_BACKGROUND BLOB,
	USER_IS_DEL BIT,
    PRIMARY KEY (USER_ID)
);

CREATE TABLE USER_OPK_KEY (
    USER_ID INT,
    USER_OPK_KEY_ID INT,
    USER_OPK_KEY BINARY(32),
    USER_OPK_KEY_IS_DEL BIT,
    PRIMARY KEY (USER_ID , USER_OPK_KEY_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID)
);

CREATE TABLE USER_TOKEN (
    USER_ID INT,
    USER_TOKEN BINARY(32),
    USER_TOKEN_TIMEOUT DATETIME,
    USER_TOKEN_IS_DEL BIT,
    PRIMARY KEY (USER_TOKEN , USER_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID)
);

CREATE TABLE USER_FRIEND (
    USER_ID_1 INT,
    USER_ID_2 INT,
    USER_FRIEND_SINCE DATETIME,
	USER_FRIEND_IS_DEL BIT,
    PRIMARY KEY (USER_ID_1 , USER_ID_2),
    FOREIGN KEY (USER_ID_1)
        REFERENCES USER (USER_ID),
    FOREIGN KEY (USER_ID_2)
        REFERENCES USER (USER_ID)
);

CREATE TABLE USER_KEY (
    USER_ID INT,
    USER_KEY_IK BINARY(32),
    USER_KEY_SPK BINARY(32),
    PRIMARY KEY (USER_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID)
);

CREATE TABLE USER_GROUP (
    USER_GROUP_ID INT AUTO_INCREMENT,
    USER_ID INT,
    USER_GROUP_SINCE DATETIME,
    PRIMARY KEY (USER_GROUP_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID)
);

CREATE TABLE GROUP_LINKER (
    USER_ID INT,
    USER_GROUP_ID INT,
    GROUP_LINKER_IS_DEL BIT,
    GROUP_LINKER_SINCE DATETIME,
    PRIMARY KEY (USER_ID , USER_GROUP_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID),
    FOREIGN KEY (USER_GROUP_ID)
        REFERENCES USER_GROUP (USER_GROUP_ID)
);

CREATE TABLE MESSAGE (
    USER_GROUP_ID INT,
    USER_ID INT,
    MESSAGE_CONTENT VARBINARY(1024),
    MESSAGE_SINCE DATETIME,
    MESSAGE_DELIVERED BIT,
    PRIMARY KEY (USER_GROUP_ID, USER_ID, MESSAGE_SINCE),
    FOREIGN KEY (USER_GROUP_ID)
        REFERENCES USER_GROUP (USER_GROUP_ID),
    FOREIGN KEY (USER_ID)
        REFERENCES USER (USER_ID)
);

CREATE TABLE USER_FRIEND_REQUEST (
    USER_ID_FROM INT,
    USER_ID_TO INT,
    USER_FRIEND_REQUEST_EK VARCHAR(255),
    USER_FRIEND_REQUEST_REJECTED BIT,
    PRIMARY KEY (USER_ID_FROM , USER_ID_TO),
    FOREIGN KEY (USER_ID_FROM)
        REFERENCES USER (USER_ID),
    FOREIGN KEY (USER_ID_TO)
        REFERENCES USER (USER_ID)
);
