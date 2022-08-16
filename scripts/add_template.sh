TEMP_FOLDER=projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample.go.tmpl

touch $TEMP_PATH

echo 'Single Item: {{.data.name}}
Item List:
{{range  .data.items}}
	- {{.name}}
{{end}}
' >> $TEMP_PATH