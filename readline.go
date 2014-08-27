package readline
import (
  "bufio"
  "os"
)

type Readline struct {
  File *os.File
}

func (r *Readline) Map(fn func(f string)) {
  scanner := bufio.NewScanner(r.File)
  for scanner.Scan() {
    fn(scanner.Text())
  }
  return
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
