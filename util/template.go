package util

const Template = `{{ $ip := .Ip }}
{{ with .Hostsfile }}# < {{ .Name }}
{{ range .Hosts }}{{ $ip }} {{ range .Domains }}{{ . }} {{ end }}# {{ .Comment }}
{{ end -}}
# {{ .Name }} >{{ end }}
`

type TemplateData struct {
	Ip        string
	Hostsfile Hostsfile
}
