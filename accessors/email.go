package accessors

type Email struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Domain   string `json:"domain"`
}

// GetEmailByID returns a email from the database by emailID
func (accessorGroup *AccessorGroup) GetEmailByID(ID int) (Email, error) {
	email := &Email{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Emails WHERE ID=?", ID).Scan(&email.ID, &email.Username, &email.Domain)
	if err != nil {
		return Email{}, err
	}

	return *email, nil
}

// GetEmailByEmail returns a email from the database by emailID
func (accessorGroup *AccessorGroup) GetEmailByEmail(username string, domain string) (Email, error) {
	email := &Email{}
	err := accessorGroup.Database.QueryRow("SELECT * FROM Emails WHERE username=? AND domain=?", username, domain).Scan(&email.ID, &email.Username, &email.Domain)
	if err != nil {
		return Email{}, err
	}

	return *email, nil
}

// GetAllEmails gets all emails from the database
func (accessorGroup *AccessorGroup) GetAllEmails() ([]Email, error) {
	rows, err := accessorGroup.Database.Query("SELECT * FROM Emails")
	if err != nil {
		return []Email{}, err
	}

	var emails []Email

	for rows.Next() {
		var newEmail = Email{}
		err = rows.Scan(&newEmail.ID, &newEmail.Username, &newEmail.Domain)
		if err != nil {
			return []Email{}, err
		}

		emails = append(emails, newEmail)
	}

	return emails, nil
}

// UpdateEmail adds a email to the database
func (accessorGroup *AccessorGroup) UpdateEmail(email Email) (Email, error) {
	_, err := accessorGroup.Database.Query("UPDATE Emails SET username=?, domain=? WHERE ID=?", email.Username, email.Domain, email.ID)
	if err != nil {
		return Email{}, err
	}

	email, err = accessorGroup.GetEmailByID(email.ID)
	if err != nil {
		return Email{}, err
	}

	return email, nil
}

// MakeEmail adds a email to the database
func (accessorGroup *AccessorGroup) MakeEmail(email Email) (Email, error) {
	_, err := accessorGroup.Database.Query("INSERT INTO Emails (username, domain) VALUES (?,?)", email.Username, email.Domain)
	if err != nil {
		return Email{}, err
	}

	email, err = accessorGroup.GetEmailByEmail(email.Username, email.Domain)
	if err != nil {
		return Email{}, err
	}

	return email, nil
}

// DeleteEmailByID deletes a email from the database
func (accessorGroup *AccessorGroup) DeleteEmailByID(ID int) error {
	_, err := accessorGroup.Database.Query("DELETE FROM Emails WHERE ID=?", ID)
	if err != nil {
		return err
	}

	return nil
}
