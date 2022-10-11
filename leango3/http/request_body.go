package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read boy failed:%v", err)
		//记住要返回，不然就还会执行后面的代码
		return
	}

	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	//尝试再去读取，啥也读不到，但是不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		//不会进来这里
		fmt.Fprintf(w, "read the data one more time got error:%v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time:[%s] and read data", string(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {

	body, err := r.GetBody()
	if err != nil {
		fmt.Fprintf(w, "read boy failed:%v", err)
		return
	}
	io.ReadAll(body)

	if r.GetBody == nil {
		fmt.Fprintf(w, "Getbody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values["name"][0]
	fmt.Fprintf(w, "query is %v\n", values)
	fmt.Sprint(name)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "befer parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before pase form %v \n", r.Form)

}
