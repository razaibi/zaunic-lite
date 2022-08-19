TEMP_FOLDER=projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_PATH=$TEMP_FOLDER/sample.pongo

# Add Liquid Template
touch $TEMP_PATH

echo 'Single Item: {{ data.name }}
Item List:
{% for item in data.items %}
    - {{ item.name }}
{% endfor %}
' >> $TEMP_PATH