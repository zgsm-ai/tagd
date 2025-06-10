# 项目规则

1. 使用gin框架开发RESTful api
2. 使用gorm框架操作数据库，默认使用sqlite3数据库
3. 使用Gin-Swagger生成swagger文档，相关文件输出到docs目录下
4. 为每个接口编写单元测试

## GO框架&组件

| 场景 | 依赖库 | 说明 |
|------|-------|-------|
| Restful API | gin | 用于实现后端API |
| 数据库 | gorm | 用于实现数据库操作 |
| 数据库 | sqlite/postgresql | 用于实现数据库操作 |
| sqlite3 | github.com/glebarez/sqlite| sqlite3数据库 |
| 日志记录 | go.uber.org/zap | 用于实现日志记录，配合es实现日志存储，搜索 |
| 日志/事件记录 |github.com/olivere/elastic/v7| 利用elasticsearch实现日志/事件记录，该日志用于linux端记录|
| 日志记录 | logrus | 用于实现日志记录(备用) |
| JSON | encoding/json | 用于实现JSON序列化 |
| YAML | gopkg.in/yaml.v3 | 采用v3版本，支持锚点等新机制，用于实现YAML序列化，不使用yaml.v2 |
| CLI | github.com/spf13/cobra | 用于实现CLI |
| swagger支持 | github.com/swaggo/swag | 用于实现swagger支持 |
| swagger支持 | github.com/swaggo/files | 用于实现swagger支持 |
| swagger支持 | github.com/swaggo/gin-swagger | 用于实现swagger支持 |
| 文件生成| github.com/alecthomas/template| 基于模板的文件生成 |
| 认证 | github.com/golang/oauth2 | 身份认证 |
| 定时任务 | github.com/robfig/cron | 定时任务 |
| 消息队列 | github.com/nsqio/go-nsq | 开源消息队列，支持windows/linux（备用） |
| 键值数据库&消息队列 | github.com/go-redis/redis/v8 | 基于redis发布订阅实现的简单消息队列，支持windows/linux |
| 向量库 | github.com/qdrant/qdrant | 支持windows的向量库 |
| 国际化 | github.com/nicksnyder/go-i18n | 实现界面内容的国际化，保证输出内容符合各国语言的特点，比如支持单复数规则 |
