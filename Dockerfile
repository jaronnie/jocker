FROM centos:7.9.2009

LABEL MAINTAINER jaron@jaronnie.com

COPY ./dist/jocker_linux_amd64/jocker /usr/bin/jocker

RUN yum -y install bash-completion \
    && gvm completion bash > /etc/bash_completion.d/gvm

