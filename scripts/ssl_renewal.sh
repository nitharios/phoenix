#!/bin/bash

/usr/local/bin/docker-compose -f /home/nitharios/golang/src/github.com/nitharios/phoenix/docker-compose.yml run certbot renew --dry-run \
&& /usr/local/bin/docker-compose -f /home/nitharios/golang/src/github.com/nitharios/phoenix/docker-compose.yml kill -s SIGHUP nginx