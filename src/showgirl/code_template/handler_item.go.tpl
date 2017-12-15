{{range $key, $method := .methods}}
//auto generated.
func Handle{{$method}}(hdr *{{$.pkg}}.STUserTrustInfo, req *{{$.pkg}}.ST{{$method}}Req) (*{{$.pkg}}.STRspHeader, *{{$.pkg}}.ST{{$method}}Rsp) {

    /////////////////////////////////////////////////////////////////////
    //auto generated, please remove this section and write your code here.
    errno := {{$.pkg}}.EErrorTypeDef_RESULT_NOT_IMPLEMENTED
    errmsg := "NOT IMPLEMENTED"
    return &{{$.pkg}}.STRspHeader{&errno, &errmsg, proto.Int64(0), nil, nil}, nil
    /////////////////////////////////////////////////////////////////////

}
{{end}}