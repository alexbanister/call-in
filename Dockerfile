# 1 choose a compiler OS
FROM golang:alpine AS builder

# 2 (optional) label the compiler image
LABEL stage=builder

# 3 (optional) install any compiler-only dependencies
RUN apk add --no-cache gcc libc-dev git
WORKDIR /workspace

# 4 copy all the source files
COPY . .

# 5 build the GO program
RUN cd src && go get -d -v
RUN cd src && CGO_ENABLED=0 GOOS=linux go build -o megaton .

# 6 choose a runtime OS
FROM alpine AS final

# 7
WORKDIR /

RUN apk add --no-cache go

# 8 copy from builder the GO executable file
COPY --from=builder /workspace/src/megaton /app/megaton
COPY --from=builder /workspace/config.yml /

EXPOSE 8080
# 9 execute the program upon start
CMD [ "./app/megaton" ]
