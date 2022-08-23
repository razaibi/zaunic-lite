TEMP_FOLDER=_projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_GO_PATH=$TEMP_FOLDER/sample.go.tmpl

# Add Go Template
touch $TEMP_GO_PATH

echo 'Single Item: {{.data.name}}
Item List:
{{range  .data.items}}
	- {{.name}}
{{end}}
' >> $TEMP_GO_PATH