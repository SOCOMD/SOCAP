FROM centos/httpd
maintainer Glowbal

# Install required dependencies
RUN yum -y install php php-mysql php-core php-zip php-zlib php-json \
 php-curl php-data php-dom php-fileinfo php-session php-sockets \
 php-tokenizer install sqlite

# Configuration and installation
COPY ./.docker/httpd.conf /etc/httpd/conf/httpd.conf
ADD web /var/www/html/

RUN chmod 777 ./var/www/html/data
RUN chmod 777 ./var/www/html/images/maps