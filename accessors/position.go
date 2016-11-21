package accessors

type Position struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetPositionByID returns a position from the database by positionID
func (accessorGroup *AccessorGroup) GetPositionByID(ID int) (Position, error) {
	position := &Position{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Positions WHERE ID=?", ID).Scan(&position.ID, &position.Name)
	if err != nil {
		return Position{}, err
	}

	return *position, nil
}

// GetPositionByName returns a position from the database by positionID
func (accessorGroup *AccessorGroup) GetPositionByName(name string) (Position, error) {
	position := &Position{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Positions WHERE name=?", name).Scan(&position.ID, &position.Name)
	if err != nil {
		return Position{}, err
	}

	return *position, nil
}

// GetPositionID returns a positionID from the database from a name
func (accessorGroup *AccessorGroup) GetPositionID(name string) (int, error) {
	position, err := accessorGroup.GetPositionByName(name)
	if err != nil {
		return 0, err
	}

	return position.ID, nil
}

// GetAllPositions gets all positions from the database
func (accessorGroup *AccessorGroup) GetAllPositions() ([]Position, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Positions")
	if err != nil {
		return []Position{}, err
	}

	var positions []Position

	for rows.Next() {
		var newPosition = Position{}
		err = rows.Scan(&newPosition.ID, &newPosition.Name)
		if err != nil {
			return []Position{}, err
		}

		positions = append(positions, newPosition)
	}

	return positions, nil
}

// UpdatePosition adds a position to the database
func (accessorGroup *AccessorGroup) UpdatePosition(position Position) (Position, error) {
	_, err := accessorGroup.Database.Query("UPDATE Positions SET name=? WHERE ID=?", position.Name, position.ID)
	if err != nil {
		return Position{}, err
	}

	position, err = accessorGroup.GetPositionByID(position.ID)
	if err != nil {
		return Position{}, err
	}

	return position, nil
}

// MakePosition adds a position to the database
func (accessorGroup *AccessorGroup) MakePosition(position Position) (Position, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Positions (name) VALUES (?)", position.Name)
	if err != nil {
		return Position{}, err
	}

	position, err = accessorGroup.GetPositionByName(position.Name)
	if err != nil {
		return Position{}, err
	}

	return position, nil
}

// DeletePositionByID deletes a position from the database
func (accessorGroup *AccessorGroup) DeletePositionByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Positions WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
