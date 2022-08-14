### Zaunic Lite

This is Golang based code generator. This is a light weight version that uses _yml_ files and _go.tmpl_ files to generate code.

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

2. Setup template for generation in *templates* folder in _go.tmpl_ format.

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

3. Create Worksheet to execute/generate in *worksheets* folder (in _.yml_ format). Worksheet drives the sequence of execution of the code generated.

Sample:
```yml
---
items:
  - name: "Generate Progam.cs."
    data: "sample-data.yml"
    template: "sample.go.tmpl"
    output: "sample.html"
```

The output value will be appended to the base path of the _output_ folder.

4. In *main.go*, setup the worksheet name you want to run.

Sample:
```go
core.ProcessWorksheet("<your-worksheet")
```

5. Simply run using:

```bash
go run main.go
```
