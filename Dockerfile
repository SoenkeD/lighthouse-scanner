FROM golang:latest as go-builder

WORKDIR /build
COPY scan-control /build

RUN go get all

RUN mkdir bin/
RUN go build -o bin/lighthouse-scanner main/main.go


FROM node:slim
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD true

RUN apt-get update && apt-get install curl gnupg -y \
  && curl --location --silent https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
  && sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list' \
  && apt-get update \
  && apt-get install google-chrome-stable -y --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY lighthouse-client/package.json .
COPY lighthouse-client/package-lock.json .
RUN npm ci

RUN useradd lighthouse
RUN mkdir /home/lighthouse
RUN chown lighthouse /app /home/lighthouse
USER lighthouse

COPY --from=go-builder build/bin/ .
COPY lighthouse-client/main.js .

ENTRYPOINT ["/app/lighthouse-scanner"]