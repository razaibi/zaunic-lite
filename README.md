### Zaunic Lite

This is Golang based code generator. This is a light weight version that uses _yml_ files and _go.tmpl_ files to generate code.

### Steps to Setup.

1. Create your new project using the script. 

```bash
./scripts/add_project.sh <name-of-your-project>
```

2. Setup Data for Generation in *projects/<name-of-your-project>/data* folder in yaml format.

Sample:
```yml
data:
  name: "Raz"
  options:
    - item: "somecontent"
      points:
        - name: "abcde"
```

> Always start with data as the root element.

3. Setup template for generation in *projects/<name-of-your-project>/templates* folder. You can choose between the *go templating engine* (_go.tmpl_ format) and *liquid templating* (_.liquid_ format). 

Sample Go Template:
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

Sample Liquid Template:
```liquid
Single Item: {{ data.name }}
Item List:
{% for item in data.items %}
    - {{ item.name }}
{% endfor %}
```


4. Create Worksheet to execute/generate in *projects/<name-of-your-project>/worksheets* folder (in _.yml_ format). Worksheet drives the sequence of execution of the code generated.

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

##### Optional

If you intend to use templating engines apart from go templates, you can explicitly mention the engine as shown below:

```yml
---
items:
  - name: "Generate Progam.cs."
    data: "sample-data.yml"
    template: "sample.liquid"
    engine: "liquid"
    output: "sample-2.html"
```

> Currently only go templates (.go.tmpl) and liquid templates (.liquid) are supported.

5. In *main.go*, setup the worksheet name you want to run.

Sample:
```go
core.ProcessWorksheet("<your-project>","<your-worksheet>")
```

6. Simply run using:

```bash
go run main.go
```

#### Other commands to get started.

##### Adding new template data.
Add new project with the below script.

```bash
./scripts/add_data.sh <name-of-your-project>
```

##### Adding a new template.
Add new project with the below script.

```bash
./scripts/add_template.sh <name-of-your-project>
```

### Knowledge Hub

#### Useful templating plugins for VSCode.

- [Go Lang Template - go.tmpl](https://marketplace.visualstudio.com/items?itemName=jinliming2.vscode-go-template)
- [Shopify Liquid - .liquid](https://marketplace.visualstudio.com/items?itemName=Shopify.theme-check-vscode)
