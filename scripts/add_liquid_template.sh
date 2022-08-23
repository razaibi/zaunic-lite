TEMP_FOLDER=_projects/$1/templates
mkdir -p $TEMP_FOLDER
TEMP_LIQUID_PATH=$TEMP_FOLDER/sample.liquid

# Add Liquid Template
touch $TEMP_LIQUID_PATH

echo 'Single Item: {{ data.name }}
Item List:
{% for item in data.items %}
    - {{ item.name }}
{% endfor %}
' >> $TEMP_LIQUID_PATH