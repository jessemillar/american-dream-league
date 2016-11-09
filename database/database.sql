# Populated from https://github.com/jessemillar/american-dream-league/blob/master/database/dev/schema.txt

DROP DATABASE IF EXISTS `american_dream_league`;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE League(leagueID int(5), name varchar(25));

CREATE TABLE User(userID int(5)), email varchar(25), password varchar(25), standing int(5));

CREATE TABLE Email( emailID int(5), screenname varchar(25), domain varchar(25));

CREATE TABLE Password(passwordID int(5), hash);

CREATE TABLE Friendship (friendshipID int(5), user int(5), user int(5));

CREATE TABLE Team(teamID int(5), mascot int(5), logo int(5), name varchar(25), ownerName varchar(25));

CREATE TABLE Mascot(mascotID int(5), name varchar(25));

CREATE TABLE Logo(logoID int(5), file varchar(25));

CREATE TABLE Pitcher(pitcherID int(5), saves int(5), innings int(5), WHIP int(5), ERA int(5), player int(5));

CREATE TABLE Batter(batterID int(5), onBasePercent int(5), battingAverage int(5), strikeOuts int(5), slugPercent int(5), player int(5), hit int(5), atBat int(5))

CREATE TABLE Hit(hitID int(5), singles int(5), doubles int(5), triples int(5), homeRuns int(5))

CREATE TABLE AtBat(atBatID int(5), count int(5))

CREATE TABLE Player(playerID int(5),age int(5), gameCount int(5), wins int(5), name int(5), position int(5));

CREATE TABLE Name(nameID int(5), firstName varchar(25), middleName varchar(25), lastName varchar(25));

CREATE TABLE Position(positionID int(5), positionName varchar(25)); 
