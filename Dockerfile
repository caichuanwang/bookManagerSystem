FROM golang:1.18.8 AS BUILD
WORKDIR /root/bookManagerSystem
COPY . .
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bms main.go


FROM scratch AS PROD
WORKDIR /usr/local/bms
COPY --from=BUILD /root/bookManagerSystem .
EXPOSE 8888


