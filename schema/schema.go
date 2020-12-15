package schema

import (
	"bytes"
	"graphql/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetRootSchema creating the complete Graphql - Schema Document
// read therefor all .graphql-files and put them togther in a single string
func GetRootSchema() *string {
	buf := bytes.Buffer{}
	for _, name := range readAllGraphqlFiles() {
		b, _ := ioutil.ReadFile(name)
		buf.Write(b)
		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}
	s := buf.String()
	return &s
}

func readAllGraphqlFiles() []string {
	var results []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			results = append(results, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	results = utils.Filter(results, func(s string) bool {
		return strings.HasSuffix(s, ".graphql")
	})
	return results
}
