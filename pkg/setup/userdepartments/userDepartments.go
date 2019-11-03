package userdepartments

import (
	"bufio"
	"encoding/csv"
	. "github.com/chasdevs/meetrics/pkg/data"
	"github.com/chasdevs/meetrics/pkg/util"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

func GetUsersFromFile() []*User {

	log.Infoln("Getting users and departments from csv file.")

	// Open the csv.
	filePath := path.Join(util.ThisFilePath(), "user-departments.csv")
	csvFile, e := os.Open(filePath)
	if csvFile != nil {
		defer func() {
			if e := csvFile.Close(); e != nil {
				panic(e)
			}
		}()
	}

	if e != nil {
		panic(e)
	}

	// Create a Reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))
	//reader.

	// Iterate through the lines and create an array of User objects.
	var users []*User
	lines := 0
	for {

		line, e := reader.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			panic(e)
		}

		lines++
		users = append(users, &User{
			Name:       line[0],
			Email:      line[1],
			Department: line[2],
		})

	}

	return users
}
