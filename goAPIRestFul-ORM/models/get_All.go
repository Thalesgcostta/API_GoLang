package models

import "apiREST/db"

// Função GetAll  que retorna um slice [] de album ou um erro
// Busca todos os albums
func GetAll() (times []Time, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM times`)
	if err != nil {
		return
	}

	for rows.Next() {
		var time Time
		err = rows.Scan(&time.ID, &time.Nome_Time, &time.Estado, &time.Tecnico, &time.Valor_Elenco)
		if err != nil {
			continue
		}

		times = append(times, time)
	}

	return
}
