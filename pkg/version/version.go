package version

// Based on: https://github.com/prometheus/common/blob/master/version/info.go

import (
	"bytes"
	"html/template"
	"runtime"
	"strings"
)

var (
	Version   string
	BuildDate string
	GoVersion = runtime.Version()
)

var versionTmpl = `
pwaudit, version {{.version}}
  build date:   {{.buildDate}}  
  go version:   {{.goVersion}}  
`

func Print() string {
	m := map[string]string{
		"version":   Version,
		"buildDate": BuildDate,
		"goVersion": GoVersion,
	}
	t := template.Must(template.New("version").Parse(versionTmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", m); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}
