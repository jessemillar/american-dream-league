DROP DATABASE IF EXISTS american_dream_league;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE Leagues(leagueID int NOT NULL, name varchar(100), PRIMARY KEY (leagueID));
CREATE TABLE Emails(emailID int NOT NULL, screenname varchar(100), domain varchar(100), PRIMARY KEY (emailID));
CREATE TABLE Users(userID int NOT NULL, emailID int, passwordID int, standing int, PRIMARY KEY (userID), FOREIGN KEY (emailID) REFERENCES Emails(emailID));
CREATE TABLE Passwords(passwordID int NOT NULL, hash varchar(100), PRIMARY KEY (passwordID));
CREATE TABLE Teams(teamID int NOT NULL, leagueID int,  mascotID int, logo int, name varchar(100), ownerName varchar(100), PRIMARY KEY (teamID));
CREATE TABLE Mascots(mascotID int NOT NULL, name varchar(100), PRIMARY KEY (mascotID));
CREATE TABLE Logos(logoID int NOT NULL, filename varchar(100), PRIMARY KEY (logoID));
CREATE TABLE Pitchers(pitcherID int NOT NULL, saves int, innings int, WHIP int, ERA int, playerID int, PRIMARY KEY (pitcherID));
CREATE TABLE Hitters(hitterID int NOT NULL, onBasePercent int, battingAverage int, strikeouts int, slugPercent int, playerID int, hitID int, atBat int, PRIMARY KEY (hitterID));
CREATE TABLE Hits(hitID int NOT NULL, singles int, doubles int, triples int, homeRuns int, PRIMARY KEY (hitID));
CREATE TABLE AtBats(atBatID int NOT NULL, count int, PRIMARY KEY (atBatID));
CREATE TABLE Players(playerID int NOT NULL, age int, gameCount int, wins int, name int, positionID int, PRIMARY KEY (playerID));
CREATE TABLE Names(nameID int NOT NULL, firstName varchar(100), middleName varchar(100), lastName varchar(100), PRIMARY KEY (nameID));
CREATE TABLE Positions(positionID int NOT NULL, positionName varchar(100), PRIMARY KEY (positionID));
