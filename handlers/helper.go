package handlers

import (
	"errors"
	"fmt"
	"goForum/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// 声明一个 *log.Logger 类型的 logger 变量作为日志处理器
var logger *log.Logger

func init() {
	// 日志文件存储到 logs/fofrum.log，通过 os.OpenFile() 打开这个日志句柄，如果文件不存在，则自动创建
	file, err := os.OpenFile("logs/goforum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to opeen log file", err)
	}
	// 通过 log.New() 初始化日志处理器并赋值给 logger。
	// 该方法需要传入日志文件、默认日志级别、以及日志格式
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 参数：...interface{} 表示可以传入的参数支持任意类型、任意个数。
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 异常处理统一重定向到错误页面
func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}