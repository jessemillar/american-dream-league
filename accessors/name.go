package accessors

type Name struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
}

// GetNameByID returns a name from the database by nameID
func (accessorGroup *AccessorGroup) GetNameByID(ID int) (Name, error) {
	name := &Name{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Names WHERE ID=?", ID).Scan(&name.ID, &name.FirstName, &name.MiddleName, &name.LastName)
	if err != nil {
		return Name{}, err
	}

	return *name, nil
}

// GetNameByName returns a name from the database by nameID
func (accessorGroup *AccessorGroup) GetNameByName(firstName string, middleName string, lastName string) (Name, error) {
	name := &Name{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Names WHERE firstName=? AND middleName=? AND lastName=?", firstName, middleName, lastName).Scan(&name.ID, &name.FirstName, &name.MiddleName, &name.LastName)
	if err != nil {
		return Name{}, err
	}

	return *name, nil
}

// GetAllNames gets all names from the database
func (accessorGroup *AccessorGroup) GetAllNames() ([]Name, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Names")
	if err != nil {
		return []Name{}, err
	}

	var names []Name

	for rows.Next() {
		var newName = Name{}
		err = rows.Scan(&newName.ID, &newName.FirstName, &newName.MiddleName, &newName.LastName)
		if err != nil {
			return []Name{}, err
		}

		names = append(names, newName)
	}

	return names, nil
}

// UpdateName adds a name to the database
func (accessorGroup *AccessorGroup) UpdateName(name Name) (Name, error) {
	_, err := accessorGroup.Database.Query("UPDATE Names SET firstName=?, middleName=?, lastName=? WHERE ID=?", name.FirstName, name.MiddleName, name.LastName, name.ID)
	if err != nil {
		return Name{}, err
	}

	name, err = accessorGroup.GetNameByID(name.ID)
	if err != nil {
		return Name{}, err
	}

	return name, nil
}

// MakeName adds a name to the database
func (accessorGroup *AccessorGroup) MakeName(name Name) (Name, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Names (firstName, middleName, lastName) VALUES (?,?,?)", name.FirstName, name.MiddleName, name.LastName)
	if err != nil {
		return Name{}, err
	}

	name, err = accessorGroup.GetNameByName(name.FirstName, name.MiddleName, name.LastName)
	if err != nil {
		return Name{}, err
	}

	return name, nil
}

// DeleteNameByID deletes a name from the database
func (accessorGroup *AccessorGroup) DeleteNameByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Names WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
