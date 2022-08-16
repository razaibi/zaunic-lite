WORKSHEET_FOLDER=projects/$1/worksheets
mkdir -p $WORKSHEET_FOLDER
WORKSHEET_PATH=$WORKSHEET_FOLDER/sample-ws.yml

touch $WORKSHEET_PATH

echo "---
items:
  - name: \"Generate File.\"
    data: \"sample-data.yml\"
    template: \"sample.go.tmpl\"
    output: \"sample.html\"
" >> $WORKSHEET_PATH