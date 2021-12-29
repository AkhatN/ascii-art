FROM golang:1.16.7-alpine3.13 

LABEL version="1.0" \
      "site.name"="ascii-art-web" \
      maintainers="AkhatN" \
      release-date="September 7, 2021" \
      description="This project does ascii-art-web" \
      authors="AkhatN"

WORKDIR /app

COPY go.mod .

RUN mkdir art
COPY art art/

RUN mkdir ascii
COPY ascii ascii/

RUN mkdir css
COPY css css/

RUN mkdir templates
COPY templates templates/

RUN mkdir img
COPY img img/

COPY README.MD .
COPY server.go .

RUN go build -o server .

CMD ["./server"]