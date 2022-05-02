FROM golang:1.17-alpine as build
RUN mkdir /build
ADD ./Server /build
WORKDIR /build
RUN go mod tidy
RUN go build -o server .

FROM alpine
COPY --from=build /build/server /app/
COPY --from=build /build/static /app/static
WORKDIR /app
CMD ["./server"]
