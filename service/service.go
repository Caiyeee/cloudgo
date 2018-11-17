package service

import(
	"net/http"
	"github.com/urfave/negroni"
	"github.com/unrolled/render"
	"github.com/gorilla/mux"
)

func NewServer() *negroni.Negroni {
	// Classic() 返回一个 Negroni 实例，会添加异常回复，日志和静态文件服务器三个中间件
	server := negroni.Classic()

	// 创建一个路由器
	router := mux.NewRouter()
 
	format := render.New(render.Options {
		IndentJSON : true,
	})

	// 初始化路由器
	router.HandleFunc("/hello/{nname}", handler(format)).Methods("GET")

	// 添加中间件
	server.UseHandler(router)

	return server
}


func handler(format *render.Render) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 获得请求
		vars := mux.Vars(request)
		// 解析请求内容
		name := vars["name"]
		// 渲染模板
		format.JSON(writer, http.StatusOK, map[string]string{name : name})
	}
}
