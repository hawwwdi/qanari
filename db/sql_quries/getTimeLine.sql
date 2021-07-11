select * from ava
where ava.Publisher in (
    select follow.Followed from follow where follow.Follower = GetID('hadi')
    and follow.Followed not in (
        select block.blocker from block where block.blocked=GetID('hadi')
            )
) order by ava.PublishTime DESC;

DELIMITER //
create procedure getTimeLine_procedure()
begin
    select *
    from ava
    where ava.Publisher in (
        select follow.Followed
        from follow
        where follow.Follower = getCurrSession()
          and follow.Followed not in (
            select block.blocker
            from block
            where block.blocked = getCurrSession()
        )
    )
    order by ava.PublishTime DESC;
end //
DELIMITER ;