# Multistage build
# Builds go binary in first stage, copy binary and install runtime
# deps in second stage

# build stage
FROM alpine:3.11 as builder
RUN apk add --no-cache ca-certificates go
ADD /src /src
RUN cd $GOPATH && go get && cd /src && go build -o megaton

# final stage
FROM alpine:3.11
RUN apk update && apk upgrade
RUN apk --no-cache add ca-certificates
COPY --from=build-env /src/megaton ./
EXPOSE 8080
CMD ["./megaton"]
