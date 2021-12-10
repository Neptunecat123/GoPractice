package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	basePath = "C:/Users/BJMX/go/src/crawler/ftcy"
	cookie   = `bid=evc-GKaAiJY; __yadk_uid=RhNUOPjqNOx6RiIBxQVZCRxz1T5MdVhU; __utmz=30149280.1634563451.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __gads=ID=bff3bacc9173990f-2203ed8daecc0085:T=1634563451:RT=1634563451:S=ALNI_MbGEVi3R8bEks-VVuop6JP6P6AJYg; douban-fav-remind=1; ll="118371"; push_doumail_num=0; __utmv=30149280.5500; dbcl2="55006472:J8QqPTwM0kw"; ct=y; ck=57pg; ap_v=0,6.0; __utmc=30149280; _pk_ses.100001.8cb4=*; push_noty_num=0; __utma=30149280.1927394790.1634563451.1636096999.1636101172.36; __utmt=1; _pk_id.100001.8cb4=cd772b288c4a1037.1634563451.36.1636101183.1636098358.; __utmb=30149280.17.5.1636101183989`
	cookie2  = `bid=_t2sTcVOtv8; _pk_ses.100001.8cb4=*; ap_v=0,6.0; __yadk_uid=3tKoE2hUcz707xJFGjULVEP1HFrhaBUe; __gads=ID=c774cf99a0b8e175-22f9512897ce0031:T=1636182388:RT=1636182388:S=ALNI_MboNRfiqZcYndIhutSPJH4IMjxpLA; dbcl2="55006472:J8QqPTwM0kw"; ck=57pg; _pk_id.100001.8cb4=637f8b727e4cda16.1636182387.1.1636182430.1636182387.; push_noty_num=0; push_doumail_num=0; __utma=30149280.281996767.1636182432.1636182432.1636182432.1; __utmc=30149280; __utmz=30149280.1636182432.1.1.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmt=1; __utmv=30149280.5500; __utmb=30149280.5.7.1636182432`
)

// func GetTopicsByLastestReply(btime, etime string) ([]string, error) {
// 	// 起始时间，终止时间为空串时，取所有topic；
// 	// 起始时间为空串，取最早到指定终止时间的帖子
// 	// 终止时间为空串，取指定起始时间到最新的所有帖子。

// }

func RandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func MakeDir(basePath, dirName string) (string, error) {
	dirPath := path.Join(basePath, dirName)
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		os.RemoveAll(dirPath)
	}
	err := os.MkdirAll(dirPath, 0766)
	if err != nil {
		return "", err
	}
	return dirPath, nil

}

// 优化：封装http.Client, 把get，封成一个函数，return rep，这样不用每次都要写req.Header.Add
func HttpClient(url string) (*http.Response, error) {
	fmt.Println("Getting url: ", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.54 Safari/537.36")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("sec-ch-ua", `"Google Chrome";v="95", "Chromium";v="95", ";Not A Brand";v="99"`)
	req.Header.Add("sec-ch-ua-platform", "Windows")

	rep, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return rep, nil

}

// 获取所有topic的url list
func GetAllTopic(url string) ([]string, error) {
	rep, err := HttpClient(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(rep.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rep.Body.Close()

	num := doc.Find(".olt").Find(".title>a").Length()
	onePageUrls := make([]string, num)
	doc.Find(".olt").Find(".title>a").Each(func(i int, selection *goquery.Selection) {

		// fmt.Println(selection.Text())
		href, ok := selection.Attr("href")

		if ok {
			onePageUrls[i] = href
			// fmt.Println(href)
		} else {
			onePageUrls[i] = ""
		}
	})

	return onePageUrls, nil
}

// 构造每页topic url列表
func PageUrlList(url string) ([]string, error) {
	// 读取第一页，获取总页数，通过总页数构造url列表
	rep, err := HttpClient(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(rep.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rep.Body.Close()
	totalPage, ok := doc.Find(".thispage").Attr("data-total-page")
	if !ok {
		fmt.Println("no data-total-page")
		return nil, nil
	}

	pageNum, err := strconv.Atoi(totalPage)
	if err != nil {
		return nil, nil
	}
	urls := make([]string, pageNum)
	urlBase := "https://www.douban.com/group/ftcy/discussion?start="
	for i := 0; i < pageNum; i++ {
		urls[i] = urlBase + strconv.Itoa(i*25)
	}

	return urls, nil

}

// 构造每页comment url列表
func CommentUrlList(baseUrl, totalPage string) ([]string, error) {
	totalNum, err := strconv.Atoi(totalPage)
	if err != nil {
		return nil, err
	}
	urls := make([]string, totalNum)
	for i := 0; i < totalNum; i++ {
		urls[i] = baseUrl + "?start=" + strconv.Itoa(i*100)
	}

	return urls, nil

}

// 下载保存html文件到指定目录
func DownLoadHtml(index int, url, downloadPath string) error {
	rep, err := HttpClient(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer rep.Body.Close()

	body, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		fmt.Println("Read Error:", err)
		return err
	}

	filePath := path.Join(downloadPath, strconv.Itoa(index)+".html")
	fmt.Println(filePath)
	htmlFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer htmlFile.Close()
	n, err := htmlFile.Write(body)
	if err == nil && n < len(body) {
		return io.ErrShortWrite
	}
	if err != nil {
		return err
	}

	return nil
}

func DownLoadOneTopic(basePath, url string) error {
	// 根据url截取topic编号，mkdir目录
	topicNum := strings.Split(url, "/")[5]
	topicDir, err := MakeDir(basePath, topicNum)
	if err != nil {
		return err
	}
	// 根据url查看comment页数
	rep, err := HttpClient(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	doc, err := goquery.NewDocumentFromReader(rep.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer rep.Body.Close()
	var urls []string
	totalPage, ok := doc.Find(".thispage").Attr("data-total-page")
	if !ok {
		totalPage = "1"
		urls = []string{url}
	} else {
		urls, err = CommentUrlList(url, totalPage)
		if err != nil {
			return err
		}
	}
	fmt.Println(urls)
	for i, commentUrl := range urls {
		if (i+1)%10 == 0 {
			time.Sleep(time.Duration(5) * time.Second)
		}
		err := DownLoadHtml(i, commentUrl, topicDir)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	fmt.Println("Finish topic:", url)
	return nil

}

// 根据每页comment url列表，下载每个topic的所有页面
func DownLoadTopicsHtml(topicDir string, urls []string) error {
	// 遍历执行 DownLoadOneTopic， todo：优化为goroutine，但要控制速度，不要被反爬ban掉
	for _, url := range urls {
		fmt.Println("Getting DownLoadTopicsHtml: ", url)
		err := DownLoadOneTopic(topicDir, url)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	return nil
}

func main() {

	firstDate := "2021-11-06"
	dirPath, err := MakeDir(basePath, firstDate)
	if err != nil {
		fmt.Println(err.Error())
	}

	allTopicsUrls := make([]string, 0)
	for i, url := range allTopicsUrls {
		fmt.Println("Get Page ", i)
		if (i+1)%3 == 0 {
			fmt.Println(allTopicsUrls)
			DownLoadTopicsHtml(dirPath, allTopicsUrls)
			sleepTime := RandInt(10, 20)
			fmt.Printf("Sleep %d sec...", sleepTime)
			time.Sleep(time.Duration(sleepTime) * time.Second)
			allTopicsUrls = make([]string, 0)
		}
		onePageUrls, err := GetAllTopic(url)
		if err != nil {
			fmt.Println(err.Error())
		}
		allTopicsUrls = append(allTopicsUrls, onePageUrls...)
		time.Sleep(time.Duration(2) * time.Second)
	}

	// gouroutine处理生成topicMysql列表
	htmlPath := "C:\\Users\\BJMX\\go\\src\\crawler\\ftcy\\2021-11-06"

	topicDatas := GetTopicsMysql(htmlPath)
	fmt.Println("len topic:", len(topicDatas))
	for _, topic := range topicDatas {
		// fmt.Println(topic)
		InsertTopicData(&topic)
	}

}

func GetTopicsMysql(htmlPath string) []TopicMysql {

	workerNum := runtime.NumCPU() - 1

	topicListChan := make(chan string, 3000)
	topicResultChan := make(chan TopicMysql, 3000)
	doneChan := make(chan bool, workerNum)

	// 生成topic channel
	go topicGenerator(topicListChan, workerNum, htmlPath)

	// 从topic channel读取每个topic html，解析生成topicMysql结构体写入topicsMysql channel
	for i := 0; i < workerNum; i++ {
		go topicPraserChan(topicListChan, topicResultChan, doneChan)
	}

	// 判断是否完成topic解析，完成则关闭topicMysql channel
	go controlTopicResultChan(doneChan, topicResultChan, workerNum)

	// 收集topicMysql结果
	topicResult := getAllTopicStruct(topicResultChan)
	return topicResult

}

func topicGenerator(topicList chan string, workNum int, htmlPath string) {
	rd, err := ioutil.ReadDir(htmlPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, dir := range rd {
		fullPath := filepath.Join(htmlPath, dir.Name())
		topicList <- fullPath
	}
	close(topicList)
}

func topicPraserChan(topicList chan string, topicResult chan TopicMysql, doneCh chan bool) {
	for topic := range topicList {
		topicResult <- GetTopicMysqlData(topic)
	}
	doneCh <- true

}

func controlTopicResultChan(doneChan chan bool, topicResult chan TopicMysql, workNum int) {
	finshedCount := 0
	for range doneChan {
		finshedCount++
		if finshedCount == workNum {
			break
		}
	}
	close(topicResult)
}

func getAllTopicStruct(topicResult chan TopicMysql) []TopicMysql {

	topicDatas := make([]TopicMysql, 0)
	for topicData := range topicResult {
		if topicData.AuthorUrl != "" {
			topicDatas = append(topicDatas, topicData)
		} else {
			fmt.Println("no data: ", topicData)
		}
	}
	return topicDatas
}
