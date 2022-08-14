DATA_FOLDER=data/$1
mkdir -p $DATA_FOLDER
DATA_PATH=$DATA_FOLDER/sample_data.yml

touch $DATA_PATH

echo '
---
data:
  name: "Some Item"
  items:
    - name: "someitem"
      place: "someotherplace"
' >> $DATA_PATH