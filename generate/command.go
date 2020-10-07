package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/FINTLabs/fint-jsonschema-generator/common/config"
	"github.com/FINTLabs/fint-jsonschema-generator/common/github"
	"github.com/FINTLabs/fint-jsonschema-generator/common/parser"
	"github.com/FINTLabs/fint-jsonschema-generator/common/types"
	"github.com/urfave/cli"
)

// CmdGenerate implements the `generate` subcommand.
func CmdGenerate(c *cli.Context) {

	var tag string
	if c.GlobalString("tag") == config.DEFAULT_TAG {
		fmt.Print("Getting latest from GitHub...")
		tag = github.GetLatest(c.GlobalString("owner"), c.GlobalString("repo"))
		fmt.Printf(" %s\n", tag)
	} else {
		tag = c.GlobalString("tag")
	}
	force := c.GlobalBool("force")
	owner := c.GlobalString("owner")
	repo := c.GlobalString("repo")
	filename := c.GlobalString("filename")

	classes, _, _, _ := parser.GetClasses(owner, repo, tag, filename, force)

	setupJSONSchemaDirStructure()
	generateJSONSchema(classes)

	fmt.Println("Done!")
}

func writeFile(path string, filename string, content []byte) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path+"/"+filename, content, 0777)
}

func includePackage(p string) bool {
	//return strings.Contains(p, "administrasjon") || strings.Contains(p, "utdanning") || strings.Contains(p, "felles")
	return true
}

func setupJSONSchemaDirStructure() {
	os.RemoveAll(config.JSON_BASE_PATH)
	os.Mkdir(config.JSON_BASE_PATH, 0777)
}

func writeJSONSchema(pkg string, schema string, content []byte) error {
	path := fmt.Sprintf("%s/schema/%s", config.JSON_BASE_PATH, types.GetComponentName(pkg))
	return writeFile(path, strings.ToLower(schema)+".json", []byte(content))
}

func generateJSONSchema(classes []*types.Class) {

	fmt.Println("Generating JSON Schema")

	var roots []*types.Class

	for _, c := range classes {
		if !c.Abstract && includePackage(c.Package) {
			fmt.Printf("  > Creating schema: %s.json\n", c.Name)
			schema := GetJSONSchema(c)
			err := writeJSONSchema(c.Package, c.Name, []byte(schema))
			if err != nil {
				fmt.Printf("Unable to write file: %s", err)
			}
			if c.Stereotype == "hovedklasse" && c.Identifiable && !strings.Contains(c.Package, "kodeverk") {
				include := false
				for _, i := range c.Identifiers {
					include = include || !i.Optional
				}
				if include {
					roots = append(roots, c)
				}
			}
		}
	}

}
