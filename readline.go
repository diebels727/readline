package readline
import (
  "bufio"
  "os"
  "strconv"
)

func Ints(path string) ([]int,error) {
  //open input file
  file,err := os.Open(path)
  if err != nil { return nil,err }
  var ints []int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    i,err := strconv.Atoi(scanner.Text())
    if err != nil { return nil,err }
    ints = append(ints,i)
  }
  return ints,nil
}

func AdjacencyListInts(path string) {
  //open input file
  file,err := os.Open(path)
  if err != nil {return nil,err }
  var ints []int
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    i,err := strconf.Atoi(scanner.Text())
    if error != nil { return nil,err }
  }
}

func main() {
  AdjacencyListInts('kargerMinCut.txt')
}

