package main

import (
    "fmt"
    "golang.org/x/net/html"
    //"html/template"
    "log"
    "os"
    "strings"
    "bufio"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    if len(os.Args) != 4 {
        fmt.Printf("Usage: ./htmlparser file.html outputfile template-name device-type\n")
        return
    }

    htmlFile := os.Args[1]
    outFile := os.Args[2]
    template := os.Args[3]
    section := os.Args[4]

    BuildDict(htmlFile, outFile, template, section)
}

func GetValue(doc *goquery.Document, key ScrapeElement, attr string) string {
    x := doc.Find("div")
    for _, v := range key.path {
        x = x.Find(v)
        //all-in-one template deviation
        if (v == ".brand-list-box") {
            if (key.index < x.Length()) {
                if (attr == "href" || attr == "data-value" || attr == "data-name") {
                    return x.Slice(key.index, x.Length()).AttrOr(attr, "")
                } else if (attr == "img") {
                    return x.Slice(key.index, x.Length()).Find("img").AttrOr("src", "")
                }
            }
        }
    }

    if (key.index < x.Length()) {
        if (attr == "img") {
            return x.Slice(key.index, x.Length()).Find("img").AttrOr("src", "")
        } else if (attr == "text") {
            return strings.TrimSpace(x.Slice(key.index, key.index+1).Text())
        } else if (attr == "video-url") {
            return x.Slice(key.index, x.Length()).Find("iframe").AttrOr("src", "")
        } else if (attr == "video-title") {
            return x.Slice(key.index, x.Length()).Find("iframe").AttrOr("data-video-title", "")
        } else if (attr == "href" || attr == "data-value" || attr == "data-name" || strings.Contains(attr, "video")) {
            return x.Slice(key.index, x.Length()).Find("a").AttrOr(attr, "")
        } else if (attr == "txt") {
            return x.Slice(key.index, key.index+1).Find("p").Text()
        } else {
            return x.Slice(key.index, x.Length()).AttrOr(attr, "")
        }

    } else {
        return ""
    }
    return ""
}

func BuildDict(fp string, fouts string, template string, section string) {
    //open html file
    fph, err := os.Open(fp)
    if err != nil {
        return
    }
    defer fph.Close()

    //create output file
    fout, err := os.Create(fouts)
    if err != nil {
        return
    }
    defer fout.Close()

    //parse html file
    doc, err := html.Parse(fph)
    if err != nil {
        return
    }
    docs := goquery.NewDocumentFromNode(doc)

    //get ScrapeElement
    list := Scrapers[template+"-"+section].(map[string]ScrapeElement)

    //get default dict content
    default_dict, err := os.Open(strings.ToLower(section)+".txt")
    if err != nil {
        return
    }
    scanner := bufio.NewScanner(default_dict)

    var sectionTitle string

    //Read the file line by line
    for scanner.Scan() {
        //Trim whitespaces before and after each line
        line := strings.TrimSpace(scanner.Text())
        if (line != "") {
            //Check if the line can be split by "="
            s := strings.Split(line, "=")
            key := strings.TrimSpace(s[0])
            if (len(s) > 1) {
                var content string
                if (key == "TitleText") {
                    content = GetValue(docs, list[sectionTitle], "text")
                } else if (key == "Image") {
                    if (GetValue(docs, list[sectionTitle], "data-video-colorbox") == "") {
                        content = GetValue(docs, list[sectionTitle], "img")
                    } else {
                        content = ""
                    }
                } else if (key == "Url") {
                    content = GetValue(docs, list[sectionTitle], "href")
                } else if (key == "DataValue") {
                    content = GetValue(docs, list[sectionTitle], "data-value")
                } else if (key == "DataName") {
                    if (GetValue(docs, list[sectionTitle], "data-video-colorbox") == "") {
                        content = GetValue(docs, list[sectionTitle], "data-name")
                    } else {
                        content = ""
                    }
                } else if (key == "VideoTitle") {
                    content = GetValue(docs, list[sectionTitle], "data-video-title")
                    if (content == "") {
                        content = GetValue(docs, list[sectionTitle], "video-title")
                    }
                } else if (key == "VideoUrl") {
                    content = GetValue(docs, list[sectionTitle], "data-video-colorbox")
                    if (content == "") {
                        content = GetValue(docs, list[sectionTitle], "video-url")
                    }
                } else if (key == "VideoId") {
                    content = GetValue(docs, list[sectionTitle], "data-video-id")
                } else if (key == "VideoName") {
                    content = GetValue(docs, list[sectionTitle], "data-name")
                } else if (key == "VideoImage") {
                    content = GetValue(docs, list[sectionTitle], "img")
                } else if (key == "Option") {
                    if (GetValue(docs, list[sectionTitle], "data-video-colorbox") == "" && GetValue(docs, list[sectionTitle], "video-url") == "") {
                        content = "image"
                    } else {
                        content = "video"
                    }
                } else if (key == "Target") {
                    url := GetValue(docs, list[sectionTitle], "href")
                    if (strings.Contains(url, "etalase") || url == "") {
                        content = "etalase"
                    } else if (strings.Contains(url, "promo")) {
                        content = "promo"
                    } else if (strings.Contains(url, "info")) {
                        content = "info"
                    } else {
                        content = "produk"
                    }
                } else if (key == "DisplayName") {
                    content = strings.TrimSpace(s[1]) + "\n"
                } else if (key == "Text") {
                    content = GetValue(docs, list[sectionTitle], "txt")
                } else if (key == "TabText") {
                    content = GetValue(docs, list[sectionTitle], "text")
                }
                log.Println("\t"+key + " = " + content)
                fout.WriteString("\t"+key + " = " + content + "\n")
            } else {
                log.Println(key)
                fout.WriteString(key + "\n")
                //Check if line is section title
                if (key[0] == '[' && key[len(key)-1] == ']') {
                    sectionTitle = key
                    sectionTitle = strings.TrimPrefix(sectionTitle, "[")
                    sectionTitle = strings.TrimSuffix(sectionTitle, "]")
                }
            }
        }
    }   
}