package utils

import (
	"database/sql"
	"github.com/Nhanderu/brdoc"
	"strings"
)

func StringsValidator(columns []string) []sql.NullString {
	columnValues := make([]sql.NullString, len(columns))

	for i, value := range columns {
		if value == "NULL" {
			columnValues[i].Valid = false
		} else {
			columnValues[i].Valid = true

			// Replace the comma with a period only for columns 4 and 5.
			if (i == 4 || i == 5) && strings.Contains(value, ",") {
				value = strings.ReplaceAll(value, ",", ".")
			}

			if i == 0 && !brdoc.IsCPF(value) {
				columnValues[i].Valid = false
			}

			if (i == 6 || i == 7) && !brdoc.IsCNPJ(value) {
				columnValues[i].Valid = false
			}

			columnValues[i].String = strings.ToUpper(value)
		}
	}
	return columnValues
}
