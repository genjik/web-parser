package main

import (
    "testing"
    "strings"
    "errors"
    "fmt"
    "io"
)

func TestParseJson(t *testing.T) {
    cases := []struct {
        JSON io.Reader
        out []website
        err bool
    }{
        {
            strings.NewReader(
            `[
                {
                    "url": "www.youtube.com",
                    "elements": [
                        { 
                            "tag": "div",
                            "attrs": [
                                {"key": "class", "val": "review"}
                            ]
                        }
                    ]

                },
                {
                    "url": "www.google.com",
                    "elements": [
                        { 
                            "tag": "div",
                            "attrs": [
                                {"key": "class", "val": "review"}
                            ]
                        }
                    ]
                }
            ]`),
            []website{
                {
                    Url: "www.youtube.com",
                    Elements: []el{
                        {
                            Tag: "div",
                            Attrs: []attr{
                                {Key: "class", Val: "review"},
                            },
                        },
                    },
                },
                {
                    Url: "www.google.com",
                    Elements: []el{
                        {
                            Tag: "div",
                            Attrs: []attr{
                                {Key: "class", Val: "review"},
                            },
                        },
                    },
                },
            },
            false,
        },
        {
            strings.NewReader(
            `[
                {
                    "url": "www.amazon.com",
                    "elements": [
                        { 
                            "tag": "span",
                            "attrs": [
                                {"key": "class", "val": "review"},
                                {"key": "class", "val": "comments"}
                            ]
                        }
                    ]
                }
            ]`),
            []website{
                {
                    Url: "www.amazon.com",
                    Elements: []el{
                        {
                            Tag: "span",
                            Attrs: []attr{
                                {Key: "class", Val: "review"},
                                {Key: "class", Val: "comments"},
                            },
                        },
                    },
                },
            },
            false,
        },
        {
            strings.NewReader(
            `[
                {
                    "url": 
                    "elements": [
                        { 
                            "tag": "span",
                            "attrs": [
                                {"key": "class", "val": "review"},
                                {"key": "class", "val": "comments"}
                            ]
                        }
                    ]
                }
            ]`),
            []website{},
            true,
        },
        {
            strings.NewReader(
            `[
                {
                    "url": true,
                    "element": [
                        { 
                            "tag": true,
                            "attrs": [
                                {"key": "class", "val": "review"},
                                {"key": "class", "val": "comments"}
                            ]
                        }
                    ]
                }
            ]`),
            []website{},
            true,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func(t *testing.T) {
            got, err := ParseJson(test.JSON)
            if len(got) != len(test.out) {
                t.Fatalf("len(got)=%d, len(out)=%d\n", len(got), len(test.out))
            }
            if err != nil && test.err == false {
                t.Fatal(err)
            }

            for j, website := range test.out {
                if website.isEqualTo(&got[j]) == false{
                    t.Errorf("%d) got=%+v, expected=%+v", j, got[j], website)
                }
            }
        })
    }
}

// used to trigger error for html.Parse()
type dummyReader struct { Error error }
func (e *dummyReader) Read(n []byte) (int, error) {
    return 0, e.Error
}

func TestParseFW(t *testing.T) {
    cases := []struct {
        fw fetchedWebsite
        out result
        err bool
    }{
        {
            fetchedWebsite{
                website{
                    Url: "google.com",
                    Elements: []el{
                        {
                            Tag: "div",
                            Limit: -1,
                            Attrs: []attr{
                                {
                                    Key: "class",
                                    Val: "specialDiv",
                                },
                            },
                        },
                    },
                },
                strings.NewReader(`
                    <html>
                        <head></head>
                        <body>
                            <div id="red">Hello World!</div>
                            <div class="specialDiv">Number one</div>
                            <div class="specialDiv">Number 2</div>
                        </body>
                    </html>
                `),
            },
            result{
                Url: "google.com",
                Data: map[string][]string{
                    "specialDiv": {"Number one", "Number 2"},
                },
            },
            false,
        },
        {
            fetchedWebsite{
                website{
                    Url: "google.com",
                    Elements: []el{
                        {
                            Tag: "div",
                            Limit: -1,
                            Attrs: []attr{
                                {
                                    Key: "class",
                                    Val: "specialDiv",
                                },
                            },
                        },
                    },
                },
                strings.NewReader(`
                    <html>
                        <head></head>
                        <body>
                            <div id="red">Hello World!</div>
                            <div class="anotherDiv">Number one</div>
                            <div class="anotherDiv">Number 2</div>
                        </body>
                    </html>
                `),
            },
            result{
                Url: "google.com",
                Data: map[string][]string{
                },
            },
            false,
        },
        {
            fetchedWebsite{
                website{
                    Url: "google.com",
                    Elements: []el{
                        {
                            Tag: "div",
                            Limit: -1,
                            Attrs: []attr{
                                {
                                    Key: "class",
                                    Val: "specialDiv",
                                },
                            },
                        },
                    },
                },
                &dummyReader{errors.New("dummyError")},
            },
            result{},
            true,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func(t *testing.T) {
            got, err := test.fw.parse();
            if (err == nil && test.err == true) || (err != nil && test.err == false) {
                t.Fatalf("err != test.err")
            }

            if got.isEqualTo(&test.out) == false {
                t.Errorf("got=%+v, expected=%+v", got, test.out)
            }
        })
    }
}

func TestGenResult(t *testing.T) {
    cases := []struct {
        fw []fetchedWebsite
        out string
    }{
        {
            []fetchedWebsite{
                {
                    website{
                        Url: "google.com",
                        Elements: []el{
                            {
                                Tag: "div",
                                Limit: -1,
                                Attrs: []attr{
                                    {
                                        Key: "class",
                                        Val: "specialDiv",
                                    },
                                },
                            },
                        },
                    },
                    strings.NewReader(`
                        <html>
                            <head></head>
                            <body>
                                <div id="red">Hello World!</div>
                                <div class="specialDiv">Number one</div>
                                <div class="specialDiv">Number 2</div>
                            </body>
                        </html>
                    `),
                },
            },
`[
  {
    "Url": "google.com",
    "Data": {
      "specialDiv": [
        "Number one",
        "Number 2"
      ]
    }
  }
]`,
        },
        {
            []fetchedWebsite{
                {
                    website{
                        Url: "google.com",
                        Elements: []el{
                            {
                                Tag: "div",
                                Limit: -1,
                                Attrs: []attr{
                                    {
                                        Key: "class",
                                        Val: "specialDiv",
                                    },
                                },
                            },
                        },
                    },
                    strings.NewReader(`
                        <html>
                            <head></head>
                            <body>
                                <div id="red">Hello World!</div>
                                <div class="specialDiv">Number one</div>
                                <div class="specialDiv">Number 2</div>
                            </body>
                        </html>
                    `),
                },
                {
                    website{
                        Url: "youtube.com",
                        Elements: []el{
                            {
                                Tag: "div",
                                Limit: -1,
                                Attrs: []attr{
                                    {
                                        Key: "class",
                                        Val: "superSpecialDiv",
                                    },
                                },
                            },
                        },
                    },
                    strings.NewReader(`
                        <html>
                            <head></head>
                            <body>
                                <div id="red">Hello World!</div>
                                <div class="specialDiv">Number one</div>
                                <div class="superSpecialDiv">Number 2</div>
                            </body>
                        </html>
                    `),
                },
            },
`[
  {
    "Url": "google.com",
    "Data": {
      "specialDiv": [
        "Number one",
        "Number 2"
      ]
    }
  },
  {
    "Url": "youtube.com",
    "Data": {
      "superSpecialDiv": [
        "Number 2"
      ]
    }
  }
]`,
        },
    }

    for i, test := range cases {
        t.Run(fmt.Sprintf("Case#%d\n", i), func(t *testing.T) {
            ch := make(chan fetchedWebsite, 2)
            done := make(chan string)
            go genResult(ch, done)

            for _, v := range test.fw {
                ch <- v
            }
            close(ch)
            
            got := <-done
            if got != test.out {
                t.Errorf("\ngot     =%+v \nexpected=%+v", got, test.out)
            }
        })
    }
}
