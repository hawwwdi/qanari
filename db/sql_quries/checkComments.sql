select * from ava
where ava.ReplyTo = 1
    and (select ava.Publisher from ava where ava.ID = 1) not in (select block.Blocker
        from block where block.Blocked = GetID('hadi'))
    and ava.Publisher not in (select block.Blocker
        from block where block.Blocked = GetID('hadi'));

delimiter //
create procedure checkComments_procedure(IN pava INT)
begin
    select *
    from ava
    where ava.ReplyTo = pava
      and (select ava.Publisher from ava where ava.ID = pava) not in (select block.Blocker
                                                                      from block
                                                                      where block.Blocked = getCurrSession())
      and ava.Publisher not in (select block.Blocker
                                from block
                                where block.Blocked = getCurrSession());
end //
delimiter ;

