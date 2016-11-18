DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE Names(ID int NOT NULL AUTO_INCREMENT, firstName varchar(500), middleName varchar(500), lastName varchar(500), PRIMARY KEY (ID));
CREATE TABLE Leagues(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Emails(ID int NOT NULL AUTO_INCREMENT, username varchar(500), domain varchar(500), PRIMARY KEY (ID));
CREATE TABLE Passwords(ID int NOT NULL AUTO_INCREMENT, hash varchar(500), PRIMARY KEY (ID));
CREATE TABLE Users(ID int NOT NULL AUTO_INCREMENT, nameId int, emailID int, passwordID int, PRIMARY KEY (ID), FOREIGN KEY (nameID) REFERENCES Names(ID), FOREIGN KEY (emailID) REFERENCES Emails(ID), FOREIGN KEY (passwordID) REFERENCES Passwords(ID));
CREATE TABLE Mascots(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Teams(ID int NOT NULL AUTO_INCREMENT, leagueID int,  mascotID int, logoID int, nameID int, PRIMARY KEY (ID), FOREIGN KEY (leagueID) REFERENCES Leagues(ID), FOREIGN KEY (mascotID) REFERENCES Mascots(ID), FOREIGN KEY (nameID) REFERENCES Names(ID));
CREATE TABLE Logos(ID int NOT NULL AUTO_INCREMENT, filename varchar(500), PRIMARY KEY (ID));
CREATE TABLE Positions(ID int NOT NULL AUTO_INCREMENT, name varchar(500), PRIMARY KEY (ID));
CREATE TABLE Players(ID int NOT NULL AUTO_INCREMENT, age int, gameCount int, wins int, nameID int, positionID int, PRIMARY KEY (ID), FOREIGN KEY (nameID) REFERENCES Names(ID),  FOREIGN KEY (positionID) REFERENCES Positions(ID));
CREATE TABLE Pitchers(ID int NOT NULL AUTO_INCREMENT, saves int, innings int, WHIP int, ERA int, playerID int, PRIMARY KEY (ID), FOREIGN KEY (playerID) REFERENCES Players(ID));
CREATE TABLE Hitters(ID int NOT NULL AUTO_INCREMENT, onBasePercent int, battingAverage int, strikeouts int, slugPercent int, playerID int, atBats int, PRIMARY KEY (ID), FOREIGN KEY (playerID) REFERENCES Players(ID));
CREATE TABLE Hits(ID int NOT NULL AUTO_INCREMENT, singles int, doubles int, triples int, homeruns int, PRIMARY KEY (ID));
CREATE TABLE AtBats(ID int NOT NULL AUTO_INCREMENT, count int, PRIMARY KEY (ID));

INSERT INTO Leagues(name) VALUES ("Test 1");
INSERT INTO Leagues(name) VALUES ("Test 2");
INSERT INTO Leagues(name) VALUES ("Test 3");
INSERT INTO Leagues(name) VALUES ("Test 4");
INSERT INTO Leagues(name) VALUES ("Test 5");
