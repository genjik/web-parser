# Web-parser

web-parser is a cli tool that concurrently parses and retrieves information from multiple websites

## How it works?
1) You pass .json file that contains urls and html elements to standard input.

2) The parser launches new goroutine for each url, fetches the html code, looks for html elements and finally generates result.

3) The result is returned to standard output in .json format

## Example
**Input json:**  
```json
[
    {
        "url": "http://example.com",
        "elements": [
            {
                "tag": "div",
                "limit": -1, 
                "attrs": [
                    {"key": "class", "val": "red-div"}
                ]
            }
        ]
    },
    {
        "url": "http://example2.com",
        "elements": [
            {
                "tag": "a",
                "limit": 1, 
                "attrs": [
                    {"key": "class", "val": "green-div"},
                    {"key": "href", "val": "#nowhere"}
                ]
            },
            {
                "tag": "div",
                "limit": 2, 
                "attrs": [
                    {"key": "class", "val": "red-div"},
                ]
            }
        ]
    }
]
```

**Output json:**  
```json
[
    {
        "url": "http://example.com",
        "data": [
            "red-div": [
                "text...",
                "text...",
                "text..."
            ]
        ]
    },
    {
        "url": "http://example2.com",
        "data": [
            "green-div #nowhere": [
                "only one result"
            ],
            "red-div": [
                "one",
                "two"
            ]
        ]
    }
]
```

## The main limitation
The parser is not flexible yet.
As of now it can't parse the following html
```html
    <div class="item">
        <div class="name">name0</div>
        <div class="price">0</div>
    </div>
    <div class="item">
        <div class="name">name1</div>
        <div class="price">1</div>
    </div>
    <div class="item">
        <div class="name">name2</div>
        <div class="price">2</div>
    </div>
```
as
```json
    {
        "url": "...",
        "data": [
            {
                "name": "name0",
                "price": 0
            },
            {
                "name": "name1",
                "price": 1
            },
            {
                "name": "name2",
                "price": 2
            }
        ]
    }
```
