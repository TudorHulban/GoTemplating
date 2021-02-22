package page

const (
	PageStart = `
<!DOCTYPE html>
<html>
`

	PageEnd = `
</html>
`

	HEAD = `
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="keywords" content="{{.Site.Subtitle}}">
<meta name="description" content="{{.Site.Subtitle}}">
<title>{{.Site.Title}}</title>
</head>
`

	BODYStart = `
<body>
`

	BODYEnd = `
</body>
`

	FOOTER = `
<footer>
<p>Copyright 2020 TARA Web</p>
<p>Contact information: <a href="mailto:admin@taraweb.eu">admin@taraweb.eu</a>.</p>
</footer>
`
)
