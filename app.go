package main

import (
    // "fmt"
    "golang.org/x/net/html"
    //"html/template"
    "log"
    "os"
    "strings"
    "bufio"
    "github.com/PuerkitoBio/goquery"
)

type ScrapeElement struct {
    path []string
    index int
}

var AdvancedDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".small-brand-logo"}, 0},
    "V2" : {[]string{".small-brand-logo"}, 1},
    "V3" : {[]string{".small-brand-logo"}, 2},
    "V4" : {[]string{".small-brand-logo"}, 3},
    "V5" : {[]string{".small-brand-logo"}, 4},
    "V6" : {[]string{".small-brand-logo"}, 5},
    "V7" : {[]string{".small-brand-logo"}, 6},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "V8" : {[]string{".official-shop-promo__banner-half"}, 0},
    "V9" : {[]string{".span6", ".row-fluid", ".mt-20"}, 0},
    "V10" : {[]string{".span6", ".row-fluid", ".span6.official-shop-promo__type2 "}, 1},
    "T": {[]string{".official-promo-top-video"}, 0},
}

var AdvancedMobile = map[string]ScrapeElement {
    "G" : {[]string{"h1"}, 0},
    "V1" : {[]string{".official-slide-category", ".swiper-slide"}, 0},
    "V2" : {[]string{".official-slide-category", ".swiper-slide"}, 1},
    "V3" : {[]string{".official-slide-category", ".swiper-slide"}, 2},
    "V4" : {[]string{".official-slide-category", ".swiper-slide"}, 3},
    "V5" : {[]string{".official-slide-category", ".swiper-slide"}, 4},
    "V6" : {[]string{".official-slide-category", ".swiper-slide"}, 5},
    "V7" : {[]string{".official-slide-category", ".swiper-slide"}, 6},
    "VV1" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 0},
    "VV2" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 1},
    "VV3" : {[]string{".official-shop-bigSlider", ".swiper-slide"}, 2},
    "V8" : {[]string{".span12"}, 0},
    "V9" : {[]string{".span6"}, 0},
    "V10" : {[]string{".span6"}, 1},
}

var AdvancedApps = AdvancedMobile

var FashionDesktop = map[string]ScrapeElement {
    "G" : {[]string{"h2"}, 0},
    "V1" : {[]string{".official-shop-promo__type1"}, 0},
    "V2" : {[]string{".official-shop-promo__type1"}, 1},
    "V3" : {[]string{".official-shop-promo__type1"}, 2},
    "V4" : {[]string{".official-shop-promo__type1"}, 3},
    "VV1" : {[]string{".swiper-slide"}, 0},
    "VV2" : {[]string{".swiper-slide"}, 1},
    "VV3" : {[]string{".swiper-slide"}, 2},
    "VV4" : {[]string{".swiper-slide"}, 3},
    "V5" : {[]string{".official-shop-promo__type2"}, 0},
    "V6" : {[]string{".official-shop-promo__type2"}, 1},
    "V7" : {[]string{".official-shop-promo__type2"}, 2},
    "V8" : {[]string{".official-shop-promo__type2"}, 3},
    "B1" : {[]string{".official-shop-branding__video"}, 0},
    "B2" : {[]string{".official-shop-branding__video"}, 1},
    "B3" : {[]string{".official-shop-branding__video"}, 2},
}

var Scrapers = map[string]interface{} {
    "advanced-Desktop" : AdvancedDesktop,
    "advanced-Mobile" : AdvancedMobile,
    "advanced-Apps" : AdvancedApps,

    "fashion-Desktop" : FashionDesktop,
}

func main() {
    // templateName := os.Args[1]
    // sectionName := os.Args[2]
    htmlFile := os.Args[1]
    outFile := os.Args[2]
    template := os.Args[3]
    section := os.Args[4]

    // log.Println(BuildTemplateMap(htmlFile))
    BuildDict(htmlFile, outFile, template, section)

}

func BuildTemplateMap(fp string) (string, error) {
    fph, err := os.Open(fp)
    if err != nil {
        return "adsf", err
    }

    defer fph.Close()
    doc, err := html.Parse(fph)
    if err != nil {
        return "adsf", err
    }
    //log.Println(doc)
    docs := goquery.NewDocumentFromNode(doc)


    // log.Println(docs.Find(".small-brand-logo").Find("a").Attr("href"))


    // list := make(map[string]*ScrapeElement)
    // list["G"] = &ScrapeElement{[]string{"h2"}, 0}
    // list["V1"] = &ScrapeElement{[]string{".small-brand-logo", "a"}, 4}
    list := AdvancedDesktop
    // var x = docs.Find("div")
    for key, val := range list {
        // x = docs.Find("div")
        log.Println(key + " = " +GetValue(docs, val, "img"))
    }
    // for _, v := range list["V1"].path {
    //     x = x.Find(v)
    //     // log.Println(x)
    // }
    // log.Println(x.Slice(list["V1"].index, x.Length()).Attr("data-name"))
    // temp := make(map[string]string)
    // temp["target"] = "asdf"
    // data["V8"] = temp

    // var f func(*html.Node)
 //    f = func(n *html.Node) {
 //        if n.Type == html.ElementNode {
 //            if (n.Data == "a" || n.Data == "img") {
 //                for _, x := range n.Attr {
 //                    if (x.Val != "") {
 //                        if (!strings.Contains(x.Key, "target")) {
 //                            //data[x.Key][x.Key] = x.Key
 //                            if (strings.Contains(x.Val, "{{")) {
 //                             log.Println(x.Val)
 //                            }
 //                        }
 //                    }
 //                }
 //            }

 //            // if the tag is a slider-div then we don't want to count the children
 //            // since html can have different number of slides and still should be counted
 //            // as valid template.
 //            if len(n.Attr) > 0 {
 //                if n.Attr[0].Key == "class" {
 //                    if n.Attr[0].Val == "swiper-wrapper" {
 //                        return
 //                    }
 //                }
 //            }
 //        }
 //        for c := n.FirstChild; c != nil; c = c.NextSibling {
 //            f(c)
 //        }
 //    }
 //    f(doc)

    // return data, nil
    return "adsf", nil
}

func GetValue(doc *goquery.Document, key ScrapeElement, attr string) string {
    x := doc.Find("div")
    for _, v := range key.path {
        x = x.Find(v)
    }
    // log.Println(key + " = "+x.Slice(val.index, x.Length()).AttrOr("href", "nein"))
    if (key.index < x.Length()) {
        if (attr == "img") {
            return x.Slice(key.index, x.Length()).Find("img").AttrOr("src", "")
        } else if (attr == "text") {
            return x.Slice(key.index, x.Length()).Text()
        } else if (attr == "video-url") {
            return x.Slice(key.index, x.Length()).Find("iframe").AttrOr("src", "")
        } else if (attr == "href" || attr == "data-value" || attr == "data-name" || strings.Contains(attr, "video")) {
            return x.Slice(key.index, x.Length()).Find("a").AttrOr(attr, "")
        } else if (attr == "txt") {
            return x.Slice(key.index, key.index+1).Find("p").Text()
        } else {
            return x.Slice(key.index, x.Length()).AttrOr(attr, "")
        }

        //log.Println(key + " = "+x.Slice(val.index, x.Length()).Find("img").AttrOr("src", "nein"))
    } else {
        //log.Println(key + fmt.Sprintf(" = none ; selection length: %d", x.Length()))
        return ""
    }
    return ""
}

func BuildDict(fp string, fouts string, template string, section string) {
    fph, err := os.Open(fp)
    if err != nil {
        return
    }

    defer fph.Close()

    fout, err := os.Create(fouts)
    if err != nil {
        return
    }
    defer fout.Close()

    doc, err := html.Parse(fph)
    if err != nil {
        return
    }
    //log.Println(doc)
    docs := goquery.NewDocumentFromNode(doc)

    list := Scrapers[template+"-"+section].(map[string]ScrapeElement)

    default_dict, err := os.Open(strings.ToLower(section)+".txt")
    if err != nil {
        return
    }
    scanner := bufio.NewScanner(default_dict)

    var sectionTitle string
    // var currentPart string
    // var partTitle string

    //Read the file line by line
    for scanner.Scan() {
        //Trim whitespaces before and after each line
        line := strings.TrimSpace(scanner.Text())
        if (line != "") {
            //Check if the line can be split by "="
            s := strings.Split(line, "=")
            key := strings.TrimSpace(s[0])
            // log.Print(key)
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
                    if (GetValue(docs, list[sectionTitle], "data-video-colorbox") == "") {
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
                // log.Println(sectionTitle)
            }
            // log.Println()
            // order = append(order, key)
        }
    }
}