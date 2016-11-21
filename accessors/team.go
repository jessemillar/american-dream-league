package accessors

type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	League  League   `json:"league"`
	Mascot  Mascot   `json:"mascot"`
	Players []Player `json:"players,omitempty"`
}

// PopulateTeam takes a mostly blank Team struct and populates the values
func (accessorGroup *AccessorGroup) PopulateTeam(teamID int, name string, leagueID int, mascotID int) (Team, error) {
	league, err := accessorGroup.GetLeagueByID(leagueID)
	if err != nil {
		return Team{}, err
	}

	mascot, err := accessorGroup.GetMascotByID(mascotID)
	if err != nil {
		return Team{}, err
	}

	players, err := accessorGroup.GetAllPlayersByTeamID(teamID)
	if err != nil {
		return Team{}, err
	}

	team := Team{}
	team.ID = teamID
	team.Name = name
	team.League = league
	team.Mascot = mascot
	team.Players = players

	return team, nil
}

// PopulateTeamMeta takes a mostly blank Team struct and populates the values without the players list
func (accessorGroup *AccessorGroup) PopulateTeamMeta(teamID int, name string, leagueID int, mascotID int) (Team, error) {
	league, err := accessorGroup.GetLeagueByID(leagueID)
	if err != nil {
		return Team{}, err
	}

	mascot, err := accessorGroup.GetMascotByID(mascotID)
	if err != nil {
		return Team{}, err
	}

	team := Team{}
	team.ID = teamID
	team.Name = name
	team.League = league
	team.Mascot = mascot

	return team, nil
}

// GetTeamByID returns a team from the database by teamID
func (accessorGroup *AccessorGroup) GetTeamByID(ID int) (Team, error) {
	team := &Team{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Teams WHERE ID=?", ID).Scan(&team.ID, &team.Name, &team.League.ID, &team.Mascot.ID)
	if err != nil {
		return Team{}, err
	}

	*team, err = accessorGroup.PopulateTeam(team.ID, team.Name, team.League.ID, team.Mascot.ID)
	if err != nil {
		return Team{}, err
	}

	return *team, nil
}

// GetTeamMetaByID returns a team from the database by teamID
func (accessorGroup *AccessorGroup) GetTeamMetaByID(ID int) (Team, error) {
	team := &Team{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Teams WHERE ID=?", ID).Scan(&team.ID, &team.Name, &team.League.ID, &team.Mascot.ID)
	if err != nil {
		return Team{}, err
	}

	*team, err = accessorGroup.PopulateTeamMeta(team.ID, team.Name, team.League.ID, team.Mascot.ID)
	if err != nil {
		return Team{}, err
	}

	return *team, nil
}

// GetTeamByName returns a team from the database by teamID
func (accessorGroup *AccessorGroup) GetTeamByName(name string) (Team, error) {
	team := &Team{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Teams WHERE name=?", name).Scan(&team.ID, &team.Name)
	if err != nil {
		return Team{}, err
	}

	return *team, nil
}

// GetAllTeams gets all teams from the database
func (accessorGroup *AccessorGroup) GetAllTeams() ([]Team, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Teams")
	if err != nil {
		return []Team{}, err
	}

	var teams []Team

	for rows.Next() {
		var newTeam = Team{}
		err = rows.Scan(&newTeam.ID, &newTeam.Name, &newTeam.League.ID, &newTeam.Mascot.ID)
		if err != nil {
			return []Team{}, err
		}

		newTeam, err = accessorGroup.PopulateTeam(newTeam.ID, newTeam.Name, newTeam.League.ID, newTeam.Mascot.ID)
		if err != nil {
			return []Team{}, err
		}

		teams = append(teams, newTeam)
	}

	return teams, nil
}

// UpdateTeam adds a team to the database
func (accessorGroup *AccessorGroup) UpdateTeam(team Team) (Team, error) {
	_, err := accessorGroup.Database.Query("UPDATE Teams SET name=? WHERE ID=?", team.Name, team.ID)
	if err != nil {
		return Team{}, err
	}

	team, err = accessorGroup.GetTeamByID(team.ID)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}

// MakeTeam adds a team to the database
func (accessorGroup *AccessorGroup) MakeTeam(team Team) (Team, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Teams (name) VALUES (?)", team.Name)
	if err != nil {
		return Team{}, err
	}

	team, err = accessorGroup.GetTeamByName(team.Name)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}

// DeleteTeamByID deletes a team from the database
func (accessorGroup *AccessorGroup) DeleteTeamByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Teams WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
