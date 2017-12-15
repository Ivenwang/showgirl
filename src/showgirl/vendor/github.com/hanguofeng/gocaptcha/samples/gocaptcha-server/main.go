// CServer project main.go
package main

import (
	"flag"
	"fmt"
	"html"
	// "image/png"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"bytes"
	"encoding/base64"

	"github.com/hanguofeng/config"
	"github.com/hanguofeng/gocaptcha"
)

var (
	ccaptcha   *gocaptcha.Captcha
	configFile = flag.String("c", "gocaptcha.conf", "the config file")
)

const (
	DEFAULT_PORT = "80"
	DEFAULT_LOG  = "logs/gocaptcha-server.log"
)

func ShowImageHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	if len(key) >= 0 {
		cimg, err := ccaptcha.GetImage(key)
		log.Println("err", err)
		if nil == err {
			// w.Header().Add("Content-Type", "image/png")
			// png.Encode(w, cimg)
			w.Header().Add("Content-Type", "image/jpeg")
			jpeg.Encode(w, cimg, &jpeg.Options{100})

			buf := new(bytes.Buffer)
			err = jpeg.Encode(buf, cimg, &jpeg.Options{100})
			// send_s3 := buf.Bytes()
			str := base64.StdEncoding.EncodeToString(buf.Bytes())
			encodeStr := fmt.Sprintf("data:image/jpeg;base64,%s", str)
			log.Print(encodeStr)
		} else {
			log.Printf("show image error:%s", err.Error())
			w.WriteHeader(500)
		}
	}

	log.Printf("[cmd:showimage][remote_addr:%s][key:%s]", r.RemoteAddr, key)
}

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	callback := html.EscapeString(r.FormValue("callback"))

	key, err := ccaptcha.GetKey(4)
	retstr := "{error_no:%d,error_msg:'%s',key:'%s'}"

	error_no := 0
	error_msg := ""

	if nil != err {
		error_no = 1
		error_msg = err.Error()
	}

	if callback != "" {
		retstr = "%s(" + retstr + ")"
		retstr = fmt.Sprintf(retstr, callback, error_no, error_msg, key)
	} else {
		retstr = fmt.Sprintf(retstr, error_no, error_msg, key)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(retstr))

	log.Printf("[cmd:getkey][remote_addr:%s][key:%s]", r.RemoteAddr, key)
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	code := r.FormValue("code")
	callback := html.EscapeString(r.FormValue("callback"))

	retstr := "{error_no:%d,error_msg:'%s',key:'%s'}"
	error_no := 0
	error_msg := ""

	suc, msg := ccaptcha.Verify(key, code)

	if false == suc {
		error_no = 1
		error_msg = msg
	}

	if callback != "" {
		retstr = "%s(" + retstr + ")"
		retstr = fmt.Sprintf(retstr, callback, error_no, error_msg, key)
	} else {
		retstr = fmt.Sprintf(retstr, error_no, error_msg, key)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(retstr))
	log.Printf("[cmd:verify][remote_addr:%s][key:%s][code:%s]", r.RemoteAddr, key, code)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	retstr := "<html>"
	retstr += "<body>"
	retstr += "<h1>gocaptcha server</h1>"
	retstr += "<h2>document</h2>"
	retstr += "<p>see:<a href='https://github.com/hanguofeng/gocaptcha/tree/master/samples/gocaptcha-server'>https://github.com/hanguofeng/gocaptcha/tree/master/samples/gocaptcha-server</a></p>"
	retstr += "<h2>interface</h2>"
	retstr += "<p><a href='/getkey'>/getkey</a></p>"
	retstr += "<p><a href='/showimage'>/showimage</a></p>"
	retstr += "<p><a href='/verify'>/verify</a></p>"
	retstr += "</body>"
	retstr += "</html>"
	w.Write([]byte(retstr))
}

func main() {

	flag.Parse()

	/* 1.load the config file and assign port/logfile */
	port := DEFAULT_PORT
	logfile := DEFAULT_LOG

	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		log.Fatalf("config file:%s not exists!", *configFile)
		os.Exit(1)
	}

	c, err := config.ReadDefault(*configFile)
	if nil != err {
		port = DEFAULT_PORT
		logfile = DEFAULT_LOG
	}
	port, err = c.String("service", "port")
	if nil != err {
		port = DEFAULT_PORT
	}
	logfile, err = c.String("service", "logfile")
	if nil != err {
		logfile = DEFAULT_LOG
	}
	log.Print("debug 1")

	os.MkdirAll(filepath.Dir(logfile), 0777)
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE, 0666)
	log.SetOutput(f)

	log.Print("debug 2")
	captcha, err := gocaptcha.CreateCaptchaFromConfigFile(*configFile)

	log.Print("debug 3")

	if nil != err {
		log.Fatalf("config load failed:%s", err.Error())
	} else {
		ccaptcha = captcha
	}

	/* 2. bind handler */
	http.HandleFunc("/showimage", ShowImageHandler)
	http.HandleFunc("/getkey", GetKeyHandler)
	http.HandleFunc("/verify", VerifyHandler)
	http.HandleFunc("/", IndexHandler)

	/* 3. run the http server */
	s := &http.Server{Addr: ":" + port}

	log.Printf("=======ready to serve=======")
	log.Fatal(s.ListenAndServe())
	f.Close()
}
