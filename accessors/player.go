package accessors

type Player struct {
	ID       int      `json:"id"`
	Age      int      `json:"age"`
	Name     Name     `json:"name"`
	Position Position `json:"position"`
}

// PopulatePlayer takes a mostly blank Player struct and populates the values
func (accessorGroup *AccessorGroup) PopulatePlayer(playerID int, age int, nameID int, positionID int) (Player, error) {
	name, err := accessorGroup.GetNameByID(nameID)
	if err != nil {
		return Player{}, err
	}

	position, err := accessorGroup.GetPositionByID(positionID)
	if err != nil {
		return Player{}, err
	}

	player := Player{}
	player.ID = playerID
	player.Age = age
	player.Name = name
	player.Position = position

	return player, nil
}

// GetPlayerByID returns a player from the database by playerID
func (accessorGroup *AccessorGroup) GetPlayerByID(ID int) (Player, error) {
	player := &Player{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Players WHERE ID=?", ID).Scan(&player.ID, &player.Age, &player.Name.ID, &player.Position.ID)
	if err != nil {
		return Player{}, err
	}

	*player, err = accessorGroup.PopulatePlayer(player.ID, player.Age, player.Name.ID, player.Position.ID)
	if err != nil {
		return Player{}, err
	}

	return *player, nil
}

// GetPlayerByName returns a player from the database by playerID
func (accessorGroup *AccessorGroup) GetPlayerByName(firstName string, middleName string, lastName string) (Player, error) {
	player := &Player{}

	name, err := accessorGroup.GetNameByName(firstName, middleName, lastName)
	if err != nil {
		return Player{}, err
	}

	err = accessorGroup.Database.QueryRow("SELECT * FROM Players WHERE nameID=?", name.ID).Scan(&player.ID, &player.Age, &player.Name.ID, &player.Position.ID)
	if err != nil {
		return Player{}, err
	}

	*player, err = accessorGroup.PopulatePlayer(player.ID, player.Age, player.Name.ID, player.Position.ID)
	if err != nil {
		return Player{}, err
	}

	return *player, nil
}

// GetAllPlayers gets all players from the database
func (accessorGroup *AccessorGroup) GetAllPlayers() ([]Player, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Players")
	if err != nil {
		return []Player{}, err
	}

	var players []Player

	for rows.Next() {
		var newPlayer = Player{}
		err = rows.Scan(&newPlayer.ID, &newPlayer.Age, &newPlayer.Name.ID, &newPlayer.Position.ID)
		if err != nil {
			return []Player{}, err
		}

		newPlayer, err = accessorGroup.PopulatePlayer(newPlayer.ID, newPlayer.Age, newPlayer.Name.ID, newPlayer.Position.ID)
		if err != nil {
			return []Player{}, err
		}

		players = append(players, newPlayer)
	}

	return players, nil
}

// UpdatePlayer updates player information in the database
func (accessorGroup *AccessorGroup) UpdatePlayer(player Player) (Player, error) {
	_, err := accessorGroup.Database.Query("UPDATE Players SET age=? AND nameID=? AND positionID=? WHERE ID=?", player.Age, player.Name.ID, player.Position.ID, player.ID)
	if err != nil {
		return Player{}, err
	}

	player, err = accessorGroup.GetPlayerByID(player.ID)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

// MakePlayer adds a player to the database
func (accessorGroup *AccessorGroup) MakePlayer(player Player) (Player, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Players (age, nameID, positionID) VALUES (?,?,?)", player.Age, player.Name.ID, player.Position.ID)
	if err != nil {
		return Player{}, err
	}

	name, err := accessorGroup.GetNameByID(player.Name.ID)
	if err != nil {
		return Player{}, err
	}

	player, err = accessorGroup.GetPlayerByName(name.FirstName, name.MiddleName, name.LastName)
	if err != nil {
		return Player{}, err
	}

	return player, nil
}

// DeletePlayerByID deletes a player from the database
func (accessorGroup *AccessorGroup) DeletePlayerByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Players WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
