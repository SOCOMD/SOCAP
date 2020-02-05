FROM centos/httpd
maintainer Glowbal

# Install required dependencies
RUN yum -y install php php-mysql php-core php-zip php-zlib php-json \
 php-curl php-data php-dom php-fileinfo php-session php-sockets \
 php-tokenizer install sqlite

RUN ln -sf /dev/stderr /var/log/httpd/error_log
RUN ln -sf /dev/stderr /var/tmp/php-error.log
# Configuration and installation
COPY ./.docker/httpd.conf /etc/httpd/conf/httpd.conf
ADD web /var/www/html/
