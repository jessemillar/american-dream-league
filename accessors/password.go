package accessors

type Password struct {
	ID   int    `json:"id"`
	Hash string `json:"hash"`
}

// GetPasswordByID returns a password from the database by passwordID
func (accessorGroup *AccessorGroup) GetPasswordByID(ID int) (Password, error) {
	password := &Password{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Passwords WHERE ID=?", ID).Scan(&password.ID, &password.Hash)
	if err != nil {
		return Password{}, err
	}

	return *password, nil
}

// GetPasswordByHash returns a password from the database by passwordID
func (accessorGroup *AccessorGroup) GetPasswordByHash(hash string) (Password, error) {
	password := &Password{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Passwords WHERE hash=?", hash).Scan(&password.ID, &password.Hash)
	if err != nil {
		return Password{}, err
	}

	return *password, nil
}

// GetPasswordID returns a passwordID from the database from a hash
func (accessorGroup *AccessorGroup) GetPasswordID(hash string) (int, error) {
	password, err := accessorGroup.GetPasswordByHash(hash)
	if err != nil {
		return 0, err
	}

	return password.ID, nil
}

// GetAllPasswords gets all passwords from the database
func (accessorGroup *AccessorGroup) GetAllPasswords() ([]Password, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Passwords")
	if err != nil {
		return []Password{}, err
	}

	var passwords []Password

	for rows.Next() {
		var newPassword = Password{}
		err = rows.Scan(&newPassword.ID, &newPassword.Hash)
		if err != nil {
			return []Password{}, err
		}

		passwords = append(passwords, newPassword)
	}

	return passwords, nil
}

// UpdatePassword adds a password to the database
func (accessorGroup *AccessorGroup) UpdatePassword(password Password) (Password, error) {
	_, err := accessorGroup.Database.Query("UPDATE Passwords SET hash=? WHERE ID=?", password.Hash, password.ID)
	if err != nil {
		return Password{}, err
	}

	password, err = accessorGroup.GetPasswordByID(password.ID)
	if err != nil {
		return Password{}, err
	}

	return password, nil
}

// MakePassword adds a password to the database
func (accessorGroup *AccessorGroup) MakePassword(password Password) (Password, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Passwords (hash) VALUES (?)", password.Hash)
	if err != nil {
		return Password{}, err
	}

	password, err = accessorGroup.GetPasswordByHash(password.Hash)
	if err != nil {
		return Password{}, err
	}

	return password, nil
}

// DeletePasswordByID deletes a password from the database
func (accessorGroup *AccessorGroup) DeletePasswordByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Passwords WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
