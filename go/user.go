package main

// User Model
type User struct {
	ID       int
	Name     string
	Address  string
	MyNumber string
	Votes    int
}

func getUser(name string, address string, myNumber string) (user User, err error) {
	row := db.QueryRow("SELECT * FROM users WHERE name = ? AND address = ? AND mynumber = ?",
		name, address, myNumber)
	err = row.Scan(&user.ID, &user.Name, &user.Address, &user.MyNumber, &user.Votes)
	return
}

// cacheUsers はメモリに全userをのせます
func cacheUsers() error {
	//usersMap = make(map[string]User, 0, 1000)
	rows, err := db.Query(`
	SELECT * FROM users
	`)
	if err != nil {
		if err != nil {
			return err
		}
	}

	defer rows.Close()
	for rows.Next() {
		var name, address, mynumber string
		var id, votes int
		err = rows.Scan(&id, &name, &address, &mynumber, &votes)
		if err != nil {
			return err
		}
		usersMap[mynumber] = &User{
			ID:       id,
			Name:     name,
			Address:  address,
			MyNumber: mynumber,
			Votes:    votes,
		}
	}
	return nil
}
