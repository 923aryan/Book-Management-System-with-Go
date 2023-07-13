LINUX(debian/arch) commands:
to stop mysql/mariadb -> sudo service mariadb stop
to start mysql/mariadb -> sudo service mariadb start
to check running servers -> sudo netstat -tulpn | grep 3306
to go into mysql -> sudo mysql
	to show databases -> SHOW DATABASES;
	to create database for eg. simplerest -> CREATE DATABASE simplerest;
	to get details about some database for eg. simplerest -> SELECT user, host, db FROM mysql.db WHERE db = 'simplerest';
	to know about priviledges about database with username for eg.aryanThakur -> SHOW GRANTS FOR 'aryanThakur'@'localhost';


