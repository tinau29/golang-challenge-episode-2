FROM golang:latest AS compiler
COPY . /project
WORKDIR /project
RUN apt update && apt install upx-ucl -y
RUN go install github.com/pwaller/goupx@latest
RUN go build -a -tags netgo -ldflags '-s -w -extldflags "-static"' -o app.run .
RUN goupx app.run

FROM scratch
COPY --from=compiler /project/app.run /server
CMD ["/server"]
