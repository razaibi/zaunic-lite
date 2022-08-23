mkdir -p _projects/$1
mkdir -p _projects/$1/data
mkdir -p _projects/$1/templates
mkdir -p _projects/$1/worksheets
mkdir -p _projects/$1/output

./scripts/add_data.sh $1
./scripts/add_go_template.sh $1
./scripts/add_liquid_template.sh $1
./scripts/add_worksheet.sh $1