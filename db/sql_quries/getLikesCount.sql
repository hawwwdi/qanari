select COUNT(*)
from postLike
where (select ava.Publisher from ava where ava.ID=1) not in
    (select block.Blocker from block where block.Blocked = getID('hadi'))
    and postLike.Post = 1;

delimiter //
create procedure getLikesCount(IN pAID INT)
begin
    select COUNT(*)
    from postLike
    where (select ava.Publisher from ava where ava.ID = pAID) not in
          (select block.Blocker from block where block.Blocked = getCurrSession())
      and postLike.Post = pAID;
end //
delimiter ;