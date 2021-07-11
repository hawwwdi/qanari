delete from follow
where Follower=getID('hadi') and Followed=getID('mahdi');


DELIMITER //
create PROCEDURE unfollow(IN pfollower varchar(20))
begin
    delete
    from follow
    where Follower = getCurrSession()
      and Followed = getID(pfollower);

end //
DELIMITER ;
