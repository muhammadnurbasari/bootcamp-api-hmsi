package query

import "database/sql"

type Customers struct {
	Id    uint
	Name  string
	Phone string
	Email string
	Age   uint
}

type DB struct {
	Conn *sql.DB
}

func (db *DB) Create(c *Customers) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1, $2, $3, $4)`)

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)

	if errExec != nil {
		return errExec
	}

	return nil
}

func (db *DB) Read() (*[]Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT id, name, phone, email, age FROM customers`)

	if errExec != nil {
		return nil, errExec
	}

	var result []Customers

	for rows.Next() {
		var each Customers
		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)

		if errScan != nil {
			return nil, errScan
		}

		result = append(result, each)
	}

	return &result, nil
}

func (db *DB) Update(c *Customers) error {
	stmt, err := db.Conn.Prepare(`UPDATE customers SET name = $1, phone = $2, email = $3, age = $4 WHERE id = $5`)

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age, c.Id)

	if errExec != nil {
		return errExec
	}

	return nil
}

func (db *DB) Delete(id uint) error {
	stmt, err := db.Conn.Prepare(`DELETE FROM customers WHERE id = $1`)

	if err != nil {
		return err
	}

	_, errExec := stmt.Exec(id)

	if errExec != nil {
		return errExec
	}

	return nil
}
