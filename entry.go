package main

import (
    "github.com/genjik/web-scraper"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "errors"
    "sync"
    //"time"
    "fmt"
    "io"
    "os"
)

func ParseJson(JSON io.Reader) ([]website, error) {
    var websites []website

    res, err := ioutil.ReadAll(JSON)
    if err != nil {
        return nil, err
    }
    if valid := json.Valid(res); valid == false {
        return nil, errors.New("Invalid json")
    }

    err = json.Unmarshal(res, &websites)

    if err != nil {
        return nil, err
    }

    return websites, nil
}

func startWorkerPool(websites []website, ch chan fetchedWebsite) {
    var wg sync.WaitGroup 

    for _, website := range websites {
        wg.Add(1)
        go fetch(website, ch, &wg)
    }
    wg.Wait()

    close(ch)
}

func fetch(website website, ch chan fetchedWebsite, wg *sync.WaitGroup) {
    defer wg.Done()

    res, err := http.Get(website.Url)

    if err != nil {
        fmt.Println(err)
        return 
    }

    ch <- fetchedWebsite{website, res.Body}
}

func genResult(ch <-chan fetchedWebsite, done chan<- string) {
    var results []result

    for fw := range ch {
        result, err := fw.parse()

        if err != nil {
            fmt.Println(err)
            continue
        }
        results = append(results, result)
    }

    output, err := json.MarshalIndent(results, "", "  ")

    if err != nil {
        fmt.Println(err)
        done <- ""
        return
    }

    done <- string(output)
}

func (fw *fetchedWebsite) parse() (result, error) {
    root, err := webscraper.GetRootElement(fw.Body)
    if err != nil {
        return result{}, err
    }

    data := make(map[string][]string)

    for _, el := range fw.Elements {
        found := root.FindAll(el.Tag, true, el.Limit, el.getAttrs()...)

        if len(found) < 1 {
            continue
        }

        var str []string

        for _, foundEl := range found {
            str = append(str, foundEl.GetText()) 
        }

        data[el.getKeys()] = str
    }        

    result := result{fw.website.Url, data}
    return result, nil
}

func startApp() error {
    stat, _ := os.Stdin.Stat()
    if (stat.Mode() & os.ModeCharDevice) != 0 {
        return errors.New("no std input")
    }

    websites, err := ParseJson(os.Stdin)
    if err != nil {
        return err
    }

    ch := make(chan fetchedWebsite) 
    done := make(chan string)

    go genResult(ch, done)
    startWorkerPool(websites, ch)

    fmt.Println(<- done)

    return nil
}

func main() {
    //startTime := time.Now()

    err := startApp()
    if err != nil {
        fmt.Println(err)
    }

    //endTime := time.Now()
    //fmt.Println("Loading time:", endTime.Sub(startTime).Seconds())
}
