-- user testdata
insert into user 
    (userName,mail,pass,sex,lang) 
values
    ('yohei','yohei@test.com','password',0,0),
    ('kyota','kyota@test.com','password',0,0),
    ('pikachu','pikachu@test.com','password',0,0),
    ('pichu','pichu@test.com','password',0,0),
    ('raichu','raichu@test.com','password',0,0);

-- talk_room testdata
insert into talk_room 
    (roomName) 
values
    ('room1'),
    ('room2'),
    ('room3'),
    ('room4'),
    ('room5');

-- friend testdata
insert into friend
    (me,you,roomId,love)
values
    ('1','2','1',true),
    ('1','3','2',true),
    ('1','4','3',true),
    ('1','5','4',true),
    ('2','1','1',true),
    ('3','1','2',true),
    ('4','1','3',true),
    ('5','1','4',true);

-- group testdata
insert into group
    (groupName,roomId)
values
    ('group10','10'),
    ('group11','11'),
    ('group12','12'),
    ('group13','13'),
    ('group14','14'),
    ('group15','15');

-- user group table
insert into user_group
    (userId,groupId)
values
    ('1','1'),
    ('1','2'),
    ('1','3'),
    ('1','4'),
    ('1','5'),
    ('1','6'),
    ('2','1'),
    ('3','1'),
    ('4','1'),
    ('5','1');

-- message data
insert into message
    (userId,message,createAt,roomId)
values
    ('1','Hello','2006-01-02 15:04:01','2'),
    ('2','Hi','2006-01-02 15:05:01','2'),
    ('1','Hello2','2006-01-02 15:06:01','2'),
    ('2','Hi2','2006-01-02 15:07:01','2'),
    ('1','Hello3','2006-01-02 15:08:01','2'),
    ('2','Hi3','2006-01-02 15:09:01','2'),
    ('1','Hello everyone1','2006-01-02 15:10:01','10'),
    ('2','Hello everyone1','2006-01-02 15:11:01','10'),
    ('3','Hello everyone1','2006-01-02 15:12:01','10'),
    ('4','Hello everyone1','2006-01-02 15:13:05','10'),
    ('5','Hello everyone1','2006-01-02 15:14:05','10'),
    ('1','Hello everyone2','2006-01-02 15:15:05','10'),
    ('2','Hello everyone2','2006-01-02 15:16:05','10'), 
    ('3','Hello everyone3','2006-01-02 15:17:05','10'),   
    ('4','Hello everyone4','2006-01-02 15:18:05','10'), 
    ('5','Hello everyone5','2006-01-02 15:19:05','10');