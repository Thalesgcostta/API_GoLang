package models

import "apiREST/db"

func Insert(time Time) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO times (Nome_Time, Estado, Tecnico, Valor_Elenco) VALUES ($1, $2, $3, $4) RETURNING id`

	err = conn.QueryRow(sql, time.Nome_Time, time.Estado, time.Tecnico, time.Valor_Elenco).Scan(&id)

	return
}
