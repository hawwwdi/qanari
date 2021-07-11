insert into user(UserName, Password, Name, LastName, Birthday, Bio)
values ('hadi', sha('password'), 'hadi', 'abbasi', NULL, 'hello world');


DELIMITER //
CREATE PROCEDURE signUp_Procudure(IN pusername varchar(20), IN ppassword varchar(128), IN pname varchar(20), IN plastname varchar(20), IN pbirthday date, IN pbio varchar(64))
BEGIN
    insert into user(UserName, Password, Name, LastName, Birthday, Bio)
        values (pusername, sha(ppassword), pname, plastname, pbirthday, pbio);
end //
DELIMITER ;

DELIMITER //
CREATE TRIGGER after_signUp
    AFTER INSERT ON user FOR EACH ROW
    BEGIN
        INSERT into signUp
            VALUES (NEW.ID, NEW.JoinDate);
    end //
END //
