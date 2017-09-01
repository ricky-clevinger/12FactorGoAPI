CREATE TABLE library (
    Library_Id int(11) NOT NULL AUTO_INCREMENT,
    Library_Name varchar(45) NOT NULL DEFAULT 'PleaseEnterALibraryName',
    Library_Address varchar(80) NOT NULL DEFAULT 'PleaseEnterAnAddress',
	Library_DueDateLength int(11) NOT NULL DEFAULT '7',
    PRIMARY KEY (Library_Id)
);