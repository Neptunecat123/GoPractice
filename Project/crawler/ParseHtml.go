package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var FullPath = "./ftcy/"
var testTopic = "229987981"
var testDate = "2021-11-06"

/*
遍历整个ftcy下目录内的0.html中获取：
topicname topicurl authorurl authorname likenum collectnum commentnum createtime latesttime iselite
todo：如何使用goroutine解析html
*/

// 获取目录下的文件或子目录全路径，返回[]string
func GetSubDir(path string) []string {
	var TopicDirList []string
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, dir := range rd {
		fullPath := filepath.Join(path, dir.Name())
		TopicDirList = append(TopicDirList, fullPath)
	}

	return TopicDirList
}

// 读取分析单个topic下的topic mysql数据，返回topicMysql结构体
func GetTopicMysqlData(htmlpath string) TopicMysql {
	var topicData TopicMysql
	firstHtml := "0.html"
	htmlfile := filepath.Join(htmlpath, firstHtml)
	r, err := os.Open(htmlfile)
	if err != nil {
		fmt.Println(err.Error())
		return TopicMysql{}
	}
	dom, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println(err.Error())
		return TopicMysql{}
	}
	defer r.Close()
	dom.Find("h1").Each(func(i int, selection *goquery.Selection) {
		topicData.TopicName = selection.Text()
	})
	dom.Find(".from").Find("a").Each((func(i int, selection *goquery.Selection) {
		topicData.AuthorName = selection.Text()
		topicData.AuthorUrl, _ = selection.Attr("href")
	}))
	dom.Find("h3").Find(".create-time").Each((func(i int, selection *goquery.Selection) {
		topicData.CreateTime = selection.Text()
	}))
	dom.Find(".action-collect").Find("a").Each((func(i int, selection *goquery.Selection) {
		topicData.TopicUrl, _ = selection.Attr("data-url")
	}))
	dom.Find(".action-collect>.collect-add>.react-num").Each((func(i int, selection *goquery.Selection) {
		if len(selection.Text()) != 0 {
			topicData.CollectNum, _ = strconv.Atoi(selection.Text())
		} else {
			topicData.CollectNum = 0
		}

	}))
	dom.Find(".action-react>.react-add>.react-num").Each((func(i int, selection *goquery.Selection) {
		if len(selection.Text()) != 0 {
			topicData.LikeNum, _ = strconv.Atoi(selection.Text())
		} else {
			topicData.LikeNum = 0
		}

	}))
	dom.Find("script").Each((func(i int, selection *goquery.Selection) {
		if s, _ := selection.Attr("type"); s == "application/ld+json" {
			b := []byte(selection.Text())
			for i, ch := range b {
				switch {
				case ch > '~':
					b[i] = ' '
				case ch == '\r':
				case ch == '\n':
					b[i] = ' '
				case ch == '\t':
				case ch < ' ':
					b[i] = ' '
				}
			}
			var m map[string]interface{}
			err := json.Unmarshal(b, &m)
			if err != nil {
				fmt.Println("unmarshal: ", err.Error())
			}
			topicData.CommentNum, err = strconv.Atoi(m["commentCount"].(string))
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}))

	htmlList := GetSubDir(htmlpath)
	latestHtml := htmlList[len(htmlList)-1]
	l, err := os.Open(latestHtml)
	if err != nil {
		fmt.Println(err.Error())
		return TopicMysql{}
	}
	dom, err = goquery.NewDocumentFromReader(l)
	if err != nil {
		fmt.Println(err.Error())
		return topicData
	}
	defer l.Close()
	topicData.LatestTime = dom.Find(".pubtime").Last().Text()

	return topicData
}

// // 读取分析单个comment mysql数据，返回[]commentMysql结构体列表
// func GetCommentMysqlData(htmlpath string) []CommentMysql {
// 	var commentMysqlList = make([]CommentMysql, 5)
// 	var commentMysql CommentMysql
// }

// 获取带解析文件的目录列表，每个解析都是不互相影响的，可以遍历也可以goroutine协程处理
// func GetAllTopicPath(date string) ([]string, error) {

// }

// func GetTopicInfo(path string) (TopicMysql, error) {
// 	// goquery.NewDocumentFromReader()
// }

// func GetCommentInfo(path string) (CommentMysql, error) {

// }

// func InsertMysql()

// func Parser(path string) error {
// 	var TopicData TopicMysql
// 	var CommentData CommentMysql
// 	DatePath := FullPath + testDate
// 	AllTopicPath, err := GetAllTopicPath(DatePath)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	// todo: 这个流程优化成goroutine
// 	for _, tp := range AllTopicPath {
// 		TopicData, err = GetTopicInfo(tp)
// 		if err != nil {
// 			fmt.Printf("Get Topic info failed: %s", tp)
// 		}
// 		InsertTopicData(&TopicData)

// 		CommentData, err = GetCommentInfo(tp)
// 		if err != nil {
// 			fmt.Printf("Get Comment info failed: %s", tp)
// 		}
// 		InsertCommentData(&CommentData)
// 	}

// }
