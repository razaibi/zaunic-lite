TEMP_FOLDER=_projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample.hbs

# Add Liquid Template
touch $TEMP_PATH

echo 'Single Item: {{data.name}}
Item List:
{{#each data.items}}
    - {{this.name}}
{{/each}}
' >> $TEMP_PATH