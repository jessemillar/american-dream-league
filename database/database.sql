# Populated from https://github.com/jessemillar/american-dream-league/blob/master/database/dev/schema.txt

DROP DATABASE IF EXISTS `american_dream_league`;
CREATE DATABASE american_dream_league;
USE american_dream_league;

CREATE TABLE League(leagueID int(5), name varchar(25));
