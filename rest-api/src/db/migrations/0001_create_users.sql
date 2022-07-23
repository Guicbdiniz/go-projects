CREATE TABLE IF NOT EXISTS users (
	ID SERIAL NOT NULL,
	Username varchar(50) NOT NULL,
	Password varchar(50) NOT NULL,
	PRIMARY KEY (ID)
);