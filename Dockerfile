FROM golang:1.15.2-alpine3.12 as VENDOR

#vendor- acts as wrapper to compose all depends
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add -no-cache ca-certificates git
WORKDIR /src
COPY go.mod ./
RUN go mod download
RUN go mod VENDOR

#use built executable from above and set creds
FROM VENDOR as builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

#Create final image from builder
FROM new as final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /app
USER nobody:nobody
ENTRYPOINT ["/app"]

#TBC - User definition:
    # Cur user is set by `nobody` user and not root. For privilege reason, may change
    # ssl certificate issues with go-sql-driver