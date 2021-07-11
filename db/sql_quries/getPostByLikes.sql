select a.ID as avaID, a.Content as content, count(User) as likes
from postLike
         right join ava a on a.ID = postLike.Post
where a.Publisher not in (
    select block.Blocker
    from block
    where block.Blocked = GetID('hadi')
)
group by a.ID
order by likes DESC;

delimiter //
create procedure getAVAsByLikes()
begin
    select a.ID as avaID, a.Content as content, count(User) as likes
    from postLike
             right join ava a on a.ID = postLike.Post
    where a.Publisher not in (
        select block.Blocker
        from block
        where block.Blocked = getCurrSession()
    )
    group by a.ID
    order by likes DESC;
end //
delimiter ;