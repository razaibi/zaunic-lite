TEMP_FOLDER=projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample.liquid

# Add Liquid Template
touch $TEMP_PATH

echo 'Single Item: {{data.name}}
Item List:
{{#data.items}}
    - {{name}}

{{/data.items}}
' >> $TEMP_PATH