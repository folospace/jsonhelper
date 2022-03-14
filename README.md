# json helper

## fill_default_to_nil.go -- when json marshal, with fill default to nil (recursive)
- *string => ""
- *int => 0
- []int, []string => []

#### how to use
```go
 bytes, err := json.Marshal(FillDefaultToNil(struct {A *string}{}))
 bytes, err := json.Marshal(FillDefaultToNil(anyVal))
```
