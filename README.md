# golang-boltdb2sql
migrate boltdb to sql 

# why
migrate data from bucket/key-value to sql table/key/value
the mapping is :
bucket name == table name
key == field of a table called key
value == field of a table called value
# how
go build boltdb2sql.go

./boltdb2sql filename_boltdb.db > textfile.sql

