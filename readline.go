package readline
import (
  "bufio"
  "os"
  "strconv"
  "strings"
  "github.com/diebels727/vertex"
  "github.com/diebels727/edge"
)

func ReadEdges(path string) ([]edge.Edge,error) {
  //open input file
  file,err := os.Open(path)
  if err != nil { return nil,err }

  var edges []edge.Edge

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {

    fields := strings.Fields(scanner.Text())
    i0,err_i0 := strconv.Atoi(fields[0])
    i1,err_i1 := strconv.Atoi(fields[1])
    if err_i0 != nil { return nil,err_i0}
    if err_i1 != nil { return nil,err_i1}
    v0 := vertex.Vertex{Id:i0}
    v1 := vertex.Vertex{Id:i1}
    e := edge.Edge{V0:v0,V1:v1}
    edges = append(edges,e)
  }
  return edges,nil
}

func Map(path string,fn func(f string)) {
  //open input file
  file,_ := os.Open(path)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fn(scanner.Text())
  }
  return
}
