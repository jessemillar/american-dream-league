DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE Leagues(leagueID int NOT NULL, name varchar(100), PRIMARY KEY (leagueID));
CREATE TABLE Users(userID int NOT NULL, emailID int(10), passwordID int(10), standing int(10), PRIMARY KEY (userID), FOREIGN KEY (emailID) REFERENCES Emails(emailID));
CREATE TABLE Emails(emailID int NOT NULL, screenname varchar(100), domain varchar(100), PRIMARY KEY (emailID));
CREATE TABLE Passwords(passwordID int NOT NULL, hash varchar(100), PRIMARY KEY (passwordID));
CREATE TABLE Teams(teamID int NOT NULL, leagueID int(10),  mascotID int(10), logo int(10), name varchar(100), ownerName varchar(100), PRIMARY KEY (teamID));
CREATE TABLE Mascots(mascotID int(10) NOT NULL, name varchar(100), PRIMARY KEY (mascotID));
CREATE TABLE Logos(logoID int NOT NULL, filename varchar(100), PRIMARY KEY (logoID));
CREATE TABLE Pitchers(pitcherID int NOT NULL, saves int(10), innings int(10), WHIP int(10), ERA int(10), playerID int(10), PRIMARY KEY (pitcherID));
CREATE TABLE Hitters(hitterID int NOT NULL, onBasePercent int(10), battingAverage int(10), strikeouts int(10), slugPercent int(10), playerID int(10), hitID int(10), atBat int(10), PRIMARY KEY (hitterID));
CREATE TABLE Hits(hitID int NOT NULL, singles int(10), doubles int(10), triples int(10), homeRuns int(10), PRIMARY KEY (hitID));
CREATE TABLE AtBats(atBatID int NOT NULL, count int(10), PRIMARY KEY (atBatID));
CREATE TABLE Players(playerID int NOT NULL, age int(10), gameCount int(10), wins int(10), name int(10), positionID int(10), PRIMARY KEY (playerID));
CREATE TABLE Names(nameID int NOT NULL, firstName varchar(100), middleName varchar(100), lastName varchar(100), PRIMARY KEY (nameID));
CREATE TABLE Positions(positionID int NOT NULL, positionName varchar(100), PRIMARY KEY (positionID));
