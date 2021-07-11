insert into ava (Publisher, Content, ReplyTo)
select getID('hadi'), 'first comment', 1
where GetID('hadi') not in (select block.Blocked from block
    where block.Blocker = (select ava.Publisher from ava
        where ava.ID = 1));

delimiter //
create procedure postComment_procedure(IN pcontent TINYTEXT, preplyTo int)
begin
    insert into ava (Publisher, Content, ReplyTo)
    select getCurrSession(), pcontent, preplyTo
    where getCurrSession() not in (select block.Blocked
                                   from block
                                   where block.Blocker = (select ava.Publisher
                                                          from ava
                                                          where ava.ID = preplyTo));
end //
delimiter ;