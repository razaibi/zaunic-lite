TEMP_FOLDER=templates/$1
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample_template.

touch $TEMP_PATH

echo '
    {{.data.name}}
	{{range  .data.items}}
		{{.name}}
	{{end}}
' >> $TEMP_PATH