FROM golang:alpine3.12 AS build
RUN mkdir /app 
ADD main.go /app/ 
WORKDIR /app 
RUN go build -o main . 

FROM alpine:3.12
RUN mkdir /app
WORKDIR /app 
COPY --from=build /app/main .
CMD ["/app/main"]