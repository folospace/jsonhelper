# json helper

## fill_default_to_nil.go -- when json marshal, with fill default to nil (recursive)
- *string => ""
- *int => 0
- []int, []string, *[]string => []
- map[int]*string = {}

#### how to use
```go
 import "github.com/folospace/jsonhelper"

 bytes, err := json.Marshal(FillDefaultToNil(struct {A *string}{}))
 fmt.Println(string(bytes)) // {"A":""}
```
