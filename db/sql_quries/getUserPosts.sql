select * from  ava
where ava.Publisher not in (select block.Blocker
    from block
    where block.Blocked = getID('hadi'))
    and ava.Publisher = getID('mahdi');

delimiter //
create procedure getUserAvas_procedure(IN puser varchar(20))
begin
    select *
    from ava
    where ava.Publisher not in (select block.Blocker
                                from block
                                where block.Blocked = getCurrSession())
      and ava.Publisher = getID(puser);
end //
delimiter ;