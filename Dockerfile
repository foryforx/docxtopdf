FROM golang:1.10 
RUN go version
RUN apt-get update
# RUN apt-get install -y build-essential libssl-dev libffi-dev python-dev
RUN apt-get install -y libreoffice
ENV GOOS=linux
WORKDIR /go/src/app
ADD . /go/src/app/
WORKDIR /go/src/app

# RUN go build -o main .
CMD ["go","run","main.go","sample2.docx","2"]




