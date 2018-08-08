package md

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsac44302fb6150a621aa9d04a0350aac972bf7e18 = "# {{ .Table.Name }}\n\n## Description\n\n{{ .Table.Comment | nl2mdnl -}}\n{{ if .Table.Def }}\n<details>\n<summary><strong>Table Definition</strong></summary>\n\n```sql\n{{ .Table.Def }}\n```\n\n</details>\n{{ end }}\n\n## Columns\n{{ range $l := .Columns }}\n|{{ range $d := $l }} {{ $d | nl2br }} |{{ end }}\n{{- end }}\n\n{{ $len := len .Constraints }}{{ if ne $len 2 -}}\n## Constraints\n{{ range $l := .Constraints }}\n|{{ range $d := $l }} {{ $d | nl2br }} |{{ end }}\n{{- end }}\n{{- end }}\n\n{{ $len := len .Indexes -}}{{ if ne $len 2 -}}\n## Indexes\n{{ range $l := .Indexes }}\n|{{ range $d := $l }} {{ $d | nl2br }} |{{ end }}\n{{- end }}\n{{- end }}\n\n{{ $len := len .Triggers -}}{{ if ne $len 2 -}}\n## Triggers\n{{ range $l := .Triggers }}\n|{{ range $d := $l }} {{ $d | nl2br }} |{{ end }}\n{{- end }}\n{{- end }}\n\n{{ if .er -}}\n## Relations\n\n![er]({{ .Table.Name }}.png)\n{{- end }}\n\n---\n\n> Generated by [tbls](https://github.com/k1LoW/tbls)"
var _Assets43889384df1c6f74d764c29d91b9d5637eb46061 = "# {{ .Schema.Name }}\n\n## Tables\n{{ range $t := .Tables }}\n|{{ range $d := $t }} {{ $d | nl2br }} |{{ end }}\n{{- end }}\n\n{{ if .er -}}\n## Relations\n\n![er](schema.png)\n{{- end }}\n\n---\n\n> Generated by [tbls](https://github.com/k1LoW/tbls)"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.md.tmpl", "table.md.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1532785399, 1532785399000000000),
		Data:     nil,
	}, "/index.md.tmpl": &assets.File{
		Path:     "/index.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1532239511, 1532239511000000000),
		Data:     []byte(_Assets43889384df1c6f74d764c29d91b9d5637eb46061),
	}, "/table.md.tmpl": &assets.File{
		Path:     "/table.md.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1532785399, 1532785399000000000),
		Data:     []byte(_Assetsac44302fb6150a621aa9d04a0350aac972bf7e18),
	}}, "")
