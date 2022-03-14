package jsonhelper

import (
    "encoding/json"
    "testing"
)

/**
 * @Author zyq
 * @Date 2021/7/7
 * @Description
 **/
func TestFillDefaultToNil(t *testing.T) {
    type myint int

    type Test struct {
        A1 *string        `json:"a_1"`
        A2 *myint         `json:"a_2"`
        A3 map[int]*myint `json:"a_3"`
        A4 []*myint       `json:"a_4"`
        A5 *Test          `json:"a_5"`
    }

    ex := &Test{A3: map[int]*myint{1: nil}, A4: []*myint{nil}}

    jb, _ := json.Marshal(FillDefaultToNil(ex))

    t.Log(string(jb))
    //{"a_1":"","a_2":0,"a_3":{"1":0},"a_4":[0],"a_5":null}
}
