
# runtime container
FROM alpine:3.19.0

WORKDIR /app
COPY ./build/go-app /app/go-app

EXPOSE 8080
ENTRYPOINT ["/app/go-app", "-port=8080"]