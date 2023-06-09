# ======================
#  GO FIRST STAGE
# ======================

FROM golang:latest as builder
USER ${USER}
WORKDIR /usr/src/carbide-backend
COPY go.mod \
  go.sum ./
RUN go mod download
COPY . ./
ENV GO111MODULE="on" \
  GOARCH="amd64" \
  GOOS="linux" \
  CGO_ENABLED="0"
RUN apt-get clean \
  && apt-get remove

# ======================
#  GO FINAL STAGE
# ======================

FROM builder
WORKDIR /usr/src/carbide-backend
RUN apt-get update \
  && apt-get install -y \
  make \
  vim \
  build-essential
COPY --from=builder . ./usr/src/carbide-backend
RUN make goprod
EXPOSE 4000
CMD ["./main"]