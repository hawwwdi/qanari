insert into postLike(Post, User)
select 1, getID('hadi')
where (select ava.Publisher from ava where ava.ID = 1) not in (
    select block.Blocker from block where block.Blocked = GetID('hadi')
    );

delimiter //
create procedure likeAva_procedure(IN pAID INT)
begin
    insert into postLike(Post, User)
    select pAID, getCurrSession()
    where (select ava.Publisher from ava where ava.ID = pAID) not in (
        select block.Blocker
        from block
        where block.Blocked = getCurrSession()
    );
end //
delimiter ;
