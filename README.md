### Zaunic Lite

This is Golang based code generator.

### Steps to Setup.

1. Setup Data for Generation in *data* folder in yaml format.

Sample:
```yml
data:
  name: "Raz"
  options:
    - item: "somecontent"
      points:
        - name: "abcde"
```

2. Setup template for generation in *templates* folder in go.tmpl format.

Sample:
```tmpl
<html>
	<head>
		<title>Sample HTML Generation</title>
	</head>
	<body>
    Name: {{.data.name}}
    Something Other Value: {{.somethingElse}}

	{{range  .data.options}}
		<li>Product Price : {{.item}}</li>
		{{range  .points}}
			<li>Product Price : {{.name}}</li>
		{{end}}
	{{end}}

	</body>
</html>

```

3. Create Worksheet to execute/generate in worksheets folder (in yaml format).

Sample:
```yml
---
items:
  - name: "Generate Progam.cs."
    data: "sample-data.yml"
    template: "sample.go.tmpl"
    output: "sample.html"
```

4. In *main.go*, setup the worksheet name you want to run.

Sample:
```go
    core.ProcessWorksheet("<your-worksheet")
```

5. Simply run using:

```bash
go run main.go
```
