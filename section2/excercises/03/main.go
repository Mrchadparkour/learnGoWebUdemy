package main

import (
  "os"
  "encoding/csv"
  "html/template"
  "strconv"
  "time"
  "log"
  "net/http"
)


type Record struct {
  Date time.Time
  Open float64
}

type Table struct {
  Headers []string
  Rows []Record
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
  table := csvParse("table.csv")

  tpl, err := template.ParseFiles("tpl.gohtml")
  if err != nil {
    log.Fatalln(err)
  }

  //writes to res -> webpage -> imports records into tpl.gohtml
  err = tpl.Execute(res, table)
  if err != nil {
    log.Fatalln(err)
  }
}

func csvParse(filePath string) Table {
  //returns a file(*file) if successful
  src, err := os.Open(filePath)
  if err != nil {
    log.Fatalln(err)
  }
  defer src.Close()
  //returns a poiinter to a reader
  rdr := csv.NewReader(src)

  //returns slices of fields
  rows, err := rdr.ReadAll()
  if err != nil {
    log.Fatalln(err)
  }

  //creates a empty slice of Record, length 0, with a capacity of the length of rows
  records := make([]Record, 0, len(rows))
  headers := make([]string, 0, len(rows))

  for i, row := range rows {
    //append to headers then skip
    if i == 0 {
      headers = row[:2]
      continue
    }

    //format date and time from column 0, and 1
    date, _ := time.Parse("2006-01-02", row[0])
    open, _ := strconv.ParseFloat(row[1], 64)

    //sets records to be records with a new Record appended
    records = append(records, Record{
      Date: date,
      Open: open,
    })
  }

  return Table {
    headers,
    records,
  }
}
