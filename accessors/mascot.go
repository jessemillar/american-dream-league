package accessors

type Mascot struct {
	ID   int    `json:"id"`
	Hash string `json:"hash"`
}

// GetMascotByID returns a mascot from the database by mascotID
func (accessorGroup *AccessorGroup) GetMascotByID(ID int) (Mascot, error) {
	mascot := &Mascot{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Mascots WHERE ID=?", ID).Scan(&mascot.ID, &mascot.Hash)
	if err != nil {
		return Mascot{}, err
	}

	return *mascot, nil
}

// GetMascotByHash returns a mascot from the database by mascotID
func (accessorGroup *AccessorGroup) GetMascotByHash(hash string) (Mascot, error) {
	mascot := &Mascot{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Mascots WHERE hash=?", hash).Scan(&mascot.ID, &mascot.Hash)
	if err != nil {
		return Mascot{}, err
	}

	return *mascot, nil
}

// GetMascotID returns a mascotID from the database from a hash
func (accessorGroup *AccessorGroup) GetMascotID(hash string) (int, error) {
	mascot, err := accessorGroup.GetMascotByHash(hash)
	if err != nil {
		return 0, err
	}

	return mascot.ID, nil
}

// GetAllMascots gets all mascots from the database
func (accessorGroup *AccessorGroup) GetAllMascots() ([]Mascot, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Mascots")
	if err != nil {
		return []Mascot{}, err
	}

	var mascots []Mascot

	for rows.Next() {
		var newMascot = Mascot{}
		err = rows.Scan(&newMascot.ID, &newMascot.Hash)
		if err != nil {
			return []Mascot{}, err
		}

		mascots = append(mascots, newMascot)
	}

	return mascots, nil
}

// UpdateMascot adds a mascot to the database
func (accessorGroup *AccessorGroup) UpdateMascot(mascot Mascot) (Mascot, error) {
	_, err := accessorGroup.Database.Query("UPDATE Mascots SET hash=? WHERE ID=?", mascot.Hash, mascot.ID)
	if err != nil {
		return Mascot{}, err
	}

	mascot, err = accessorGroup.GetMascotByID(mascot.ID)
	if err != nil {
		return Mascot{}, err
	}

	return mascot, nil
}

// MakeMascot adds a mascot to the database
func (accessorGroup *AccessorGroup) MakeMascot(mascot Mascot) (Mascot, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Mascots (hash) VALUES (?)", mascot.Hash)
	if err != nil {
		return Mascot{}, err
	}

	mascot, err = accessorGroup.GetMascotByHash(mascot.Hash)
	if err != nil {
		return Mascot{}, err
	}

	return mascot, nil
}

// DeleteMascotByID deletes a mascot from the database
func (accessorGroup *AccessorGroup) DeleteMascotByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Mascots WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
