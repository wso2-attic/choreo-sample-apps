FROM php:8.1-apache-buster

# enable apache mod_rewrite
RUN a2enmod rewrite

# copy source files
COPY . /var/www/html

# these directories need to be writable by Apache
RUN chown -R www-data:www-data /var/www/html/storage /var/www/html/bootstrap/cache

# copy env file for our Docker image
COPY .env /var/www/html/.env

RUN php artisan config:cache && \
    php artisan route:cache && \
    chmod 777 -R /var/www/html/storage/ && \
    chown -R www-data:www-data /var/www/

# DB migration
RUN php artisan migrate

# Create a new user with UID 10014
RUN useradd -u 10014 -m choreouser
USER 10014

### Docker image metadata
VOLUME ["/var/www/html/storage", "/var/www/html/bootstrap/cache"]
