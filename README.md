### Zaunic Lite

This is Golang based code generator. This is a light weight version that uses _yml_ files and _go.tmpl_ files to generate code.

### Steps to Setup.

1. Create your new project using the script. 

```bash
./scripts/add_project.sh <name-of-your-project>
```

2. Setup Data for Generation in *_projects/<name-of-your-project>/data* folder in yaml format.

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

3. Setup template for generation in *_projects/<name-of-your-project>/templates* folder. You can choose between the *go templating engine* (_go.tmpl_ format) and *liquid templating* (_.liquid_ format). 

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


4. Create Worksheet to execute/generate in *_projects/<name-of-your-project>/worksheets* folder (in _.yml_ format). Worksheet drives the sequence of execution of the code generated.

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

> Currently the following template types are supported:

- *Go Templates* (.go.tmpl), 
- *Liquid* (.liquid)
- *Pongo* (.pongo)
- *Mustache* (.mustache)

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

#### Secrets Management

- Zaunic-Lite supports pulling secrets from cloud and embedding it into your generated code.

- To do this, your environment variables are used.

- Let us take an example to illustrate. If you have your AWS access key id and secret key stored as environment variables, you can use them to pull up your secrets like below.

```yml
secrets:
  - name: "some-awesome-secret"
    source: "aws"
    env: "dev"
    region: "us-east-1"
```

Core will start looking environment variables.

- Ensure these environment variables (in all caps) are set based on the cloud provider.

For AWS, 

```bash
AWS_<your-environment-name>_ACCESS_KEY_ID
AWS_<your-environment-name>_SECRET_ACCESS_KEY
```

#### Useful templating plugins for VSCode.

- [Go Lang Template - go.tmpl](https://marketplace.visualstudio.com/items?itemName=jinliming2.vscode-go-template)

- [Shopify Liquid - .liquid](https://marketplace.visualstudio.com/items?itemName=Shopify.theme-check-vscode)

- [Mustache Template - .mustache](https://marketplace.visualstudio.com/items?itemName=dawhite.mustache)

#### Other Templating engine links.

- [Pongo - .pongo](https://github.com/flosch/pongo2)
