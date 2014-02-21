package main
import (
  "io"
  "os"
  // "strconv"
  "fmt"
  "encoding/csv"
  // "reflect"
)

func AdjacencyListInts(path string) {
  //open input file
  file,_ := os.Open(path)
  defer file.Close()
  reader := csv.NewReader(file)
  reader.TrailingComma = true
  var adj map[int][]string
  adj = make(map[int][]string)

  i := 0
  for {
    var err error
    fmt.Println(i)
    adj[i],err = reader.Read()
    // record,err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("Error:",err)
      return
    }
    // fmt.Println(i)
    fmt.Println(adj[i])
    i = i+1
  }

}

func main() {
  AdjacencyListInts("kargerMinCut.csv")
}