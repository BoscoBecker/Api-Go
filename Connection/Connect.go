package Connect

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectSQLServer() {
	server := "DESKTOP-G0TT3JA\\SQLEXPRESS"
	port := 1433
	user := "sa"
	password := "hollemax"
	database := "AcessoEtapaFitness"

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	// rows, err := db.Query("SELECT * FROM minha_tabela")
	// ...
}
