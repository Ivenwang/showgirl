appname  = showgirl
ListenTCP4 = true
httpaddr = "0.0.0.0"
httpport = ${httpport}
runmode  = dev
copyrequestbody = true
RecoverPanic = true


module = ${module}


[log]
;DEBUG
;INFO
;WARN
;FATAL
level = ${level}

[DefaultSvr]
url = ${url}
[AccountSvr]
[RecommendSvr]

[GoFront]
MaxHttpBodyLength = 8192

[CommonConfig]
;依赖微服务环境
DepEnv = ${DepEnv}
SecretKey = IamSecretKey
CookieActiveTime = 31536000

[Mysql]
MysqlConn = ${MysqlConn}
MysqlConnSlave = ${MysqlConnSlave}
MaxConnent = 30
MaxIdle = 6

[Session]
cookiekey = Z9W4BAtQGHMMu0txqmVQGkRwVjW8R7Lb
maxonlinenum = 1

[zhanweifu]
name=value



