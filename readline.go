package readline
import (
  "bufio"
  "os"
  "strconv"
  "strings"
)

func Map(path string,fn func(f string)) {
  //open input file
  file,_ := os.Open(path)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fn(scanner.Text())
  }
  return
}
