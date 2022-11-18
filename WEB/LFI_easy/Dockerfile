FROM php:7.3-apache

COPY index.php /var/www/html

RUN  chown www-data:www-data  /var/www/html/index.php

RUN  php -r "echo 'flag{'.md5(time()+1).'}';" > /flag.txt

WORKDIR /var/www/html

expose 80
