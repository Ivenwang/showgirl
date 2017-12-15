package utils

import (
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

//SendAndRecv 通用发收包
func SendAndRecv(method, urlStr string, body io.Reader, mapHeaderParam map[string]string, flowid int64) ([]byte, error) {
	//构造请求
	client := &http.Client{

		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					Warn(flowid, "SendAndRecv DialTimeout error, err = %s", err)
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}

	request, _ := http.NewRequest(method, urlStr, body)

	//设置header参数
	for idx, value := range mapHeaderParam {
		request.Header.Set(idx, value)
	}

	//执行请求
	rsp, err := client.Do(request)
	if err != nil {
		Warn(flowid, "SendAndRecv client do error, err = %s", err)
		return nil, err
	}

	defer rsp.Body.Close()

	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		Warn(flowid, "SendAndRecv read body error, err = %s", err)
		return nil, err
	}

	Debug(flowid, "SendAndRecv rsp data, body = %s", rspBody)

	return rspBody, nil
}
