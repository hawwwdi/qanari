CREATE TABLE avaLog (
    AID INT,
    Publisher INT,
    PublishTime datetime,
    FOREIGN KEY (AID) references ava(ID),
    FOREIGN KEY (Publisher) references user(ID)
);


DELIMITER //
CREATE TRIGGER after_sendAVA
    AFTER INSERT ON ava FOR EACH ROW
    BEGIN
        INSERT into avaLog values (NEW.ID, New.Publisher, NEW.PublishTime);
    end //
DELIMITER ;
