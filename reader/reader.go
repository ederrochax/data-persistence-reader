package reader

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"reader/utils"
	"strings"
	"sync"
)

func PersistentScanner(file *os.File, conn *sql.DB) error {
	scanner := bufio.NewScanner(file)

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
			continue
		}
		line := scanner.Text()
		line = strings.TrimSpace(line)
		columns := strings.Fields(line)

		columnValues := utils.StringsValidator(columns)

		wg.Add(1)
		go func() {
			defer wg.Done()
			stmt, err := tx.Prepare("INSERT INTO customer (cpf, private, incomplete, date_last_purchase, ticket_average, ticket_last_purchase, cnpj_most_frequent_store, cnpj_last_purchase_store) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
			if err != nil {
				log.Println(err)
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(columnValues[0], columnValues[1], columnValues[2], columnValues[3], columnValues[4], columnValues[5], columnValues[6], columnValues[7])
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}
	wg.Wait()

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
