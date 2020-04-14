FROM ubuntu:18.04

ADD ./webserver/webserver /usr/local/bin
ADD ./runserver.sh /usr/local/bin

ENTRYPOINT /usr/local/bin/runserver.sh

EXPOSE 8338
