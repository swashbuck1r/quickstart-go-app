
# runtime container
FROM alpine:3.19.0

WORKDIR /app
COPY ./build/quickstart-go-app /app/quickstart-go-app

EXPOSE 8080
ENTRYPOINT ["/app/quickstart-go-app", "-port=8080"]