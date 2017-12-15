package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"

	"github.com/astaxie/beego"
	//"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type ProtoFile struct {
	module string
	file   string
}

type MethodInfo struct {
	MethodName     string
	CookieIsNeeded bool
	OpenApi        bool
}

func getProtolist(path string) ([]ProtoFile, error) {
	protos := []ProtoFile{}
	files, _ := ioutil.ReadDir(path)
	reg, _ := regexp.Compile("^([a-zA-Z0-9]+)Svr.proto$")
	for _, fi := range files {
		if !fi.IsDir() {
			m := reg.FindStringSubmatch(fi.Name())
			if len(m) < 2 || m[1] == "" {
				continue
			}
			module := m[1]
			protos = append(protos, ProtoFile{
				module: module,
				file:   path + "/" + fi.Name(),
			})
			//log.Print("Find module:", module)
		}
	}
	return protos, nil
}

func parseProto(file string) (*[]string, *[](MethodInfo), error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("Open file for stdout failed. err:", err)
		return nil, nil, err
	}
	defer f.Close()
	reg, _ := regexp.Compile("^\\s*message\\s+ST([a-zA-Z0-9]+)Req(\\s*//\\s*[Pp]ublic)?(\\s*\\|\\s*[Nn]o[Cc]ookie)?(\\s*\\|\\s*[O]pen)?")

	methods := ([]string{})
	publicMethods := ([](MethodInfo){})

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//log.Print("EOF")
				break
			}
			log.Fatal("err:", err)
			break
		}

		line = strings.TrimSpace(line)
		m := reg.FindStringSubmatch(line)
		if len(m) < 2 || m[1] == "" {
			continue
		}
		method := m[1]
		methods = append(methods, method)
		if len(m) >= 5 && len(m[4]) > 0 {
			publicMethods = append(publicMethods, MethodInfo{method, false, true})
		} else if len(m) >= 4 && len(m[3]) > 0 {
			publicMethods = append(publicMethods, MethodInfo{method, false, false})
		} else if len(m) >= 3 && len(m[2]) > 0 {
			publicMethods = append(publicMethods, MethodInfo{method, true, false})
		}

	}
	//log.Print("Parse proto done. module:", file, " methods:", methods,
	//	" publicMethod:", publicMethods)
	return &methods, &publicMethods, nil
}

func parseHandlerSrcFile(file string) (*[]string, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("Open file for stdout failed. err:", err)
		return nil, err
	}
	defer f.Close()
	//func HandleQueryRoomInfo(hdr *STUserTrustInfo, req *STQueryRoomInfoReq) (*STRspHeader, *STQueryRoomInfoRsp) {
	reg, _ := regexp.Compile("Handle.*STUserTrustInfo.*ST([a-zA-Z0-9]+)Req.*STRspHeader.*ST([a-zA-Z0-9]+)Rsp")

	methods := ([]string{})
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				//log.Print("EOF")
				break
			}
			log.Fatal("err:", err)
			break
		}

		line = strings.TrimSpace(line)
		m := reg.FindStringSubmatch(line)
		if len(m) < 2 || m[1] == "" {
			continue
		}
		method := m[1]
		methods = append(methods, method)
	}
	//log.Print("Parse proto done. module:", file, " methods:", methods)
	return &methods, nil
}

func hello(in string) (out string) {
	out = in + "world"
	return
}

func LowerCase(in string) string {
	return strings.ToLower(in)
}

func renderSrcFile(path, tpl string, data map[string]interface{}, appending bool) error {
	//log.Print("Start to render path:", path, " data:", data)

	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	if appending == true {
		mode = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	}

	fOut, err := os.OpenFile(path, mode, 0644)
	if err != nil {
		log.Fatal("Open file for stdout failed. err:", err)
		return err
	}
	defer fOut.Close()

	if beego.BeeTemplates[tpl] == nil {
		errmsg := fmt.Sprintf("BeeTemplates[%s] is nil.", tpl)
		log.Fatal(errmsg)
		return errors.New(errmsg)
	}

	if err := beego.BeeTemplates[tpl].ExecuteTemplate(fOut, tpl, &data); err != nil {
		log.Fatal(err)
	}
	return nil
}

func diff(a, b []string) *[]string {
	res := []string{}
	for _, ai := range a {
		found := false
		for _, bi := range b {
			if ai == bi {
				found = true
				break
			}
		}
		if found == false {
			res = append(res, ai)
		}
	}
	return &res
}

func main() {
	flag.Parse()
	args := flag.Args()
	protoDir := "../../../proto"
	srcDir := ".."
	if len(args) >= 1 {
		protoDir = args[0]
	}
	if len(args) >= 2 {
		srcDir = args[1]
	}
	templateDir := srcDir + "/code_template"
	clientDir := srcDir + "/client"
	controllerDir := srcDir + "/controllers"
	routerDir := srcDir + "/routers"

	log.Print("Start. protoDir:", protoDir, " templateDir:", templateDir, " clientDir:", clientDir,
		" controllerDir:", controllerDir, " routerDir:", routerDir)

	files := []string{
		"router.go.tpl",
		"client.go.tpl",
		"controller.go.tpl",
		"handler.go.tpl",
		"handler_item.go.tpl",
		"go_front_controller.go.tpl",
		"op_gateway_controller.go.tpl",
	}
	beego.AddFuncMap("lowerCase", LowerCase)
	for _, f := range files {
		if err := beego.BuildTemplate(templateDir, f); err != nil {
			log.Fatal(err)
		}
	}

	protoList, err := getProtolist(protoDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Get proto list:", protoList)
	methodMap := make(map[string](*[]string))
	publicMethodMap := make(map[string](*[](MethodInfo)))

	for _, p := range protoList {
		methods, publicMethods, err := parseProto(p.file)
		if err != nil {
			log.Fatal(err)
			continue
		}

		//handle business modules
		methodMap[p.module] = methods
		publicMethodMap[p.module] = publicMethods

		data := make(map[string]interface{})
		data["pkg"] = "client"
		data["module"] = p.module
		data["methods"] = *methods

		if err := renderSrcFile(clientDir+"/"+p.module+"Client.go",
			"client.go.tpl", data, false); err != nil {
			log.Fatal(err)
		}

		if err := renderSrcFile(controllerDir+"/"+p.module+"Controller.go",
			"controller.go.tpl", data, false); err != nil {
			log.Fatal(err)
		}
		handlerSrcFile := controllerDir + "/" + p.module + "Handler.go"
		if _, err := os.Stat(handlerSrcFile); os.IsNotExist(err) {
			// handlerSrcFile does not exist
			if err := renderSrcFile(handlerSrcFile, "handler.go.tpl", data, false); err != nil {
				log.Fatal(err)
			}
		} else {
			methodsExist, err := parseHandlerSrcFile(handlerSrcFile)
			if err != nil {
				log.Fatal(err)
			}
			newMethods := diff(*methods, *methodsExist)
			if len(*newMethods) > 0 {
				data["methods"] = *newMethods
				if err := renderSrcFile(handlerSrcFile, "handler_item.go.tpl", data, true); err != nil {
					log.Fatal(err)
				}
			}
		}

	}

	// generate router.go
	data := make(map[string]interface{})
	data["pkg"] = "client"
	data["modules"] = methodMap
	data["publics"] = publicMethodMap
	if err := renderSrcFile(routerDir+"/router.go", "router.go.tpl", data, false); err != nil {
		log.Fatal(err)
	}

	{
		data := make(map[string]interface{})
		data["pkg"] = "client"
		data["modules"] = methodMap
		if err := renderSrcFile(controllerDir+"/OpGatewayController.go",
			"op_gateway_controller.go.tpl", data, false); err != nil {
			log.Fatal(err)
		}
	}

	{
		data := make(map[string]interface{})
		data["pkg"] = "client"
		data["publics"] = publicMethodMap
		if err := renderSrcFile(controllerDir+"/GoFrontController.go",
			"go_front_controller.go.tpl", data, false); err != nil {
			log.Fatal(err)
		}
	}

}
