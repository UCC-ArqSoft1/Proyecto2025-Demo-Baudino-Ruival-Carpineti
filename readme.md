# Run MySQL

```
docker run --name mysql-users \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_DATABASE=backend \
    -p 3306:3306 \
    mysql:5.6
```