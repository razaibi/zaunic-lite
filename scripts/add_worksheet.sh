WORKSHEET_FOLDER=worksheets/$1
mkdir -p $WORKSHEET_FOLDER
WORKSHEET_PATH=$WORKSHEET_FOLDER/sample_ws.yml

touch $WORKSHEET_PATH

echo "---
items:
  - name: \"Generate File.\"
    data: \"$1\sample-data.yml\"
    template: \"$1\sample.go.tmpl\"
    output: \"$1.html\"
" >> $WORKSHEET_PATH