insert into ava(Publisher, Content)
values (GetID('hadi'), 'second ava');


DELIMITER //
CREATE PROCEDURE postAVA_procedure(IN pcontent TINYTEXT)
begin
    insert into ava(Publisher, Content)
    values (getCurrSession(), pcontent);
end //
DELIMITER ;


DELIMITER //
CREATE TRIGGER after_sendAVA
    AFTER INSERT ON ava FOR EACH ROW
    BEGIN
        INSERT into avaLog values (NEW.ID, New.Publisher, NEW.PublishTime);
    end //
DELIMITER ;

CREATE trigger extract_hashtag
    after insert on ava for each row
    begin
        set @counter = 1;
        loop_label: loop
        if length(new.Content) < @counter then
            leave loop_label;
        end if ;
        if substr(new.Content, @counter, 6) like '#_____' then
            call addTagToAPost(NEW.ID, substr(new.Content, @counter, 6));
        end if ;
    end loop ;
    end ;
