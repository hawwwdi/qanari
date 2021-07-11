delete from block
where Blocker=GetID('hadi') and Blocked=GetID('mahdi');

DELIMITER //
create PROCEDURE unblock_procedure(IN pblocked varchar(20))
begin
    delete
    from block
    where Blocker = getCurrSession()
      and Blocked = GetID(pblocked);
end //
DELIMITER ;