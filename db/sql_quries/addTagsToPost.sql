insert into tag(hashtag)
values ('#hadi');

insert into ava_tag(AID, TID)
values (1 ,(select tag.ID from tag where tag.hashtag = '#hadi'));


create procedure addTagToAPost(IN PAID int, IN pTag varchar(6))
begin
    declare res INT default NULL;
    select tag.ID into res from tag
        where tag.hashtag = pTag;
    IF res = NULL then
        insert into tag values (pTag);
        select tag.ID into res from tag
        where tag.hashtag = pTag;
    end if;
    insert into ava_tag values (PAID, res);
end ;
