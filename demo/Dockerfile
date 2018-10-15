#DANGLING IMAG
FROM golang:1.6.2 AS BUILD_BASE

MAINTAINER ****@*******

COPY . /go/src/demo

WORKDIR  /go/src/demo/cmd/demo

RUN go build -v  -o /go/bin/hello  -ldflags "-X main.VERSION=1.0 -s -w" *.go

WORKDIR /go/bin/

ENTRYPOINT ["hello","-city=hangzhou"]

#IMAGE 2
FROM debian

MAINTAINER ****@*******

ENV GOPATH /go \
    PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

COPY --from=BUILD_BASE /go/bin/hello /bin/hello

EXPOSE 8000
EXPOSE 9000

ENTRYPOINT ["hello","-city=hangzhou"]




#FROM debian:latest
#
#MAINTAINER ****@*******
#
#LABEL "demo"="demo"
#
#ENV GOPATH /go \
#    PATH $GOPATH/bin:/usr/local/go/bin:$PATH
#
#RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
#
#ADD ./cmd/demo/hello /bin/hello
##COPY https://docs.projectcalico.org/v2.6/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml /bin/hello
#
#EXPOSE 8000
#
#ENTRYPOINT ["hello","-city=hangzhou"]
#
#


FROM ubuntu:14.04

MAINTAINER ****@*******

RUN apt-get update \
    && apt-get install -y libldap2-dev \
    && rm -r /var/lib/apt/lists/*

COPY . /harbor
WORKDIR /harbor/ui

ADD harbor_ui /usr/bin/harbor_ui

ENV MYSQL_USR root \
    MYSQL_PWD root \
    REGISTRY_URL localhost:5000

COPY views /views
COPY static /static
COPY favicon.ico /favicon.ico
COPY Deploy/jsminify.sh /tmp/jsminify.sh

RUN chmod u+x /usr/bin/harbor_ui \
    && sed -i 's/TLS_CACERT/#TLS_CAERT/g' /etc/ldap/ldap.conf \
    && sed -i '$a\TLS_REQCERT allow' /etc/ldap/ldap.conf \
    && /tmp/jsminify.sh /views/sections/script-include.htm /static/resources/js/harbor.app.min.js \
    && rm -rf /var/cache \
    && rm -rf /go/src/github.com/vmware/harbor

ENTRYPOINT ["/usr/bin/harbor_ui"]

EXPOSE 80