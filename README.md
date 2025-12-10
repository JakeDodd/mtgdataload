## Requirements
1. Install GO and Postgres
    1. Create a Database User with a password, and a Database, if running locally
    2. Make sure that the User you create has the proper privileges on the database you create
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
5. Run the SQL script to set up all the tables and database objects
```
# login
psql -d $DATABASE -U $USER

# run script
\i $PROJECT_DIR/sql/createTables.sql
```
5. Build and run the project
```
go build
./mtgdataload
```
