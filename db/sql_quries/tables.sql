create table user
(
    ID       INT PRIMARY KEY AUTO_INCREMENT,
    UserName varchar(20) UNIQUE,
    Password varchar(128) NOT NULL,
    Name     varchar(20)  NOT NULL,
    LastName varchar(20)  NOT NULL,
    Birthday date,
    JoinDate date DEFAULT CURRENT_TIMESTAMP,
    Bio      varchar(64)
);

create table block
(
    Blocker INT,
    Blocked INT,
    FOREIGN KEY (Blocker) references user (ID),
    FOREIGN KEY (Blocked) references user (ID),
    CONSTRAINT PK_block PRIMARY KEY (Blocker, Blocked)
);

create table follow
(
    Follower int,
    Followed int,
    FOREIGN KEY (Follower) references user(ID),
    FOREIGN KEY (Followed) references user(ID),
    CONSTRAINT PK_follow Primary Key (Follower, Followed)
);

create table logIn
(
    User int,
    LogInTime datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (User) references user(ID)
);

create table ava
(
    ID          int PRIMARY KEY AUTO_INCREMENT,
    Publisher   int      NOT NULL,
    Content     TINYTEXT not null,
    ReplyTo     int,
    PublishTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Publisher) references user (ID),
    FOREIGN KEY (ReplyTo) references ava (ID)
);

create table postLike
(
    Post int,
    User int,
    FOREIGN KEY (Post) references ava(ID),
    FOREIGN KEY (User) references user(ID),
    CONSTRAINT PK_like PRIMARY KEY (Post, User)
);


create table message
(
    ID       int PRIMARY KEY AUTO_INCREMENT,
    Sender   int,
    Receiver int,
    Content  TINYTEXT,
    PostID   int,
    SendTime datetime DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (Sender) references user (ID),
    FOREIGN KEY (Receiver) references user (ID),
    foreign key (PostID) references ava (ID),
    check ( Content IS NULL or PostID IS NULL )
);

CREATE table tag
(
    ID int PRIMARY KEY AUTO_INCREMENT,
    hashtag varchar(6) not null,
    check ( hashtag like '#%' )
);


create table ava_tag (
    AID int,
    TID int,
    FOREIGN KEY (AID) references ava(ID),
    FOREIGN KEY (TID) references tag(ID),
    CONSTRAINT PK_ava_tag PRIMARY KEY (AID, TID)
);
