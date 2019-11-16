## httpclient包

- 加载httpclient包

```
"github.com/meixiu/httpclient"
```

- 普通GET请求

```
resp, err := httpclient.Get("http://127.0.0.1:8081/get", nil)
fmt.Println(resp.String(), err)
```

- 带参数的GET请求

```
// 支持自定义 struct (添加 url:"name" 标签来指定参数名)
type queryParams struct {
    Foo  string `url:"foo"`
    Bar  string `url:"bar"`
}
resp, err := httpclient.Get("http://127.0.0.1:8081/get", queryParams{
    "foo1", 
    "bar1",
})
fmt.Println(resp.String(), err)


// 支持原生 url.Values 和 支持map[string]string 
b := url.Values{
    "foo": []string{"foo1"},
    "bar": []string{"bar1"},
}
resp, err := httpclient.Get("http://127.0.0.1:8081/get", b)
fmt.Println(resp.String(), err)
```

- FORM表单数据格式的POST请求

```
// 和GET请求一样, 支持 struct 和 url.Values, map[string]string 数据类型
resp, err := httpclient.PostForm("http://127.0.0.1:8081/post", queryParams{
    "name2", "value2",
})
fmt.Println(resp.String(), err)
```

- JSON数据格式POST请求

```
// 支持所有可以转换成JSON的数据类型
d := interface{}{"foo", "bar"}
resp, err := httpclient.PostJson("http://127.0.0.1:8081/post", d)
fmt.Println(resp.String(), err)

// 将请求结果映射到map或者struct
resultData := make(map[string]interface{})
err := resp.Decode(&resultData)
fmt.Println(resultData, err)
```

- 如果需要自定义Header或者其它配置, 可以新建一个client对象

```
// 新建client对象
client := httpclient.New()

// 设置基本请求地址
client.Base("http://127.0.0.1:8081")

// 开启或者关闭调试
client.SetDebug(true)

// 设置超时时间  默认为30秒, 设置为0 则没有超时
client.SetTimeout(1 * time.Second)

// 添加请求头
client.SetHeader("TestHeader", "new header add")

// 发送请求
resp, err := client.Get("/get", nil)
fmt.Println(resp.String(), err)
```