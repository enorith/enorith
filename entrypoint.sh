#!/bin/sh
ln -sf /usr/share/zoneinfo/${APP_TIMEZONE} /etc/localtime && echo "${APP_TIMEZONE}" > /etc/timezone
/app/enorith