FROM golang:1.19-alpine

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh isi folder day 7 ke app
ADD ../day7/. /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["/app/main"]