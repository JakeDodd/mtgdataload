## Requirements
### Install GO and Postgres
### Clone this repo
### Download all-cards file from [CARDS](https://scryfall.com/docs/api/bulk-data)
#### Place this file in the root directory of the project
#### This file is a few GB's in size
### Inside the root directory of the project, create a .env file
```
pg_host=localhost
pg_port=5432
pg_user=[Username that you created]
pg_password=[Password]
pg_dbname=[Database that you created]
data_filename="[File Name Downloaded above]"
```
### Build and run the project
```
go build
./mtgdataload
```
