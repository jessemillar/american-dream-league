DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE Names(ID int NOT NULL AUTO_INCREMENT, firstName varchar(500), lastName varchar(500), PRIMARY KEY (ID));
CREATE TABLE Leagues(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Emails(ID int NOT NULL AUTO_INCREMENT, username varchar(500), domain varchar(500), PRIMARY KEY (ID));
CREATE TABLE Passwords(ID int NOT NULL AUTO_INCREMENT, hash varchar(500), PRIMARY KEY (ID));
CREATE TABLE Users(ID int NOT NULL AUTO_INCREMENT, nameID int, emailID int, passwordID int, PRIMARY KEY (ID), FOREIGN KEY (nameID) REFERENCES Names(ID), FOREIGN KEY (emailID) REFERENCES Emails(ID), FOREIGN KEY (passwordID) REFERENCES Passwords(ID));
CREATE TABLE Mascots(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Teams(ID int NOT NULL AUTO_INCREMENT, name varchar(500), leagueID int, mascotID int, PRIMARY KEY (ID), FOREIGN KEY (leagueID) REFERENCES Leagues(ID), FOREIGN KEY (mascotID) REFERENCES Mascots(ID));
CREATE TABLE Positions(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Players(ID int NOT NULL AUTO_INCREMENT, age int, nameID int, teamID int, positionID int, PRIMARY KEY (ID), FOREIGN KEY (teamID) REFERENCES Teams(ID), FOREIGN KEY (nameID) REFERENCES Names(ID), FOREIGN KEY (positionID) REFERENCES Positions(ID));
CREATE TABLE Hitters(ID int NOT NULL AUTO_INCREMENT, playerID int, homeruns int, runsBattedIn int, battingAverage float(7,4), PRIMARY KEY (ID), FOREIGN KEY (playerID) REFERENCES Players(ID));
CREATE TABLE Pitchers(ID int NOT NULL AUTO_INCREMENT, playerID int, strikeouts int, wins int, earnedRunAverage float(7,4), PRIMARY KEY (ID), FOREIGN KEY (playerID) REFERENCES Players(ID));

delimiter //
CREATE TRIGGER SimulateGameHitter BEFORE INSERT ON Hitters
FOR EACH ROW 
BEGIN
	# 0-2 homeruns, 0-10 runsBattedIn, 0.1-0.4 battingAverage
	SET NEW.homeruns = FLOOR(RAND()*2);
	SET NEW.runsBattedIn = FLOOR(RAND()*11);
	SET NEW.battingAverage = RAND()*(0.4-0.1)+0.1;
END; //
delimiter ;

delimiter //
CREATE TRIGGER SimulateGamePitcher BEFORE INSERT ON Pitchers
FOR EACH ROW 
BEGIN
	# 0-20 strikeouts, 0-1 wins, 0.5-4 earnedRunAverage
	SET NEW.strikeouts = FLOOR(RAND()*20);
	SET NEW.wins = FLOOR(RAND()*1);
	SET NEW.earnedRunAverage = RAND()*(4-0.5)+0.5;
END; //
delimiter ;

INSERT INTO Positions(name) VALUES ("Pitcher");
INSERT INTO Positions(name) VALUES ("Catcher");
INSERT INTO Positions(name) VALUES ("First Base");
INSERT INTO Positions(name) VALUES ("Second Base");
INSERT INTO Positions(name) VALUES ("Third Base");
INSERT INTO Positions(name) VALUES ("Shortstop");
INSERT INTO Positions(name) VALUES ("Outfielder");

INSERT INTO Leagues(name) VALUES ("Draft");
INSERT INTO Mascots(name) VALUES ("Draft");
INSERT INTO Teams(leagueID, mascotID, name) VALUES (1,1,"Draft");

# Test data below this point
INSERT INTO Emails(username,domain) VALUES ("test1","gmail.com");
INSERT INTO Emails(username,domain) VALUES ("test2","gmail.com");

INSERT INTO Passwords(hash) VALUES ("badpassword");
INSERT INTO Passwords(hash) VALUES ("worsepassword");

INSERT INTO Names(firstName,lastName) VALUES ("Jesse","Millar");
INSERT INTO Names(firstName,lastName) VALUES ("Derek","Hill");

INSERT INTO Users(nameID,emailID,passwordID) VALUES (1,1,1);
INSERT INTO Users(nameID,emailID,passwordID) VALUES (2,2,2);
