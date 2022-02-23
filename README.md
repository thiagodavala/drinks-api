# Drinks API

## BD

```
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=senha -d mysql/mysql-server:latest
```

```
docker exec -it mysql mysql -uroot -p
```

```
CREATE USER 'sistema'@'localhost' IDENTIFIED BY 'senha';
GRANT ALL PRIVILEGES ON * . * TO 'sistema'@'localhost';
FLUSH PRIVILEGES;
```

```
create database drinks;
use drinks;
```

```
CREATE TABLE cocktails (
id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(30) NOT NULL,
ingredients VARCHAR(256) NOT NULL,
method VARCHAR(256) NOT NULL,
garnish VARCHAR(256) NOT NULL);
```

```
INSERT INTO cocktails (name, ingredients, method, garnish)
VALUES ("Caipirinha", "60 ml Cachaça; 1 Lime cut into small wedges; 4 Teaspoons White Cane Sugar;", "Place lime and sugar into a double old fashioned glass and muddle gently. Fill the glass with cracked ice and add Cachaça. Stir gently to involve ingredients.","N/A");
```

## WEB API

```
go get -u github.com/gin-gonic/gie
go mod init drinks-api
go mod tidy
go get -u github.com/swaggo/swag/cmd/swag
swag init
```

