# gin-basic
golang 和 gin web framework 的基本練習

### 建立 PostgreSQL users Table
```sql
CREATE TABLE users
(
id                  BIGSERIAL,
name                TEXT,
PRIMARY KEY (id)
); 
```
註: 目前程式在啟動時，若 users Table 不存在，會自動去把 Table 建出來，可以不用手動下上面的SQL沒關係。

### 測試環境建立方式
#### 方式一
##### 先用 docker 起 postgres_db
```docker
docker container run -d -p 5432:5432 --name postgres_db -v postgres_data:/var/lib/postgresql/data -e POSTGRES_PASSWORD=pw_postgres postgres:latest
```
##### 再跑本專案的 image : vant/gin-basic
```
docker container run -d --name go_app -p 8080:8080 --link postgres_db \
-e DB_HOST=postgres_db \
-e DB_PORT=5432 \
-e DB_USER=postgres \
-e DB_PASSWORD=pw_postgres \
-e DB_NAME=postgres \
vant/gin-basic
```
#### 方式二
##### 直接用 docker compose 啟動
```
docker compose up
```

### Test files
如果你有裝 VS code 的 REST Client Extension 的話，可以再 ./api-test/ 目錄下，直接修改檔案內容去測試
