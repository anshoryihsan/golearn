CREATE TABLE localdata_go.user (
	username VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	PRIMARY KEY (username)
)ENGINE =InnoDB;
SELECT * FROM localdata_go.user;
INSERT INTO localdata_go.user(username, password) VALUES('admin','admin');

create table localdata_go.comments(
id int not null auto_increment,
email varchar(100) not null,
comment text,
primary key (id)
)engine = InnoDB
desc localdata_go.comments
