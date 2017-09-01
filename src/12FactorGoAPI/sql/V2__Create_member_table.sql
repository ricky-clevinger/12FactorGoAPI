CREATE TABLE member (
    Member_Id int(11) NOT NULL AUTO_INCREMENT,
    Member_FName varchar(45) NOT NULL,
    Member_LName varchar(45) NOT NULL,
    Email varchar(45) NOT NULL,
	Password char(64) NOT NULL,
	Role varchar(45) NOT NULL DEFAULT 'user',
    PRIMARY KEY (Member_Id)
);