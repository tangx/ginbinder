# ginbinder

一次绑定 Request 中所有需要的数据

![gopher-binding](docs/imgs/gopher-binding.jpeg)


## Usage

1. **废弃/不可用**: 弃用原生 tag `form` tag。
2. **保持**: 使用 tag `path` 绑定路径中的参数。 **作用于某个字段**
    + 就是 `example.com/:some/:path` 中 **冒号后面的**
3. **保持**: 使用 tag `header` 绑定 header。 **作用于某个字段**
4. **新增**: 新增 tag `query` tag 绑定通过 Query 传递的参数。 **作用于某个字段**
    + 就是 `example.com/some/path?a=1&b=2` 中 问号后面的那一串
5. **新增**: 新增 tag `cookie` 绑定 cookie 中 **简单** 的键值对。 **作用于某个字段**
6. **新增**: 新增 tag `body` 绑定 `request.Body` 中的数据。 **必须作用于一个 `struct` 上。**
    + Body 数据的解析器通过 `content-type` 判断， 默认为 `Json`
    + Body 中的字段不能包含 `path, query, cookie, header` 这些 tag， 否则会 panic。
    + `body` **只能有一个**
7. **新增**: 新增 tag `mime` 在代码中 **指定** 解析器， 优先于 **Request** 中的 `content-type`。 与 `body` 的协同，无法单独使用。
    + `json`: `application/json`
    + `yaml`: `application/x-yaml`
    + `xml`: `application/xml`
    + `form`: `form` 的行为逻辑 **必须** 依赖 request 中的 `Content-Type`。
        + `application/octet-stream`: 默认， `net/http/request.go#L1212` 中定义。
        + `application/x-www-form-urlencoded`: form 表单
        + `multipart/form-data`: 多行数据

## Bind and Bindonly

1. `BindRequest` / `ShouldBindRequest` 强制进行参数有效性 `validate` 验证。
2. `BindOnlyRequest` / `ShouldBindOnlyRequest` 不进行参数有效性 `validate` 验证。

### 按照 Request 结构管理参数

```go
type Params struct {
	Name          string `path:"name"`
	Age           int    `query:"age,default=18"`
	Money         int32  `query:"money" binding:"required"`
	Authorization string `cookie:"Authorization"`
	UserAgent     string `header:"User-Agent"`
	Data          struct {
		Replicas *int32 `json:"replicas" yaml:"replicas" xml:"replicas" form:"replicas"`
	} `body:"body" mime:"json"`
}
```

如果层抓包捕获过 http 请求并观察过其结构。 那么应该很好理解这个格式。

这里放一张 `Params` 与 `Request` 的关系图以帮助理解。

![params-request-relationship.png](./docs/imgs/params-request-relationship.png)

> 图中有一个错误， `12/13 行` 的 `tag` 应该是 `query`。

按照 Request 结构设计参数很 **常用** 也很清晰。

![result-demo1](./docs/imgs/result-demo1.png)


### 绑定 Json Body 到 map 中

[/__example__/bind_map.go](/__example__/bind_map.go)

注意， 如果 Data 的数据类型与传入不匹配， 将会引发 panic。

```go
type paramsMap struct {
	Data map[string]interface{} `body:"body" mime:"json"`
}

func handlerBindMapAsJson(c *gin.Context) {
	parmas := &paramsMap{}

	err := ginbinder.ShouldBindRequest(c, parmas)
	if err != nil {
		panic(err)
	}

	c.JSON(200, parmas)
}
```

## exmaple

Demo 参考 [exmaple](./__example__)

使用 VsCode 的 Rest Client 可以快速测试


```bash
cd __example__

go run .
```