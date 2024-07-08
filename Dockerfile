FROM golang:1.21-alpine

# Install dependencies
RUN apk --no-cache --update add \
    tzdata \
    ca-certificates \
    openssh \
    git \
    wget \
    curl \
    && update-ca-certificates \
    && rm -rf /var/cache/* /tmp/*

# Install command line tools
RUN go install github.com/revel/cmd/revel@v1.1.2

ENV PATH $PATH:/go/bin
EXPOSE 9000


RUN mkdir -p /go/src
COPY init.sh /usr/local/bin/
RUN chmod u+x /usr/local/bin/init.sh

WORKDIR "/go/src"

CMD ["init.sh"]