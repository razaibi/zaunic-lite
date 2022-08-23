TEMP_FOLDER=_projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample.mustache

# Add Liquid Template
touch $TEMP_PATH

echo 'Single Item: {{data.name}}
Item List:
{{#data.items}}
    - {{name}}

{{/data.items}}
' >> $TEMP_PATH