package models

import "apiREST/db"

func Get(id int64) (time Time, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM times WHERE id=$1`, id)

	err = row.Scan(&time.ID, &time.Nome_Time, &time.Estado, &time.Tecnico, &time.Valor_Elenco)

	return

}
