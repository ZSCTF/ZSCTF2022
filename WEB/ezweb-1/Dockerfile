FROM luisffc/imagick:latest

COPY index.php /var/www/html

COPY 1.xml  /etc/ImageMagick-6/policy.xml

RUN  chown www-data:www-data  /var/www/html/index.php

RUN  php -r "echo 'flag{'.md5(time()+1).'}';" > /var/www/html/flag.php

WORKDIR /var/www/html

expose 80
