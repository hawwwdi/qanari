select user.ID, user.Username
from user, postLike
where (select ava.Publisher from ava where ava.ID=1) not in
    (select block.Blocker from block where block.Blocked = getID('hadi'))
    and postLike.User = user.ID and postLike.Post = 1;

delimiter //
create procedure getLikers(IN pAID INT)
begin
    select user.ID, user.Username
    from user,
         postLike
    where (select ava.Publisher from ava where ava.ID = pAID) not in
          (select block.Blocker from block where block.Blocked = getCurrSession())
      and postLike.User = user.ID
      and postLike.Post = pAID;
end //
delimiter ;