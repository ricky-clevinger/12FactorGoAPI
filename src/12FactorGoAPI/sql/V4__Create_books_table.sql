CREATE TABLE books (
    Book_Id int(11) NOT NULL AUTO_INCREMENT,
    Book_Title varchar(45) NOT NULL,
    Book_AuthFName varchar(45) NOT NULL,
    Book_AuthLName varchar(45) NOT NULL,
	Library_Id int(11) NOT NULL DEFAULT '1',
	Book_Check int(11) NOT NULL DEFAULT '1',
	Mid int(11) NOT NULL DEFAULT '0',
	Book_Out_Date date DEFAULT NULL,
    PRIMARY KEY (Book_Id),
	FOREIGN KEY (Library_Id) REFERENCES library(Library_Id)
);