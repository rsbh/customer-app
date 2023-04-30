FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o ./dist/customer-app main.go


FROM alpine:3.14
COPY --from=builder /app/dist /usr/bin/
EXPOSE 8080
ENTRYPOINT [ "customer-app" ]









