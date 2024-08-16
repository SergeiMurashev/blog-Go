CREATE TABLE "Users" (
	name				VARCHAR(255),
	email				VARCHAR(255) 	PRIMARY KEY,
	password			VARCHAR(255),
	create_date			DATE
);

CREATE TABLE "Post" (
	id					serial				PRIMARY KEY,
	title				VARCHAR(255),
	text				VARCHAR(255),
	create_date			DATE,
	author				VARCHAR(255) 	REFERENCES "Users"(email)
);

CREATE TABLE "Comment" (
	id					serial PRIMARY KEY,
	text				VARCHAR(255),
	create_date			DATE,
	author				VARCHAR(255) 	REFERENCES "Users"(email),
	post				INT 			REFERENCES "Post"(id) on DELETE CASCADE
);







