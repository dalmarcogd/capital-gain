FROM golang:1.18-buster

ENV GOOS=linux \
  GO111MODULE="on"

RUN mkdir -m 700 /root/.ssh; \
  touch -m 600 /root/.ssh/known_hosts; \
  ssh-keyscan github.com > /root/.ssh/known_hosts; \
  git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /code

COPY . .

# Download Go dependencies (include private modules)
RUN --mount=type=ssh go mod download

# Build app
RUN go install -ldflags="-s -w" ./cmd/capgain

ENTRYPOINT ["capgain", "calculator"]
