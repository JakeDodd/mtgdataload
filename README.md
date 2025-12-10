## Requirements
1. Install GO and Postgres
2. Clone this repo
3. Download all-cards file from [CARDS](https://scryfall.com/docs/api/bulk-data)
    1. Place this file in the root directory of the project
    2. This file is a few GB's in size
4. Inside the root directory of the project, create a .env file
```
pg_host=localhost
pg_port=5432
pg_user=[Username that you created]
pg_password=[Password]
pg_dbname=[Database that you created]
data_filename="[File Name Downloaded above]"
```
5. Build and run the project
```
go build
./mtgdataload
```
