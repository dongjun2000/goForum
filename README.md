# goForum
用Go开发的一个论坛，便于Go爱好者学习和借鉴。

### 部署

* 在本地开发环境编译 linux 上运行的可执行程序

```
$ GOOS=linux GOARCH = amd64 go build
```

* 在服务器上项目根目录配置 `config.json` 配置文件

```json
{
  "App": {
    "Address": "127.0.0.1:8000",
    "Static": "public",
    "Log": "logs",
    "Locale": "locales",
    "Language": "zh"
  },
  "Db": {
    "Driver": "mysql",
    "Address": "127.0.0.1:3306",
    "Database": "goforum",
    "User": "xxx",
    "Password": "xxx"
  }
}
```

* 确保 `logs` 日志目录对 `Web` 用户具有写权限

```
$ sudo chmod 777 logs
```

* 初始化数据库

在MySQL数据库中创建 `goforum` 数据库，并初始化对应数据表

```sql
create table users (
  id         bigint auto_increment primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at datetime(6) not null
);

create table sessions (
  id         bigint auto_increment primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    bigint NOT NULL,
  created_at datetime(6) not null
);

create table threads (
  id         bigint auto_increment primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    bigint NOT NULL,
  created_at datetime(6) not null
);

create table posts (
  id         bigint auto_increment primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    bigint NOT NULL,
  thread_id  bigint NOT NULL,
  created_at datetime(6) NOT NULL
);
```

* 在项目根目录下运行 `goForum` 二进制文件启动应用

```
$ ./goForum
```

* 部署 Nginx 做反向代理，nginx.conf

```
server {
    listen      80; 
    server_name goforum.test www.goforum.test;
    
    # 静态资源交由 Nginx 管理，并缓存一天
    location /static {
        root        /var/www/goForum/public;
        expires     1d;
        add_header  Cache-Control public;
        access_log  off;
        try_files $uri @goweb;
    }
    
    location / {
        try_files /_not_exists_ @goweb;
    }
    
    # 动态请求默认通过 Go Web 服务器处理
    location @goweb {
        proxy_set_header Host $http_host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Scheme $scheme;
        proxy_redirect off;
        proxy_pass http://127.0.0.1:8000;
    }
    
    error_log /var/log/nginx/goforum_error.log;
    access_log /var/log/nginx/goforum_access.log;
}
```

* 通过 `Supervisor` 维护应用守护进程

我们需要借助第三方进程监控工具帮我们实现 `Go Web` 应用进程以后台守护进程的方式运行。

首先创建对应的 `Supervisor` 配置文件 `/etc/supervisor/conf.d/goforum.conf`，这里需要设置进程启动目录及命令、进程意外挂掉后是否自动重启、以及日志文件路径等：

```
[program:goForum]
process_name=%(program_name)s
directory=/var/www/goForum
command=/var/www/goForum/goForum
autostart=true
autorestart=true
user=root
redirect_stderr=true
stdout_logfile=/var/www/goForum/logs/goForum.log
```

```
$ supervisorctl reload
$ supervisorctl update
$ supervisorctl start goforum
```

这里的部署方式比较原始，对于多人协作的大型项目，可以借助持续集成工具(比如 Jenkins)进行自动化部署。