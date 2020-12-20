package main
import(
"github.com/boltdb/bolt"
"fmt"
"os"
)

func ExternalDBListOfBuckets(db *bolt.DB) []string{
var list_of_strings []string
db.View(func(tx *bolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
            list_of_strings=append(list_of_strings,string(name))
            return nil
        })
    })
return list_of_strings
}

func SqlCreateTable(bucket string) string{
return "create table if not exists "+bucket+"(k text primary key unique,v text);"
}
func SqlInsert(bucket,key,value string) string{
return `insert into `+bucket+`(k,v) values("`+key+`","`+value+`");`
}

func ExternaDBCreateSql(db *bolt.DB){
lista:=ExternalDBListOfBuckets(db)

for m:=range(lista){

fmt.Println(SqlCreateTable(lista[m]))
db.View(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte(lista[m]))
	c := b.Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
	fmt.Println(SqlInsert(lista[m],string(k),string(v)))
	}

	return nil
})
}


}

func main(){
arg := os.Args[1]
db, err := bolt.Open(arg, 0600, nil)
if err==nil{
ExternaDBCreateSql(db)
}
}
