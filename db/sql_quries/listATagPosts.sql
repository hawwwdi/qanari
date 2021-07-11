select *
from ava
where ava.ID in (
    select ava_tag.AID
    from ava_tag
    where ava_tag.TID = (select tag.ID from tag where tag.hashtag = '#hadi')
)
  and ava.Publisher not in
      (select block.Blocker from block where block.Blocked = getID('hadi'));


delimiter //
create procedure getPostByTag_procedure(IN hashtag varchar(6))
begin
    select *
    from ava
    where ava.ID in (
        select ava_tag.AID
        from ava_tag
        where ava_tag.TID = (select tag.ID from tag where tag.hashtag = hashtag)
    )
      and ava.Publisher not in
          (select block.Blocker from block where block.Blocked = getCurrSession());
end //
delimiter ;