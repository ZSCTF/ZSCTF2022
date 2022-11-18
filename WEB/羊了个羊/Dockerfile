FROM luisffc/imagick:latest

COPY index.html /var/www/html

COPY cShc.html /var/www/html

RUN  chown www-data:www-data  /var/www/html/index.html

WORKDIR /var/www/html

expose 80
