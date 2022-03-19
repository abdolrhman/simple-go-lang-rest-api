# simple go-lang customers query
simple go-lang customers query, is an project made to query sqlite db to get customers
with filtering and grouping customers by country according phone number
also pagination and search suppoerted

## Database Support
UGin uses **gorm** as an ORM. So **Sqlite3**, **MySQL** and **PostgreSQL** is supported. You just need to edit **config.yml** file according to your setup. 

**config.yml** content:
```
database:
  driver: "postgres"
  dbname: "ugin"
  username: "user"
  password: "password"
  host: "localhost"
  port: "5432"
```

## Main Models

**/model/customer-model.go** content:
```
// Data is mainle generated for filtering and pagination
type Data struct {
	TotalData    int64
	FilteredData int64
	Data         map[string][]Customer
}

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

type Customer struct {
	ID    int
	Name  string
	Phone string
}

```

## Filtering, Search and Pagination
**UGin** has it's own filtering, search and pagination system. You just need to use these parameters.

**Query parameters:**
```
/customers/?Limit=2
/customers/?Offset=0
/customers/?Sort=ID
/customers/?Order=DESC
/customers/?Search=hello
```

Full: **http://localhost:8081/customers/?Limit=25&Offset=0&Sort=ID&Order=DESC&Search=hello**

## Running

with Docker, firstly build an image:
```
make build-image
```

To run with sqlite:
```
make run-app-sqlite
```


Application will be served at ":8081"

## Logging
very powerful logging logic. There is **application jumiaMonitor.log **, **database log (jumiaMonitor.db.log)** and **access log (jumiaMonitor.access.log)**

### jumiaMonitor.log:
```
INFO 2021-09-19T00:33:32+03:00 Server is starting at 127.0.0.1:8081
ERROR 2021-09-19T00:39:19+03:00 Failed to open log file jumiaMonitor.log
```
### jumiaMonitor.db.log:
```
2021/09/19 00:33:32 /home/user/projects/jumiaMonitor/pkg/database/database.go:76
[0.023ms] [rows:-] SELECT * FROM `posts` LIMIT 1

2021/09/19 00:33:32 /home/user/go/pkg/mod/gorm.io/driver/sqlite@v1.1.5/migrator.go:261
[0.015ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type = "index" AND tbl_name = "posts" AND name = "idx_posts_deleted_at"

2021/09/19 00:33:32 /home/user/go/pkg/mod/gorm.io/driver/sqlite@v1.1.5/migrator.go:32
[0.010ms] [rows:-] SELECT count(*) FROM sqlite_master WHERE type='table' AND name="tags"

2021/09/19 00:33:32 /home/user/projects/jumiaMonitor/pkg/database/database.go:76
[0.011ms] [rows:-] SELECT * FROM `tags` LIMIT 1
```
### jumiaMonitor.access.log:
```
[GIN] 2021/09/19 - 00:33:43 | 200 |    9.255625ms |       127.0.0.1 | GET      "/posts/"
[GIN] 2021/09/19 - 00:41:51 | 200 |     6.41675ms |       127.0.0.1 | GET      "/posts/4"
```

## Routes
Default **UGin** routes are listed below. 

| METHOD  | ROUTE            | FUNCTION                                                      |
|---------|------------------|---------------------------------------------------------------|
| GET     | /posts/          | github.com/yakuter/jumiaMonitor/controller.(*Controller).GetPosts     |

## Gin Running Mode
Gin framework listens **GIN_MODE** environment variable to set running mode. This mode enables/disables access log. Just run one of these commands before running app:
```bash
// Debug mod
export GIN_MODE=debug
// Test mod
export GIN_MODE=test
// Release mod
export GIN_MODE=release
```


## Packages
great open source projects list below: **Gin** for main framework, **Gorm** for database and **Viper** for configuration.
```
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/postgres
go get -u github.com/jinzhu/gorm/dialects/sqlite
go get -u github.com/jinzhu/gorm/dialects/mysql
go get -u github.com/spf13/viper
```
## Middlewares
### 1. Logger and Recovery Middlewares
Gin has 2 important built-in middlewares: **Logger** and **Recovery**. app calls these two in default.
```
router := gin.Default()
```

This is same with the following lines.
```
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
```

### 2. CORS Middleware
CORS is important for API's and UGin has it's own CORS middleware in **include/middleware.go**. CORS middleware is called with the code below.
```
router.Use(include.CORS())
```
There is also a good repo for this: https://github.com/gin-contrib/cors

