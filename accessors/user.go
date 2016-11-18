package accessors

import "github.com/labstack/echo"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUserByID returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserByID(ID int) (User, error) {
	user := &User{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Users WHERE ID=?", ID).Scan(&user.ID, &user.Name)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserByName returns a user from the database by userID
func (accessorGroup *AccessorGroup) GetUserByName(name string) (User, error) {
	user := &User{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Users WHERE name=?", name).Scan(&user.ID, &user.Name)
	if err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserID returns a userID from the database from a name
func (accessorGroup *AccessorGroup) GetUserID(name string) (int, error) {
	user, err := accessorGroup.GetUserByName(name)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
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
		err = rows.Scan(&newUser.ID, &newUser.Name)
		if err != nil {
			return []User{}, err
		}

		users = append(users, newUser)
	}

	return users, nil
}

// UpdateUser adds a user to the database
func (accessorGroup *AccessorGroup) UpdateUser(context echo.Context) (User, error) {
	user := User{}
	err := context.Bind(&user)
	if err != nil {
		return User{}, err
	}

	_, err = accessorGroup.Database.Query("UPDATE Users SET name=? WHERE ID=?", user.Name, user.ID)
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
func (accessorGroup *AccessorGroup) MakeUser(context echo.Context) (User, error) {
	user := User{}
	err := context.Bind(&user)
	if err != nil {
		return User{}, err
	}

	_, err = accessorGroup.Database.Query("INSERT INTO Users (name) VALUES (?)", user.Name)
	if err != nil {
		return User{}, err
	}

	user, err = accessorGroup.GetUserByName(user.Name)
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
