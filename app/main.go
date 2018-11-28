package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSourceName := os.Getenv("DATASOURCENAME")
	if dataSourceName == "" {
		dataSourceName = "root:password@tcp(127.0.0.1:13306)/sampledb"
	}
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// health check
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <title>サンプルアプリ</title>
</head>
<body>
  <h1>1day サンプルアプリ</h1>
  <a href="https://voyagegroup.com/" alt="VOYAGE GROUP"><img src="https://voyagegroup.com/wp-content/themes/voyagegroup/common/img/img_ogp_cp.png" width="400"></a>

  <h2>DB接続確認リンク</h2>
  <ul>
    <li><a href="event?name=hoge&value=fuga">event?name=hoge&value=fuga</a>
  </ul>
</body>
</html>
		`))
	})

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		stmt, err := db.Prepare("INSERT INTO eventlog(at, name, value) values(NOW(), ?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		name := r.URL.Query().Get("name")
		value := r.URL.Query().Get("value")
		_, _ = stmt.Exec(name, value)

		w.WriteHeader(200)
		w.Write([]byte("success"))
	})

	// start server
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
