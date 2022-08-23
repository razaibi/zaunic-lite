DATA_FOLDER=_projects/$1/data
mkdir -p $DATA_FOLDER
DATA_PATH=$DATA_FOLDER/sample-data.yml

touch $DATA_PATH

echo '---
data:
  name: "Some Name Value"
  items:
    - name: "First Item Name"
      place: "First Item Place"
    - name: "Second Item Name"
      place: "Second Item Place"
' >> $DATA_PATH