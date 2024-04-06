package Connect

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func ConnectSQLServer() {
	server := "DESKTOP-G0TT3JA\\SQLEXPRESS"
	port := 1433
	user := "sa"
	password := "cidadao"
	database := "DB"

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}

func ExecutarConsulta(consulta string) (string, error) {
	// Verifica se a conexão com o banco de dados foi estabelecida
	if db == nil {
		return "", fmt.Errorf("conexão com o banco de dados não foi estabelecida")
	}

	// Executar a consulta SELECT
	rows, err := db.Query(consulta)
	if err != nil {
		return "", fmt.Errorf("erro ao executar a consulta: %v", err)
	}
	defer rows.Close()

	// Obter as colunas do resultado
	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("erro ao obter colunas: %v", err)
	}

	// Criar uma slice para armazenar os valores de cada linha
	var resultados []map[string]interface{}

	// Iterar sobre os resultados
	for rows.Next() {
		// Criar um slice para armazenar os valores de cada coluna na linha
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// Escanear os valores da linha
		if err := rows.Scan(scanArgs...); err != nil {
			return "", fmt.Errorf("erro ao escanear os resultados: %v", err)
		}

		// Criar um mapa para armazenar os valores de cada linha
		linha := make(map[string]interface{})
		for i, col := range columns {
			linha[col] = values[i]
		}

		// Adicionar o mapa de valores à slice de resultados
		resultados = append(resultados, linha)
	}

	// Verificar se ocorreu algum erro durante a iteração
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("erro ao iterar sobre os resultados: %v", err)
	}

	// Serializar os resultados em JSON
	jsonData, err := json.Marshal(resultados)
	if err != nil {
		return "", fmt.Errorf("erro ao transformar em JSON: %v", err)
	}

	return string(jsonData), nil
}

func ExecutarConsultaById(consulta string, id int) (string, error) {
	// Verifica se a conexão com o banco de dados foi estabelecida
	if db == nil {
		return "", fmt.Errorf("conexão com o banco de dados não foi estabelecida")
	}
	param := strconv.Itoa(id)

	// Executar a consulta SELECT
	rows, err := db.Query(consulta + " where id  = " + param)
	if err != nil {
		return "", fmt.Errorf("erro ao executar a consulta: %v", err)
	}
	defer rows.Close()

	// Obter as colunas do resultado
	columns, err := rows.Columns()
	if err != nil {
		return "", fmt.Errorf("erro ao obter colunas: %v", err)
	}

	// Criar uma slice para armazenar os valores de cada linha
	var resultados []map[string]interface{}

	// Iterar sobre os resultados
	for rows.Next() {
		// Criar um slice para armazenar os valores de cada coluna na linha
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		// Escanear os valores da linha
		if err := rows.Scan(scanArgs...); err != nil {
			return "", fmt.Errorf("erro ao escanear os resultados: %v", err)
		}

		// Criar um mapa para armazenar os valores de cada linha
		linha := make(map[string]interface{})
		for i, col := range columns {
			linha[col] = values[i]
		}

		// Adicionar o mapa de valores à slice de resultados
		resultados = append(resultados, linha)
	}

	// Verificar se ocorreu algum erro durante a iteração
	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("erro ao iterar sobre os resultados: %v", err)
	}

	// Serializar os resultados em JSON
	jsonData, err := json.Marshal(resultados)
	if err != nil {
		return "", fmt.Errorf("erro ao transformar em JSON: %v", err)
	}

	return string(jsonData), nil
}
