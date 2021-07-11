select * from ava
where Publisher = getID('hadi');

DELIMITER //
CREATE PROCEDURE getPersonalAvas()
begin
    select * from ava
    where Publisher = getCurrSession();
end //
DELIMITER ;