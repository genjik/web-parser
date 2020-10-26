package main

import (
    "strings"
    "reflect"
    "fmt"
    "io"
)

type website struct {
    Url string
    Elements []el
}

type el struct {
    Tag string
    Limit int
    Attrs []attr
}

type attr struct {
    Key string
    Val string
}

type fetchedWebsite struct {
    website
    Body io.Reader
}

type result struct {
    Url string `json:"url"`
    Data map[string][]string `json:"data"`
}

func (w *website) isEqualTo(w2 *website) bool {
    if strings.EqualFold(w.Url, w2.Url) == false {
        return false
    }

    if len(w.Elements) != len(w2.Elements) {
        return false
    }

    for i, el := range w.Elements {
        if el.isEqualTo(&w2.Elements[i]) == false {
            return false
        }
    }

    return true
}

func (e *el) isEqualTo(e2 *el) bool {
    if strings.EqualFold(e.Tag, e2.Tag) == false {
        return false
    }

    if e.Limit != e2.Limit {
        return false
    }
    
    if len(e.Attrs) != len(e2.Attrs) {
        return false
    }

    for i:=0; i < len(e.Attrs); i++ {
        if strings.EqualFold(e.Attrs[i].Key, e2.Attrs[i].Key) == false {
            return false
        }
        if strings.EqualFold(e.Attrs[i].Val, e2.Attrs[i].Val) == false {
            return false
        }
    }

    return true
}

func (r *result) isEqualTo(r2 *result) bool {
    if r.Url != r2.Url {
        return false
    }

    if reflect.DeepEqual(r.Data, r2.Data) == false {
        return false
    }

    return true
}

func (e *el) getAttrs() []string {
    var result []string

    for _, attr := range e.Attrs {
        if attr.Key == "" || attr.Val == "" {
            continue
        }
        result = append(result, attr.Key, attr.Val)    
    }

    return result
}

func (e *el) getKeys() string {
    var result strings.Builder

    for _, attr := range e.Attrs {
        if attr.Key == "" || attr.Val == "" {
            continue
        }
        result.WriteString(fmt.Sprintf("%s ", attr.Val))
    }
    
    return strings.TrimSuffix(result.String(), " ")
}

func compareStr(s1, s2 []string) bool {
    if len(s1) != len(s2) {
        return false
    }

    for i, v := range s1 {
        if strings.EqualFold(v, s2[i]) == false {
            return false
        }
    }

    return true
}
