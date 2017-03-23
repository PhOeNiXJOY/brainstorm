package main

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mirrr/govalidator"
	"golang.org/x/crypto/sha3"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"io/ioutil"
	"math/rand"
	// "net/http"
	"os"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	dmu     sync.Mutex
	dots    = map[string]string{}
	spliter = regexp.MustCompile("[\\/]+")
)

type (
	obj map[string]interface{}
	arr []interface{}
)

func init() {
	if os.Getenv("ENV") == "development" {
		fmt.Println("Init templates cleanup")
		go func() {
			for range time.Tick(time.Second) {
				func() {
					dmu.Lock()
					defer dmu.Unlock()
					dots = map[string]string{}
				}()
			}
		}()
	}
}

// плагин к шаблонизатору, подключающий файлы для doT.js без их парсинга
// template.FuncMap{"dot": dot}
func dot(name string) string {
	dmu.Lock()
	defer dmu.Unlock()
	if _, exists := dots[name]; !exists {
		dots[name] = "<!-- Template '" + name + "' not found! -->\n"
		if dat, err := ioutil.ReadFile("templates/doT.js/" + name); err == nil {
			s := strings.Split(name, ".")
			tplName := spliter.Split(s[0], -1)
			if s[len(s)-1] == "js" { // js темплейты
				dots[name] = "<!-- doT.js template - " + name + " -->\n" +
					"<script type='text/javascript' id='tpl_" + tplName[len(tplName)-1] + "'>\n" + string(dat) + "</script>\n"

			} else { // html темплейты
				dots[name] = "<!-- doT.js template - " + name + " -->\n" +
					"<script type='text/html' id='tpl_" + tplName[len(tplName)-1] + "'>\n" + string(dat) + "</script>\n"
			}
		}
	}
	return dots[name]
}

func postBind(c *gin.Context, ret interface{}) bool {
	c.BindWith(ret, binding.Form)
	if _, err := govalidator.ValidateStruct(ret); err != nil {
		ers := []string{}
		for k, v := range govalidator.ErrorsByField(err) {
			v = strings.Replace(v, "non zero value required", "не может быть пустым", -1)
			v = strings.Replace(v, "does not validate by 'min' tag", "меньше необходимого", -1)
			v = strings.Replace(v, "does not validate by 'max' tag", "больше необходимого", -1)
			v = strings.Replace(v, "does not validate as email", "- не электронный адрес", -1)
			ers = append(ers, k+": "+v)
		}
		c.String(400, strings.Join(ers, "<hr />"))
		return false
	}
	return true
}

func indexOf(arr interface{}, v interface{}) int {
	V := reflect.ValueOf(v)
	Arr := reflect.ValueOf(arr)
	if t := reflect.TypeOf(arr).Kind(); t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}
	for i := 0; i < Arr.Len(); i++ {
		if Arr.Index(i).Interface() == V.Interface() {
			return i
		}
	}
	return -1
}

func PassRand() string {

	rand.Seed(time.Now().UnixNano())

	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func H3hash(s string) string {
	h3 := sha3.New512()
	io.WriteString(h3, s)
	return fmt.Sprintf("%x", h3.Sum(nil))
}
