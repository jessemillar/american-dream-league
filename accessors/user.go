package accessors

type User struct {
	ID       int      `json:"id"`
	Name     Name     `json:"name"`
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

// PopulateUser takes a mostly blank User struct and populates the values
func (accessorGroup *AccessorGroup) PopulateUser(userID int, nameID int, emailID int, passwordID int) (User, error) {
	name, err := accessorGroup.GetNameByID(nameID)
	if err != nil {
		return User{}, err
	}

	email, err := accessorGroup.GetEmailByID(emailID)
	if err != nil {
		return User{}, err
	}

	password, err := accessorGroup.GetPasswordByID(passwordID)
	if err != nil {
		return User{}, err
	}

	user := User{}
	user.ID = userID
	user.Name = name
	user.Email = email
	user.Password = password

	return user, nil
}

// GetUserByID returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserByID(ID int) (User, error) {
	user := &User{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Users WHERE ID=?", ID).Scan(&user.ID, &user.Name.ID, &user.Email.ID, &user.Password.ID)
	if err != nil {
		return User{}, err
	}

	*user, err = accessorGroup.PopulateUser(user.ID, user.Name.ID, user.Email.ID, user.Password.ID)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserByName returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserByName(firstName string, lastName string) (User, error) {
	user := &User{}

	name, err := accessorGroup.GetNameByName(firstName, lastName)
	if err != nil {
		return User{}, err
	}

	err = accessorGroup.Database.QueryRow("SELECT * FROM Users WHERE nameID=?", name.ID).Scan(&user.ID, &user.Name.ID, &user.Email.ID, &user.Password.ID)
	if err != nil {
		return User{}, err
	}

	*user, err = accessorGroup.PopulateUser(user.ID, user.Name.ID, user.Email.ID, user.Password.ID)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetAllUsers gets all users from the database
func (accessorGroup *AccessorGroup) GetAllUsers() ([]User, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Users")
	if err != nil {
		return []User{}, err
	}

	var users []User

	for rows.Next() {
		var newUser = User{}
		err = rows.Scan(&newUser.ID, &newUser.Name.ID, &newUser.Email.ID, &newUser.Password.ID)
		if err != nil {
			return []User{}, err
		}

		newUser, err = accessorGroup.PopulateUser(newUser.ID, newUser.Name.ID, newUser.Email.ID, newUser.Password.ID)
		if err != nil {
			return []User{}, err
		}

		users = append(users, newUser)
	}

	return users, nil
}

// UpdateUser updates user information in the database
func (accessorGroup *AccessorGroup) UpdateUser(user User) (User, error) {
	_, err := accessorGroup.Database.Query("UPDATE Users SET nameID=? AND emailID=? AND passwordID=? WHERE ID=?", user.Name.ID, user.Email.ID, user.Password.ID, user.ID)
	if err != nil {
		return User{}, err
	}

	user, err = accessorGroup.GetUserByID(user.ID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// MakeUser adds a user to the database
func (accessorGroup *AccessorGroup) MakeUser(user User) (User, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Users (nameID, emailID, passwordID) VALUES (?,?,?)", user.Name.ID, user.Email.ID, user.Password.ID)
	if err != nil {
		return User{}, err
	}

	name, err := accessorGroup.GetNameByID(user.Name.ID)
	if err != nil {
		return User{}, err
	}

	user, err = accessorGroup.GetUserByName(name.FirstName, name.LastName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// DeleteUserByID deletes a user from the database
func (accessorGroup *AccessorGroup) DeleteUserByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Users WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
