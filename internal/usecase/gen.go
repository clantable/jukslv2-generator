package usecase

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/clantable/jukslv2-generator/config"
	"github.com/iancoleman/strcase"
)

//go:embed template/template.txt
var templateTxt string

func Gen(c config.Config, model string) {
	params := map[string]any{
		"service":        strcase.ToLowerCamel(c.ServiceName),
		"domainPath":     strings.Split(c.DomainConfig.Path, ".")[1],
		"model":          strcase.ToLowerCamel(model),
		"Model":          strcase.ToCamel(model),
		"package":        strcase.ToSnake(c.UsecaseConfig.Package),
		"repositoryPath": strings.Split(c.RepositoryConfig.Path, ".")[1],
	}
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())

	// if you need format, uncomment this
	// body, err := format.Source(out.Bytes())
	// if err != nil {
	// 	// The user can compile the output to see the error.
	// 	log.Fatalf("warning: internal error: invalid Go generated: %s", err)
	// }
	if err := ioutil.WriteFile(c.UsecaseConfig.Path+"/"+model+".go", out.Bytes(), 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
