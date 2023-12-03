FROM golang:1.20.4-alpine AS build_stage
COPY ./ /go/src/app/
WORKDIR /go/src/app
RUN  go mod download
RUN  go install ./cmd/main.go

FROM alpine AS run_stage
WORKDIR /app_binary
COPY --from=build_stage /go/bin/main /app_binary/
COPY --from=build_stage /go/src/app/web /app_binary/static
RUN chmod +x main
EXPOSE 8080/tcp
ENTRYPOINT ./main

EXPOSE 8080/tcp
CMD [ "main" ]