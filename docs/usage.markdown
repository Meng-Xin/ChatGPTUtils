// @Summary	创建会话场景
// @Produce	json
// @Param		request	body		request.AddToScenesRequest	true	"新建场景：1聊天、2画图可选"
// @Success	200		{object}	public.Response				"成功"
// @Failure	400		{object}	public.Response				"请求错误"
// @Router		/api/v1/openai/addScenes [post]

## swagger用法
### 1.@Summary	创建会话场景
### 2.@Produce  数据格式
### 3.@Param    数据
+ request 参数名称
+ body 数据存放位置 head|body
+ request.AddToScenesRequest 具体类型
+ true 是否必填
+ “” 注释
### 4.@Success  成功返回
+ 200 HTTP.Status
+ {object} 对象类型
+ public.Response 具体参数烈性
+ 注释信息
### 5.@Failure  失败返回
+ 400 HTTP.Status
+ {object} 对象类型
+ public.Response 具体参数烈性
+ 注释信息
### 6.@Router   路由
+ Path
+ RequestType

> @Param		chatMsg		body		request.AddToScenesRequest		true "请求信息"
> param解析  参数名称  参数类型 head|body   参数类型  是否必填   注释
