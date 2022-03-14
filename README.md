# json helper

## fill_default_to_nil.go -- when json marshal, with fill default to nil (recursive)
- *string => ""
- *int => 0
- []int, []string => []

#### how to use
- json.Marshal(FillDefaultToNil(struct {A *string}{}))
- json.Marshal(FillDefaultToNil(anyVal))