
insert into message (Sender, Receiver, Content, PostID)
select getID('hadi'), getID('mahdi'), null, 1
where (select ava.Publisher from ava where ava.ID = 1) not in
      (select block.Blocker from block where block.Blocked = getID('hadi'));


delimiter //
create procedure sendMessage(IN preceiver varchar(20), IN pcontent TINYTEXT, IN pAID int)
begin

    insert into message (Sender, Receiver, Content, PostID)
    select getCurrSession(), getID(preceiver), pcontent, pAID
    where (select ava.Publisher from ava where ava.ID = pAID) not in
          (select block.Blocker from block where block.Blocked = getCurrSession());
end //
delimiter ;