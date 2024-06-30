package generate

import (
	"fmt"
	"os"
	"strings"
)

type IPathSchema interface {
	ToYaml() (string, error)
}

func GeneratePathYamlFile(p IPathSchema, path, filename string) error {
	yaml, err := p.ToYaml()
	if err != nil {
		return err
	}

	os.MkdirAll(path, 0755)

	allName := func(fileName string) string {
		if strings.Contains(fileName, ".yaml") {
			return fmt.Sprintf("%s/%s", path, fileName)
		}

		return fmt.Sprintf("%s/%s.yaml", path, fileName)
	}(filename)

	return os.WriteFile(allName, []byte(yaml), 0644)
}
