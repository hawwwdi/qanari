create function GetID(uname varchar(20))
    returns int
    DETERMINISTIC
begin
    declare uid int;
    set uid = (select ID
              from user
              where user.UserName = uname);
    return uid;
end;
