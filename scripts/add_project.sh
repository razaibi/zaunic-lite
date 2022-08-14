mkdir -p data/$1
mkdir -p templates/$1
mkdir -p worksheets/$1
mkdir -p output/$1

./scripts/add_data.sh $1
./scripts/add_template.sh $1
./scripts/add_worksheet.sh $1