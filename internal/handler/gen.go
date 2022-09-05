package handler

import (
	"bytes"
	_ "embed"
	"go/format"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/clantable/jukslv2-generator/config"
	"github.com/iancoleman/strcase"
)

//go:embed template/template.txt
var templateTxt string

func Gen(c config.Config, model string) {
	params := map[string]any{
		"model":   strcase.ToLowerCamel(model),
		"Model":   strcase.ToCamel(model),
		"service": strcase.ToLowerCamel(c.ServiceName),
		"Service": strcase.ToCamel(c.ServiceName),
		"package": strcase.ToSnake(c.HandlerConfig.Package),
	}
	var buf bytes.Buffer
	t := template.Must(template.New("meta-txt").Parse(templateTxt))
	t.Execute(&buf, params)

	var out bytes.Buffer
	out.Write(buf.Bytes())
	body, err := format.Source(out.Bytes())
	if err != nil {
		// The user can compile the output to see the error.
		log.Fatalf("warning: internal error: invalid Go generated: %s", err)
	}
	if err := ioutil.WriteFile(c.HandlerConfig.Path+"/"+model+".go", body, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
