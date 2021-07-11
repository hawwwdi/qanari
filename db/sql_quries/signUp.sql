CREATE Table signUp (
    User INT,
    SingUpTime datetime default CURRENT_TIMESTAMP,
    foreign key (User) references user(ID)
);
