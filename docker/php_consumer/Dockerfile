FROM php:7.4-fpm-alpine AS build

ARG TZ=UTC
ARG UID=1000
ARG GID=1000

RUN apk add --no-cache shadow git libzip-dev icu-dev autoconf build-base rabbitmq-c

RUN docker-php-ext-install zip pdo_mysql intl
RUN pecl install amqp && docker-php-ext-enable amqp

COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

RUN usermod -u $UID www-data && groupmod -g $GID www-data

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 9000

FROM build AS production

USER www-data:www-data

FROM build AS development

RUN apk update && apk add --no-cache $PHPIZE_DEPS

RUN pecl install xdebug && docker-php-ext-enable xdebug

USER www-data:www-data

