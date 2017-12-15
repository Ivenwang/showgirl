package main

import (
	"os"
	"bufio"
	"strings"
	"gopkg.in/ini.v1"
	"fmt"
	"io"
	"errors"
	"regexp"
)

func init() {

}

func parseFile(templatefile, filterfile, descfile string) error {

	tf, err := os.Open(templatefile)
	if err != nil {
		return err
	}
	defer tf.Close()

	tfreader := bufio.NewReader(tf)

	df, err := os.OpenFile(descfile, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer df.Close()

	dfwriter := bufio.NewWriter(df)

	sectionRe,_ := regexp.Compile("\\[(.*)\\]")
	varRe,_ := regexp.Compile("\\$\\{(.*)\\}")
	profile, err := loadProfile(filterfile)

	if err != nil {
		return err
	}

	//初始化为默认Section
	curSection := ""

	//逐行替换变量
	for {
		line, err := tfreader.ReadString('\n')
		if err != nil {
			if (err == io.EOF) {
				break
			}
		}
		line = strings.TrimSpace(line)

		//非注释内容
		if !strings.HasPrefix(line, ";") {
			submatch := sectionRe.FindSubmatch([]byte(line))
			if len(submatch) > 1 {
				curSection = string(submatch[1])
			}
			key := varRe.FindSubmatch([]byte(line))
			if len(key) > 1 {
				keystr := string(key[1])
				section, _ := profile.GetSection(curSection)
				if section != nil {
					value,_ := section.GetKey(keystr)

					if value != nil {
						line = strings.Replace(line, "${" + keystr + "}", value.String(), -1)
					} else {
						return errors.New("unkown variable " + keystr)
					}
				} else {
					return errors.New("unkown variable " + keystr)
				}
			}
		}


		dfwriter.WriteString(line)
		dfwriter.WriteString("\n")

	}
	dfwriter.Flush()

	return nil
}

func loadProfile(file string) (*ini.File, error)   {
	cfg, err := ini.Load(file)

	return cfg, err
}

func main()  {
	arglen := len(os.Args)
	if arglen < 4 {
		fmt.Println("usage:go run filter_config.go templatefile filterfile descfile")
		os.Exit(-1)
	}
	err := parseFile(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}