select LogInTime
from logIn
where User = getID('hadi')
order by LogInTime DESC;

DELIMITER //
CREATE PROCEDURE getLogInLog(IN pusername varchar(20))
begin
    select LogInTime
        from logIn
        where User = getID(pusername)
        order by LogInTime DESC;
end //
DELIMITER ;
