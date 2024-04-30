drop database Direct_Backend_DB;
create database Direct_Backend_DB;
use Direct_Backend_DB;
select * from USER;
CREATE TABLE USER (
    USER_ID INT AUTO_INCREMENT,
    USER_EMAIL CHAR(64) UNIQUE,
    USER_PASSWORD BINARY(32),
    USER_NAME CHAR(64),
    USER_AVATAR VARCHAR(128),
    USER_BACKGROUND VARCHAR(128),
    USER_IS_PRIVATE BIT,
    USER_IS_DEL BIT,
    PRIMARY KEY (USER_ID)
);

CREATE TABLE USER_OPK_KEY (
    USER_ID INT,
    USER_OPK_KEY BINARY(32),
    USER_OPK_KEY_IS_DEL BIT,
    PRIMARY KEY (USER_ID , USER_OPK_KEY),
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
    USER_GROUP_2_PERSON BIT, 
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
    USER_ID_FROM INT,
    USER_ID_TO INT,
    USER_GROUP_ID INT,
    MESSAGE_CONTENT VARcHAR(128),
    MESSAGE_SINCE DATETIME,
    MESSAGE_IS_ENCRYPT BIT,
    PRIMARY KEY (USER_ID_FROM, MESSAGE_SINCE),
    FOREIGN KEY (USER_GROUP_ID)
        REFERENCES USER_GROUP (USER_GROUP_ID),
    FOREIGN KEY (USER_ID_FROM)
        REFERENCES USER (USER_ID),
	FOREIGN KEY (USER_ID_TO)
		references USER (USER_ID)
);

CREATE TABLE USER_FRIEND_REQUEST (
    USER_ID_FROM INT,
    USER_ID_TO INT,
    USER_FRIEND_REQUEST_EK binary(32),
    USER_FRIEND_REQUEST_OPK binary(32),
    USER_FRIEND_REQUEST_IS_DEL BIT,
    PRIMARY KEY (USER_ID_FROM , USER_ID_TO),
    FOREIGN KEY (USER_ID_FROM)
        REFERENCES USER (USER_ID),
    FOREIGN KEY (USER_ID_TO)
        REFERENCES USER (USER_ID)
);
