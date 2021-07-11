insert into follow
values (getID('hadi'), GetID('mahdi'));

DELIMITER //
create PROCEDURE follow(IN pfollowed varchar(20))
begin
    insert into follow
    values (getCurrSession(), GetID(pfollowed));
end //
DELIMITER ;