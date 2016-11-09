# Populated from https://github.com/jessemillar/american-dream-league/blob/master/database/dev/schema.txt

DROP DATABASE IF EXISTS `american_dream_league`;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE League(leagueID int(5), name varchar(25));

CREATE TABLE User(userID int(5)), email varchar(25), password varchar(25), standing int(5));

CREATE TABLE Email( emailID int(5), screenname varchar(25), domain varchar(25))

CREATE TABLE Password(passwordID int(5), hash)

CREATE TABLE Friendship (friendshipID int(5), user int(5), user int(5))

CREATE TABLE Team(
  teamID,
  mascot [mascotID],
  logo [logoID],
  name,
  ownerName
)

CREATE TABLE Mascot(
  mascotID,
  name
)

CREATE TABLE Logo(
  logoID,
  file
)

CREATE TABLE Pitcher(
  pitcherID,
  saves,
  innings,
  WHIP,
  ERA,
  player [playerID]
)

CREATE TABLE Batter(
  batterID,
  onBasePercent,
  battingAverage,
  strikeOuts,
  slugPercent,
  player [playerID],
  hit [hitID],
  atBat [atBatID]
)

CREATE TABLE Hit(
  hitID,
  singles,
  doubles,
  triples,
  homeRuns
)

CREATE TABLE AtBat(
  atBatID,
  count
)

CREATE TABLE Player(playerID int(5),age int(5), gameCount int(5), wins int(5), name int(5), position int(5));

CREATE TABLE Name(nameID int(5), firstName varchar(25), middleName varchar(25), lastName varchar(25));

CREATE TABLE Position(positionID int(5), positionName varchar(25)); 
