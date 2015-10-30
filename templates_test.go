package templates

import (
	"html/template"
	"os"
	"testing"
)

func TestParseRecursive(t *testing.T) {
	tmpls, err := ParseRecursive("./test", HTMLTemplate, ".html", ".tmpl")
	if err != nil {
		t.Errorf("Expected template parsing to succeed, err was %s", err)
		return
	}
	_, ok := tmpls.(*template.Template)
	if !ok {
		t.Error("Expected to cast Templater to html/template, failed.")
		return
	}
}

func TestHelperFuncs(t *testing.T) {
	tmpl, err := template.New("func").Funcs(TemplateFuncs).ParseFiles("./test/foo/bar/exec.func")
	if err != nil {
		t.Errorf("Expected template parsing to succeed, err was %s", err)
		return
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		t.Errorf("Expected tmpl executing func to success, err was \n%s", err)
	}
}
