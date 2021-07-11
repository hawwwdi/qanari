
insert into logIn(User)
select ID
from user
where UserName='hadi' and Password=sha('password');


DELIMITER //
CREATE PROCEDURE logIn_Procedure(IN pusername varchar(20), IN ppassword varchar(128))
begin
    insert into logIn(User)
        select ID
        from user
        where UserName=pusername and Password=sha(ppassword);
end //
DELIMITER ;
