package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"strconv"

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
}

func getProtolist(path string) ([]ProtoFile, error) {
	protos := []ProtoFile{}
	protos = append(protos, ProtoFile{
		module: "TianTianLive",
		file:   path + "/TianTianLive.proto",
	})
	//	protos = append(protos, ProtoFile{
	//		module: "OpenAPI",
	//		file:   path + "/OpenAPI.proto",
	//	})
	protos = append(protos, ProtoFile{
		module: "CommonDef",
		file:   path + "/CommonDef.proto",
	})

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
			log.Print("Find module:", module)
		}
	}
	return protos, nil
}

const (
	OUT        = 0
	IN_ENUM    = 1
	IN_MESSAGE = 2
)

func genCode(code string) string {
	return fmt.Sprintf("<li><code><span>%s</span></code></li>", code)
}

type STPropertyInfo struct {
	Name     string
	Type     string
	Required string
}
type TypeInfoMap map[string]([]STPropertyInfo)

var (
	reg, _      = regexp.Compile("^\\s*((message)|(enum))\\s+([a-zA-Z0-9_]+)\\s*(//(.*))?")
	regField, _ = regexp.Compile("^\\s*((required)|(optional)|(repeated))\\s+([a-zA-Z0-9_]+)\\s+([a-zA-Z0-9_]+)\\s*=\\s*\\d+\\s*;\\s*(//(.*))?")
)

//@brief 解析proto文件得到message对应的类型信息并写入TypeInfoMap
//@param string proto文件名
//@param typeInfoMap 类型信息Map
//@return TypeInfoMap 更新后的类型信息Map
//@return error
func parseTypeInfo(file string, typeInfoMap TypeInfoMap) (TypeInfoMap, error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("Open file readonly failed. err:", err, " ,file:", file)
		return nil, err
	}
	defer f.Close()

	curTypeName := ""
	where := OUT
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Print("EOF")
				break
			}
			log.Fatal("err:", err)
			break
		}
		line = strings.TrimSpace(line)
		m := reg.FindStringSubmatch(line)
		//log.Print("len(m): ", len(m), ", m:", m)

		if len(m) >= 5 && len(m[4]) > 0 { //message or enum
			if len(m[2]) > 0 { //message
				where = IN_MESSAGE
				curTypeName = m[4]
				typeInfoMap[curTypeName] = []STPropertyInfo{}
			} else { // enum
				where = IN_ENUM
			}
		} else if where == IN_MESSAGE {
			mFields := regField.FindStringSubmatch(line)
			if len(mFields) > 0 { // message field description
				requiredOrOptionalOrRepeated := mFields[1]
				fieldTypeName := mFields[5]
				fieldName := mFields[6]
				typeInfoMap[curTypeName] = append(typeInfoMap[curTypeName], STPropertyInfo{
					Name:     fieldName,
					Type:     fieldTypeName,
					Required: requiredOrOptionalOrRepeated,
				})
			}
		} else {
			if where == IN_ENUM {
			} else { // OUT
			}
		}
	}

	for k, slc := range typeInfoMap {
		log.Printf("Type: %s", k)
		for _, property := range slc {
			log.Printf("  %s %s", property.Name, property.Type)
		}
	}
	return typeInfoMap, nil
}

func genSampleReq(typeInfoMap TypeInfoMap, method string, isPublic bool, isOpenApi bool) string {
	code := "请求示例：<br/>"
	if isPublic {
		code += " HttpHeader:<br/>"
		code += "  Statistics: {\"cv\":\"1.6.3\",\"osv\":\"9.3.3\",\"os\":1,\"net\":1,\"uuid\":\"sdfa-jgkadd-fadaf\",\"phone\":\"xiaomi2\"}<br/>"
		code += "  UserKey: vMrO9nV5mnM9AydMA0Vu4dvZnRY=<br/>"
	}
	code += " HttpPostBody: "

	code += genPropertySample(typeInfoMap, "ST"+method+"Req", 1)
	return code
}

func genSampleRsp(typeInfoMap TypeInfoMap, method string, isPublic bool, isOpenApi bool) string {
	code := "应答示例：<br/>{<br/>"
	code += fmt.Sprintf("&nbsp;\"RspHeader\" : %s", genPropertySample(typeInfoMap, "STRspHeader", 1))
	code += fmt.Sprintf("&nbsp;\"RspJson\" : %s", genPropertySample(typeInfoMap, "ST"+method+"Rsp", 1))
	code += "}<br/>"
	return code
}

func genPropertySample(typeInfoMap TypeInfoMap, typeName string, depth int) string {

	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "&nbsp;&nbsp;"
	}

	if sli, found := typeInfoMap[typeName]; found {
		code := "{<br/>"
		for _, v := range sli {
			repeatedFlag := ""
			if v.Required == "repeated" {
				repeatedFlag = "[] "
			}
			code += fmt.Sprintf(prefix+"&nbsp;&nbsp;\"%s\" : %s%s<br/>", v.Name, repeatedFlag, genPropertySample(typeInfoMap, v.Type, depth+1))
		}
		code += prefix + "}<br/>"
		return code
	}

	return typeName

}

//@brief 解析proto文件并生成对应的Md文档
//@param file proto文件名
//@param fout 文件指针
//@param onlyPublic 是否只包含Public方法
//@param TypeInfoMap 只读类型信息Map
//@return error
func renderMarkdown(file string, fout *os.File, moduleName string, onlyPublic bool, typeInfoMap TypeInfoMap) error {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("Open file readonly failed. err:", err, " ,file:", file)
		return err
	}
	defer f.Close()

	regForMethod, _ := regexp.Compile("^\\s*message\\s+ST([a-zA-Z0-9]+)Req(\\s*//\\s*[Pp]ublic)?(\\s*\\|\\s*[Nn]o[Cc]ookie)?(\\s*\\|\\s*[Oo]pen)?")
	regExplain, _ := regexp.Compile("^\\s*//(.*)")
	regExplainPath, _ := regexp.Compile("^\\s*//\\s*([Pp]ath)\\s*:\\s*(.*)")
	regExplainBrief, _ := regexp.Compile("^\\s*//\\s*([Bb]rief)\\s*:\\s*(.*)")
	regEndEnum, _ := regexp.Compile("^\\s*}")

	typeNames := []string{}
	methods := ([]string{})
	publicMethods := ([](MethodInfo){})

	where := OUT
	explainStr := ""
	methodPath := ""
	brief := ""
	buf := bufio.NewReader(f)
	fout.WriteString(fmt.Sprintf("# %s\n\n", moduleName))
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Print("EOF")
				break
			}
			log.Fatal("err:", err)
			break
		}
		line = strings.TrimSpace(line)
		m := reg.FindStringSubmatch(line)
		//log.Print("len(m): ", len(m), ", m:", m)

		if len(m) >= 5 && len(m[4]) > 0 { //message or enum
			typeName := m[4]
			typeNames = append(typeNames, typeName)

			mMethods := regForMethod.FindStringSubmatch(line)
			if len(mMethods) >= 2 && len(mMethods[1]) > 0 { //method
				method := mMethods[1]
				needCookie := true
				isPublic := false
				isOpenApi := false
				if len(mMethods) >= 5 && len(mMethods[4]) > 0 {
					isOpenApi = true
					needCookie = false
					isPublic = true
				} else if len(mMethods) >= 4 && len(mMethods[3]) > 0 {
					needCookie = false
					isPublic = true
					isOpenApi = false
				} else if len(mMethods) >= 3 && len(mMethods[2]) > 0 {
					isPublic = true
					needCookie = true
					isOpenApi = false
				}

				if !(!isPublic && onlyPublic) {
					if len(brief) > 0 {
						fout.WriteString(fmt.Sprintf("## %s\n\n", brief))
					} else {
						fout.WriteString(fmt.Sprintf("## %s\n\n", method))
					}
				} else {
					if len(brief) > 0 {
						fout.WriteString(fmt.Sprintf("<b> %s</b>\n\n", brief))
					} else {
						fout.WriteString(fmt.Sprintf("<b> %s</b>\n\n", method))
					}
				}
				brief = ""

				fout.WriteString("<ul class=\"my-code-list\">")

				if len(methodPath) > 0 {
					if isPublic {
						if isOpenApi {
							fout.WriteString(genCode(fmt.Sprintf("Url: https://open.qyvideo.net/v2.0/%s?BusinessId=1010&CurTime=1480000000&Sign=xxx<br/>\n", methodPath)))
							fout.WriteString(genCode("OpenAPI Url参数和请求应答简介请参考<a href=\"#103\">链接</a><br/>\n"))
						} else {
							fout.WriteString(genCode(fmt.Sprintf("Url: https://live.qyvideo.net/v2.0/%s<br/>\n", methodPath)))
						}

					} else {
						fout.WriteString(genCode("注意: Private接口, 允许内部调用或OpGateway调用<br/>\n"))
						fout.WriteString(genCode(fmt.Sprintf("RPCPath: /op/%s<br/>\n", methodPath)))
						fout.WriteString(genCode(fmt.Sprintf("OpUrl: https://cop.qyvideo.net/route/op/%s<br/>\n", methodPath)))
					}
				}
				methodPath = ""

				if needCookie && isPublic {
					fout.WriteString(genCode("注意: 请求Http Header中应存在UserKey字段<br/>\n"))
				}
				fout.WriteString("<code>")
				fout.WriteString(genSampleReq(typeInfoMap, method, isPublic, isOpenApi))
				fout.WriteString(genSampleRsp(typeInfoMap, method, isPublic, isOpenApi))
				fout.WriteString("</code>")
				fout.WriteString(fmt.Sprintf("%s</ul>\n\n", explainStr))

				fout.WriteString(fmt.Sprintf("<a name=\"%s\" id=\"%s\"> <b>%s</b> </a>\n\n", typeName, typeName, typeName))
			} else { // normal message or enum
				explainStr += m[6]
				fout.WriteString(fmt.Sprintf("<a name=\"%s\" id=\"%s\"> <b>%s</b> </a>\n\n", typeName, typeName, typeName))
				fout.WriteString(fmt.Sprintf("<ul class=\"my-code-list\">%s</ul>\n\n", explainStr))
			}

			explainStr = ""

			if len(m[2]) > 0 { //message
				fout.WriteString(fmt.Sprintf("| 字段 | 必须/可选/数组 | 类型 | 说明 |\n"))
				fout.WriteString(fmt.Sprintf("|:---------:|:--------:|:----:|:------- |\n"))
				where = IN_MESSAGE
			} else { // enum
				fout.WriteString(fmt.Sprintf("%s\n\n", line))
				where = IN_ENUM
			}

		} else if where == IN_MESSAGE {
			mFields := regField.FindStringSubmatch(line)
			if len(mFields) > 0 { // message field description
				required := mFields[1]
				typeName := mFields[5]
				fieldName := mFields[6]
				explainStr += mFields[8]
				typeNameWithHRef := fmt.Sprintf("<a href=\"#%s\">%s</a>",
					typeName, typeName)
				fout.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					fieldName, required, typeNameWithHRef, explainStr))
				explainStr = ""
			} else { // explains
				mExplains := regExplain.FindStringSubmatch(line)
				if len(mExplains) >= 2 && len(mExplains[1]) > 0 {
					exp := mExplains[1] + "<br/>"
					explainStr += exp
				} else if regEndEnum.MatchString(line) { // message end
					where = OUT
				}
			}
		} else {
			if where == IN_ENUM {
				fout.WriteString(fmt.Sprintf("    %s\n\n", line))
				if regEndEnum.MatchString(line) { // enum end
					where = OUT
				}
			} else { // OUT
				mExplains := regExplain.FindStringSubmatch(line)
				if len(mExplains) >= 2 && len(mExplains[1]) > 0 {
					mExplainPaths := regExplainPath.FindStringSubmatch(line)
					if len(mExplainPaths) >= 3 && len(mExplainPaths[2]) > 0 {
						methodPath = strings.ToLower(mExplainPaths[2])
					} else {
						mExplainBrief := regExplainBrief.FindStringSubmatch(line)
						//log.Print("line: ", line)
						//log.Print("len(mExplainBrief): ", len(mExplainBrief), ", mExplainBrief:", m)

						if len(mExplainBrief) >= 3 && len(mExplainBrief[2]) > 0 {
							brief = mExplainBrief[2]
							//exp := brief + "<br/>"
							//explainStr += exp
						} else {
							//exp := mExplains[1] + "<br/>"
							exp := genCode(mExplains[1])
							explainStr += exp
						}
					}

				}
			}
		}
	}
	log.Print("Parse proto done. module:", file, " methods:", methods,
		" publicMethod:", publicMethods)
	return nil
}

func LowerCase(in string) string {
	return strings.ToLower(in)
}

func renderSrcFile(path, tpl string, data map[string]interface{}, appending bool) error {
	log.Print("Start to render path:", path, " data:", data)

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

func printUsage() {
	fmt.Printf("*****************************************************\n")
	fmt.Printf("md_generator <protoDir> [<outMdFile> [<onlyPublic>] ]\n")
	fmt.Printf("*****************************************************\n")
}

func main() {
	protoDir := "."

	flag.Parse()
	args := flag.Args()
	if len(args) >= 1 {
		protoDir = args[0]
	} else {
		printUsage()
		return
	}

	outMdFile := "./api.md"
	if len(args) >= 2 {
		outMdFile = args[1]
	}

	onlyPublic := false
	if len(args) >= 3 {
		tmp, _ := strconv.Atoi(args[2])
		onlyPublic = (tmp != 0)
	}

	protoList, err := getProtolist(protoDir)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Get proto list:", protoList)
	//methodMap := make(map[string](*[]string))
	//publicMethodMap := make(map[string](*[](MethodInfo)))

	fout, err := os.OpenFile(outMdFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("Open file writeonly failed. err:", err)
		return
	}

	typeInfoMap := TypeInfoMap{}
	for _, p := range protoList {
		typeInfoMap, err = parseTypeInfo(p.file, typeInfoMap)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}

	for _, p := range protoList {
		err := renderMarkdown(p.file, fout, p.module, onlyPublic, typeInfoMap)
		if err != nil {
			log.Fatal(err)
			continue
		}
	}
	_ = fout.Close()

}
