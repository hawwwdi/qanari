select * from  message as m
where m.Receiver = GetID('hadi') and m.Sender=GetID('mahdi')
    and (select ava.Publisher from ava where ava.ID = m.PostID) not in
      (select block.Blocker from block where block.Blocked = getID('hadi'))
order by m.SendTime DESC;


delimiter //
create procedure getMessagesFromAUser(IN pSender varchar(20))
begin
    select *
    from message as m
    where m.Receiver = getCurrSession()
      and m.Sender = GetID(pSender)
      and (select ava.Publisher from ava where ava.ID = m.PostID) not in
          (select block.Blocker from block where block.Blocked = getCurrSession())
    order by m.SendTime DESC;
end //
delimiter ;
