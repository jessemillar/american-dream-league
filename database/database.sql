DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE League(leagueID int NOT NULL, name varchar(25), PRIMARY KEY (leagueID));
CREATE TABLE Users(userID int NOT NULL, email varchar(25), password varchar(25), standing int(5), PRIMARY KEY (userID));
CREATE TABLE Email(emailID int NOT NULL, screenname varchar(25), domain varchar(25), PRIMARY KEY (emailID));
CREATE TABLE Password(passwordID int NOT NULL, hash varchar(100), PRIMARY KEY (passwordID));
CREATE TABLE Team(teamID int NOT NULL, mascot int(5), logo int(5), name varchar(25), ownerName varchar(25), PRIMARY KEY (teamID));
CREATE TABLE Mascot(mascotID int(5) NOT  NULL, name varchar(25), PRIMARY KEY (mascotID));
CREATE TABLE Logo(logoID int NOT NULL, file varchar(25), PRIMARY KEY (logoID));
CREATE TABLE Pitcher(pitcherID int NOT NULL, saves int(5), innings int(5), WHIP int(5), ERA int(5), player int(5), PRIMARY KEY (pitcherID));
CREATE TABLE Batter(batterID int NOT NULL, onBasePercent int(5), battingAverage int(5), strikeOuts int(5), slugPercent int(5), player int(5), hit int(5), atBat int(5), PRIMARY KEY (batterID));
CREATE TABLE Hit(hitID int NOT NULL, singles int(5), doubles int(5), triples int(5), homeRuns int(5), PRIMARY KEY (hitID));
CREATE TABLE AtBat(atBatID int NOT NULL, count int(5), PRIMARY KEY (atBatID));
CREATE TABLE Player(playerID int NOT NULL,age int(5), gameCount int(5), wins int(5), name int(5), position int(5), PRIMARY KEY (playerID));
CREATE TABLE Name(nameID int NOT NULL, firstName varchar(25), middleName varchar(25), lastName varchar(25), PRIMARY KEY (nameID));
CREATE TABLE Positions(positionID int NOT NULL, positionName varchar(25), PRIMARY KEY (positionID));
