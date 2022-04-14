![Screenshot at 2022-04-13 23-00-58](https://user-images.githubusercontent.com/38928236/163292499-23d754da-f1fb-4213-bd7e-291c5fed8bd2.png)
<br>
This is my version of Have I been pwned web app written in Go. App is searching for pwned passwords stored in mysql, and is working very very fast. For example it can find data in database of 3 billions records, just in milliseconds.

So, first you need to create your own database and table in mysql.
```
CREATE DATABASE data;
```
```
CREATE TABLE pwned (idproducts INT(6) AUTO_INCREMENT PRIMARY KEY, email VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);
```
and add some data...
```
INSERT INTO pwned (email, password) VALUES ('joe@example.com', 'password');
```
Do not forget to change mysql username and password in main.go file.

<b>db, err = sql.Open("mysql", "root:password.@tcp(localhost:3306)/data")</b>

Run the program and enjoy :)
```
sudo go run main.go
```
The program is running on port 8080, and of course you can change it.

In case of bigger quantity of data I recommend you to create index in database:
```
ALTER TABLE pwned ADD INDEX (email);
```

