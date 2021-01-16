-- drop
drop table if exists  talkRoom_message;
drop table if exists  message;
drop table if exists  user_group;
drop table if exists  `group`;
drop table if exists  friend;
drop table if exists  talk_room;
drop table if exists  user;


-- user table
CREATE TABLE user(
	id INT NOT NULL AUTO_INCREMENT,
	userName VARCHAR(20) not null,
	mail VARCHAR(50) not null,
	pass VARCHAR(20) not null,
	sex INT,
	lang INT,
	primary key(id)
);

-- talk room
CREATE TABLE talk_room(
	id INT NOT NULL AUTO_INCREMENT,
	roomName VARCHAR(50),
	primary key(id)
);

-- friend
CREATE TABLE friend(
	me INT not null,
	you INT not null,
	roomId INT not null,
	love BOOLEAN not null default 1 ,
	primary key(me,you),
	FOREIGN KEY (me) REFERENCES user(id),
	FOREIGN KEY (you) REFERENCES user(id),
	FOREIGN KEY (roomId) REFERENCES talk_room(id)
);

-- group
CREATE TABLE `group`(
	id INT NOT NULL AUTO_INCREMENT,
	groupName VARChAR(50) not null,
	roomId INT not null,
	primary key(id),
	FOREIGN KEY (roomId) REFERENCES talk_room(id)
);

-- user group
CREATE TABLE user_group(
	userId INT NOT NULL ,
	groupId INT NOT NULL,
	primary key(userid,groupId),
	FOREIGN KEY (userId) REFERENCES user(id),
	FOREIGN KEY (groupId) REFERENCES `group`(id)
);

-- message
CREATE TABLE message(
	id INT NOT NULL AUTO_INCREMENT,
	userId INT NOT NULL,
	message TEXT,
	createAt TIMESTAMP,
	roomId INT NOT NULL,
	primary key(id),
	FOREIGN KEY (userId) REFERENCES user(id),
	FOREIGN KEY (roomId) REFERENCES talk_room(id)
);





