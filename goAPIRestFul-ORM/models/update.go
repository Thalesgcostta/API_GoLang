package models

import "apiREST/db"

func Update(id int64, time Time) (int64, error) {
	//Abrindo a conexão com o banco
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE times SET Nome_Time=$2, Estado=$3, Tecnico=$4, Valor_Elenco=$5
	WHERE id=$1`, id, time.Nome_Time, time.Estado, time.Tecnico, time.Valor_Elenco)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
