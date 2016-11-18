DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE Names(nameID int NOT NULL, firstName varchar(500), middleName varchar(500), lastName varchar(500), PRIMARY KEY (nameID));
CREATE TABLE Leagues(leagueID int NOT NULL, name varchar(500), PRIMARY KEY (leagueID));
CREATE TABLE Emails(emailID int NOT NULL, screenname varchar(500), provider varchar(500), PRIMARY KEY (emailID));
CREATE TABLE Passwords(passwordID int NOT NULL, hash varchar(500), PRIMARY KEY (passwordID));
CREATE TABLE Users(userID int NOT NULL, emailID int, passwordID int, standing int, PRIMARY KEY (userID), FOREIGN KEY (emailID) REFERENCES Emails(emailID), FOREIGN KEY (passwordID) REFERENCES Passwords(passwordID));
CREATE TABLE Mascots(mascotID int NOT NULL, name varchar(500), PRIMARY KEY (mascotID));
CREATE TABLE Teams(teamID int NOT NULL, leagueID int,  mascotID int, logo int, name varchar(500), ownerID int, PRIMARY KEY (teamID), FOREIGN KEY (leagueID) REFERENCES Leagues(leagueID), FOREIGN KEY (mascotID) REFERENCES Mascots(mascotID), FOREIGN KEY (ownerID) REFERENCES Names(nameID));
CREATE TABLE Logos(logoID int NOT NULL, filename varchar(500), PRIMARY KEY (logoID));
CREATE TABLE Players(playerID int NOT NULL, age int, gameCount int, wins int, name int, positionID int, PRIMARY KEY (playerID), FOREIGN KEY (name) REFERENCES Names(nameID));
CREATE TABLE Pitchers(pitcherID int NOT NULL, saves int, innings int, WHIP int, ERA int, playerID int, PRIMARY KEY (pitcherID), FOREIGN KEY (playerID) REFERENCES Players(playerID));
CREATE TABLE Hitters(hitterID int NOT NULL, onBasePercent int, battingAverage int, strikeouts int, slugPercent int, playerID int, atBats int, PRIMARY KEY (hitterID), FOREIGN KEY (playerID) REFERENCES Players(playerID));
CREATE TABLE Hits(hitID int NOT NULL, singles int, doubles int, triples int, homeruns int, PRIMARY KEY (hitID));
CREATE TABLE AtBats(atBatID int NOT NULL, count int, PRIMARY KEY (atBatID));
CREATE TABLE Positions(positionID int NOT NULL, name varchar(500), PRIMARY KEY (positionID));
