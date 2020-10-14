package main

import (
    "testing"
    "fmt"
)

func TestIsEqualToWebsite(t *testing.T) {
    cases := []struct {
        w website
        w2 website
        out bool
    }{
        {
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            true,
        },
        {
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            website{
                Url: "youtube.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            false,
        },
        {
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "span",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            false,
        },
        {
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "id", Val: "red"}, 
                        },
                    },
                },
            },
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            false,
        },
        {
            website{
                Url: "google.com",
                Elements: []el{
                    {
                        Tag: "div",
                        Attrs: []attr{
                            {Key: "class", Val: "red"}, 
                        },
                    },
                },
            },
            website{
                Url: "google.com",
            },
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func(t *testing.T) {
            if got := test.w.isEqualTo(&test.w2); got != test.out {
                t.Errorf("got=%t, expected=%t\n", got, test.out)
            }
        })
    }
}

func TestIsEqualToEl(t *testing.T) {
    cases := []struct {
        e el
        e2 el
        out bool
    }{
        {
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            true,
        },
        {
            el{
                Tag: "Div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "CLASS", Val: "reD"}, 
                },
            },
            true,
        },
        {
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "span",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            false,
        },
        {
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "id", Val: "red"}, 
                },
            },
            false,
        },
        {
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                    {Key: "id", Val: "red"}, 
                },
            },
            false,
        },
        {
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "green"}, 
                },
            },
            false,
        },
        {
            el{
                Tag: "div",
                Limit: 1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            el{
                Tag: "div",
                Limit: -1,
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func (t *testing.T) {
            if got := test.e.isEqualTo(&test.e2); got != test.out {
                t.Errorf("got=%t, expected=%t\n", got, test.out)
            }
        })
    }
}

func TestIsEqualToResult(t *testing.T) {
    cases := []struct {
        r result
        r2 result
        out bool
    }{
        {
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            true,
        },
        {
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            result{
                Url: "youtube.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            false,
        },
        {
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                    "anotherDiv": {"text1", "text2"},
                },
            },
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            false,
        },
        {
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text1", "text2"},
                },
            },
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"text2", "text2"},
                },
            },
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func (t *testing.T) {
            if got := test.r.isEqualTo(&test.r2); got != test.out {
                t.Errorf("got=%t, expected=%t\n", got, test.out)
            }
        })
    }
}

func TestGetAttrsEl(t *testing.T) {
    cases := []struct {
        el el
        out []string
    }{
        {
            el{
                Tag: "div",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                    {Key: "id", Val: "red"}, 
                },
            },
            []string{"class", "red", "id", "red"},
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            []string{"class", "red"},
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class"}, 
                },
            },
            []string{},
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                    {Key: "id"},
                },
            },
            []string{"class", "red"},
        },
        {
            el{
                Tag: "span",
            },
            []string{},
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func (t *testing.T) {
            got := test.el.getAttrs()
            if len(got) != len(test.out) {
                t.Fatalf("len(got)=%d, expected=%d\n", len(got), len(test.out))
            }

            if compareStr(got, test.out) == false {
                t.Errorf("got=%+v, expected=%+v\n", got, test.out)
            }
        })
    }
}

func TestGetKeysEl(t *testing.T) {
    cases := []struct {
        el el
        out string
    }{
        {
            el{
                Tag: "div",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                    {Key: "id", Val: "red"}, 
                },
            },
            "red red",
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                },
            },
            "red",
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class"}, 
                },
            },
            "",
        },
        {
            el{
                Tag: "span",
                Attrs: []attr{
                    {Key: "class", Val: "red"}, 
                    {Key: "id"},
                },
            },
            "red", 
        },
        {
            el{
                Tag: "span",
            },
            "",
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func (t *testing.T) {
            got := test.el.getKeys()

            if got != test.out {
                t.Errorf("got=%+v, expected=%+v\n", got, test.out)
            }
        })
    }
}

func TestCompareStr(t *testing.T) {
    cases := []struct{
        a []string
        b []string
        out bool
    }{
        {
            []string{"class", "red", "id", "red"},
            []string{"class", "red", "id", "red"},
            true,
        },
        {
            []string{"class", "red", "id", "red"},
            []string{"blue", "red", "id", "red"},
            false,
        },
        {
            []string{"red", "id", "red"},
            []string{"blue", "red", "id", "red"},
            false,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case #%d\n", i), func(t *testing.T) {
            got := compareStr(test.a, test.b)

            if got != test.out {
                t.Errorf("output=%t, expected=%t\n", got, test.out)
            }
        })
    }
}
