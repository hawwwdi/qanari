insert into block
values (GetID('hadi'), GetID('mahdi'));

DELIMITER //
create procedure block_procedure(IN blocked varchar(20))
begin
    insert into block
    values (getCurrSession(), GetID(Blocked));
end //
DELIMITER //