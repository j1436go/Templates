package templates

import (
	"fmt"
	html "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	text "text/template"

	"bitbucket.org/j1436go/inflect"
)

type TemplateType int

const (
	HTMLTemplate = iota
	TextTemplate
)

// Templater is an interface for both text/template
// and html/template.
// Use type hinting to access template specific members.
type Templater interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

func ParseRecursive(root string, typ TemplateType, exts ...string) (Templater, error) {
	var paths []string

	filepath.Walk(root, func(
		path string,
		info os.FileInfo,
		err error) error {
		if strInSlice(filepath.Ext(path), exts) {
			paths = append(paths, path)
		}
		return nil
	})

	switch typ {
	case HTMLTemplate:
		return html.ParseFiles(paths...)
	case TextTemplate:
		return text.ParseFiles(paths...)
	}

	return nil, fmt.Errorf("Undefined template type.")
}

func strInSlice(str string, sl []string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}

// HelperFuncMap provides all string functions in the strings package
// along with simple arithmetic and inflection functions.
var TemplateFuncs = map[string]interface{}{
	"Contains":       strings.Contains,
	"ContainsAny":    strings.ContainsAny,
	"ContainsRune":   strings.ContainsRune,
	"Count":          strings.Count,
	"EqualFold":      strings.EqualFold,
	"Fields":         strings.Fields,
	"HasPrefix":      strings.HasPrefix,
	"HasSuffix":      strings.HasSuffix,
	"Index":          strings.Index,
	"IndexAny":       strings.IndexAny,
	"IndexByte":      strings.IndexByte,
	"IndexFunc":      strings.IndexFunc,
	"IndexRune":      strings.IndexRune,
	"Join":           strings.Join,
	"LastIndex":      strings.LastIndex,
	"LastIndexAny":   strings.LastIndexAny,
	"Map":            strings.Map,
	"Repeat":         strings.Repeat,
	"Replace":        strings.Replace,
	"Split":          strings.Split,
	"SplitAfter":     strings.SplitAfter,
	"SplitAfterN":    strings.SplitAfterN,
	"SplitN":         strings.SplitN,
	"Title":          strings.Title,
	"ToLower":        strings.ToLower,
	"ToLowerSpecial": strings.ToLowerSpecial,
	"ToTitle":        strings.ToTitle,
	"ToTitleSpecial": strings.ToTitleSpecial,
	"ToUpper":        strings.ToUpper,
	"ToUpperSpecial": strings.ToUpperSpecial,
	"Trim":           strings.Trim,
	"TrimFunc":       strings.TrimFunc,
	"TrimLeft":       strings.TrimLeft,
	"TrimLeftFunc":   strings.TrimLeftFunc,
	"TrimPrefix":     strings.TrimPrefix,
	"TrimRight":      strings.TrimRight,
	"TrimRightFunc":  strings.TrimRightFunc,
	"TrimSpace":      strings.TrimSpace,
	"TrimSuffix":     strings.TrimSuffix,

	"Asciify":             inflect.Asciify,
	"Camelize":            inflect.Camelize,
	"CamelizeDownFirst":   inflect.CamelizeDownFirst,
	"Capitalize":          inflect.Capitalize,
	"Dasherize":           inflect.Dasherize,
	"ForeignKey":          inflect.ForeignKey,
	"ForeignKeyCondensed": inflect.ForeignKeyCondensed,
	"Humanize":            inflect.Humanize,
	"Ordinalize":          inflect.Ordinalize,
	"Parameterize":        inflect.Parameterize,
	"ParameterizeJoin":    inflect.ParameterizeJoin,
	"Pluralize":           inflect.Pluralize,
	"Singularize":         inflect.Singularize,
	"Tableize":            inflect.Tableize,
	"Titleize":            inflect.Titleize,
	"Typeify":             inflect.Typeify,
	"Uncountables":        inflect.Uncountables,
	"Underscore":          inflect.Underscore,

	"Sprintf": fmt.Sprintf,

	"Incr": func(a int64) int64 { return a + 1 },
	"Decr": func(a int64) int64 { return a - 1 },
}
