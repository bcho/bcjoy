package render

import (
	"io"
	"io/ioutil"

	"github.com/bcho/bcjoy/files"
	"github.com/gin-gonic/contrib/renders/multitemplate"
)

func Make() multitemplate.Render {
	r := multitemplate.New()

	mustAddFromTemplateFile(r, "join", "join.html")

	return r
}

func mustReadStringFromTemplateFile(r io.ReadCloser) string {
	defer r.Close()
	if b, err := ioutil.ReadAll(r); err != nil {
		panic(err)
	} else {
		return string(b)
	}
}

func mustAddFromTemplateFile(r multitemplate.Render, name, templateName string) {
	template, err := files.Open(templateName)
	if err != nil {
		panic(err)
	}

	r.AddFromString(name, mustReadStringFromTemplateFile(template))
}
