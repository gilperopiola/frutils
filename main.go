package frutils

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/* gin */

func GetUserID(c *gin.Context) int {
	idUser, _ := c.Get("ID")
	return ToInt(idUser.(string))
}

func GetToken(c *gin.Context) string {
	return strings.Trim(strings.TrimSuffix(c.Request.Header.Get("Authorization"), "\n"), `"`)
}

func GetRequestIP(c *gin.Context) string {
	return c.Request.RemoteAddr
}

/* type conversions */

func ToString(i int) string {
	return strconv.Itoa(i)
}

func ToString64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func ToBool(s string) bool {
	return strings.TrimSpace(s) == "true" || strings.TrimSpace(s) == "1" || strings.TrimSpace(s) == "TRUE"
}

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateToStringFull(date time.Time) string {
	return date.Format("2006-01-02T15:04:05-07:00")
}

func IntToAlphabetPosition(i int) string {
	return string('A' - 1 + i)
}

/* dates */
func GetDaysBetween(date1 time.Time, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}

/* string manipulations */

func RemoveAllCharactersBefore(str string, separator string) string {
	hasAppeared := false
	result := ""

	for _, c := range str {
		if string(c) == separator {
			hasAppeared = true
		}

		if hasAppeared {
			result += string(c)
		}
	}

	return result
}

func RemoveNewLinesAndWhiteSpace(str string) string {
	return strings.Trim(strings.Trim(str, "\n"), `"`)
}

func RemoveExteriorSquareBrackets(str string) string {
	return strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(str, "["), "\n"), "]")
}

func GetFilenameExtension(filename string) string {
	splitted := strings.Split(filename, ".")
	return splitted[len(splitted)-1]
}

/* databases */

func GetID(result sql.Result) int {
	id, _ := result.LastInsertId()
	return int(id)
}

func GetQueryString(query string, args ...interface{}) string {
	var buffer bytes.Buffer
	nArgs := len(args)
	for i, part := range strings.Split(query, "?") {
		buffer.WriteString(part)
		if i < nArgs {
			switch a := args[i].(type) {
			case int64:
				buffer.WriteString(fmt.Sprintf("%d", a))
			case bool:
				buffer.WriteString(fmt.Sprintf("%t", a))
			case sql.NullBool:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%t", a.Bool))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullInt64:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%d", a.Int64))
				} else {
					buffer.WriteString("NULL")
				}
			case sql.NullString:
				if a.Valid {
					buffer.WriteString(fmt.Sprintf("%q", a.String))
				} else {
					buffer.WriteString("NULL")
				}
			default:
				buffer.WriteString(fmt.Sprintf("%q", a))
			}
		}
	}
	return buffer.String()
}
