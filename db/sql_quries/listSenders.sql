select distinct ID, Username
from (
         select u.ID, u.UserName
         from message
                  join user
                  join user u on message.Sender = u.ID
         where message.Receiver = getID('mahdi')
         order by message.SendTime DESC
     ) as IUN;


delimiter //
create procedure getSenders()
begin
    select distinct ID, Username
    from (
             select u.ID, u.UserName
             from message
                      join user
                      join user u on message.Sender = u.ID
             where message.Receiver = getCurrSession()
             order by message.SendTime DESC
         ) as IUN;
end //
delimiter ;