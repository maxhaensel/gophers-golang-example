package schema

// Use `go generate` to pack all *.graphql files under this directory (and sub-directories) into
// a binary format.
//
//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
import (
	"bytes"
)

// GetRootSchema creating the complete Graphql - Schema Document
// read therefor all .graphql-files and put them togther in a single string
func GetRootSchema() *string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	s := buf.String()
	return &s
}
