CREATE TABLE register(
	id char(36),
   `from` varchar(255),
	destination varchar(255),
	created_at datetime
);


CREATE TABLE object(
	id char(36),
   name varchar(255),
   register_id char(36)
);