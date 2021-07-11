DELIMITER //
CREATE function getCurrSession() returns INT DETERMINISTIC
begin
    declare curr INT default 0;
    select user INTO curr from logIn
        where loginTime in (select max(logintime) from logIn);
    return curr;
end //
DELIMITER ;