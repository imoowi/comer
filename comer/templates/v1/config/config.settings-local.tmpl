application: # dev开发环境 test测试环境 prod线上环境
  mode: dev
  name: {{.moduleName}} # 服务名称
server:
  host: 0.0.0.0 # 服务器ip，默认使用 0.0.0.0
  port: 8000 # 服务端口号
  readtimeout: 60 # 读超时时间
  writertimeout: 60 # 写超时时间
logger:
  path: runtime/logs/log # 日志存放路径
  stdout: "" # 日志输出，file：文件，default：命令行，其他：命令行
  level: trace # 日志等级, trace, debug, info, warn, error, fatal
  maxAge: 168h # 日志最长保存时间，7天, ns、us、ms、s、m、h
  rotationTime: 24h # 日志切割级别
ratelimit:
  # 每秒放多少个令牌
  cap: 1000
  # 每秒取多少个令牌
  quantum: 1000
jwt:
  secret: {{.moduleName}}-admin # token 密钥，生产环境时及的修改
  timeout: 2000h0m0s # token 过期时间 格式：0h0m0s
  refresh_token_timeout: 0h5m0s # token 过期时间减去的时间，用于刷新token
mysql:
  dsn: root:123456@tcp(127.0.0.1:3306)/{{.dbName}}?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  casbin: root:123456@tcp(127.0.0.1:3306)/{{.dbName}}
# influxdb:
#   addr: http://127.0.0.1:8086
#   token: [token string]
#   org: [orgnization name]
#   bucket: [bucket name]
#   testSwitchOn: false
redis:
  addr: 127.0.0.1:6379
  password: "123456"
  db: 0
cache:
  driver: redis
  prefix: "{{.dbName}}:cache"
