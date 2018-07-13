FROM golang:1.9.2 as basic

MAINTAINER jiangd@vmware.com

COPY src/. /go/src/github.com/vmware/harbor/src
WORKDIR /go/src/github.com/vmware/harbor/src/ui

RUN go build -a -o /go/bin/harbor_ui

COPY src/ui/views /go/bin/views
COPY src/ui/static /go/bin/static
COPY src/favicon.ico /go/bin/favicon.ico

RUN mkdir /go/bin/harbor/
COPY VERSION /go/bin/harbor/UIVERSION

FROM vmware/photon:1.0

COPY --from=basic /go/bin/favicon.ico /go/bin/harbor/UIVERSION /go/bin/harbor_ui  /harbor/
COPY --from=basic /go/bin/views /harbor/views
COPY --from=basic /go/bin/static /harbor/static
COPY --from=basic /go/bin/favicon.ico /go/bin/harbor/UIVERSION /go/bin/harbor_ui /harbor/

RUN chmod u+x /harbor/harbor_ui

WORKDIR /harbor/
ENTRYPOINT ["/harbor/harbor_ui"]

EXPOSE 80

