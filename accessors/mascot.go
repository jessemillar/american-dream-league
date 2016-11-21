package accessors

type Mascot struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetMascotByID returns a mascot from the database by mascotID
func (accessorGroup *AccessorGroup) GetMascotByID(ID int) (Mascot, error) {
	mascot := &Mascot{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Mascots WHERE ID=?", ID).Scan(&mascot.ID, &mascot.Name)
	if err != nil {
		return Mascot{}, err
	}

	return *mascot, nil
}

// GetMascotByName returns a mascot from the database by mascotID
func (accessorGroup *AccessorGroup) GetMascotByName(name string) (Mascot, error) {
	mascot := &Mascot{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Mascots WHERE name=?", name).Scan(&mascot.ID, &mascot.Name)
	if err != nil {
		return Mascot{}, err
	}

	return *mascot, nil
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
		err = rows.Scan(&newMascot.ID, &newMascot.Name)
		if err != nil {
			return []Mascot{}, err
		}

		mascots = append(mascots, newMascot)
	}

	return mascots, nil
}

// UpdateMascot adds a mascot to the database
func (accessorGroup *AccessorGroup) UpdateMascot(mascot Mascot) (Mascot, error) {
	_, err := accessorGroup.Database.Query("UPDATE Mascots SET name=? WHERE ID=?", mascot.Name, mascot.ID)
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
	_, err := accessorGroup.Database.Query("INSERT INTO Mascots (name) VALUES (?)", mascot.Name)
	if err != nil {
		return Mascot{}, err
	}

	mascot, err = accessorGroup.GetMascotByName(mascot.Name)
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
