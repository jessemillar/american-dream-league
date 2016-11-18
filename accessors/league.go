package accessors

import "github.com/labstack/echo"

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetLeagueByID returns a league from the database by leagueID
func (accessorGroup *AccessorGroup) GetLeagueByID(ID int) (League, error) {
	league := &League{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Leagues WHERE ID=?", ID).Scan(&league.ID, &league.Name)
	if err != nil {
		return League{}, err
	}

	return *league, nil
}

// GetLeagueByName returns a league from the database by leagueID
func (accessorGroup *AccessorGroup) GetLeagueByName(name string) (League, error) {
	league := &League{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Leagues WHERE name=?", name).Scan(&league.ID, &league.Name)
	if err != nil {
		return League{}, err
	}

	return *league, nil
}

// GetLeagueID returns a leagueID from the database from a name
func (accessorGroup *AccessorGroup) GetLeagueID(name string) (int, error) {
	league, err := accessorGroup.GetLeagueByName(name)
	if err != nil {
		return 0, err
	}

	return league.ID, nil
}

// GetAllLeagues gets all leagues from the database
func (accessorGroup *AccessorGroup) GetAllLeagues() ([]League, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Leagues")
	if err != nil {
		return []League{}, err
	}

	var leagues []League

	for rows.Next() {
		var newLeague = League{}
		err = rows.Scan(&newLeague.ID, &newLeague.Name)
		if err != nil {
			return []League{}, err
		}

		leagues = append(leagues, newLeague)
	}

	return leagues, nil
}

// UpdateLeague adds a league to the database
func (accessorGroup *AccessorGroup) UpdateLeague(context echo.Context) (League, error) {
	league := League{}
	err := context.Bind(&league)
	if err != nil {
		return League{}, err
	}

	_, err = accessorGroup.Database.Query("UPDATE Leagues SET name=? WHERE ID=?", league.Name, league.ID)
	if err != nil {
		return League{}, err
	}

	league, err = accessorGroup.GetLeagueByID(league.ID)
	if err != nil {
		return League{}, err
	}

	return league, nil
}

// MakeLeague adds a league to the database
func (accessorGroup *AccessorGroup) MakeLeague(context echo.Context) (League, error) {
	league := League{}
	err := context.Bind(&league)
	if err != nil {
		return League{}, err
	}

	_, err = accessorGroup.Database.Query("INSERT INTO Leagues (name) VALUES (?)", league.Name)
	if err != nil {
		return League{}, err
	}

	league, err = accessorGroup.GetLeagueByName(league.Name)
	if err != nil {
		return League{}, err
	}

	return league, nil
}

// DeleteLeagueByID deletes a league from the database
func (accessorGroup *AccessorGroup) DeleteLeagueByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Leagues WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
