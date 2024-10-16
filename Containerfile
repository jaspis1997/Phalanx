FROM golang:1.22-alpine3.20 as BackendBuilder

RUN apk update && apk add make

ARG APPLICATION="phalanx"

WORKDIR /app

# Copy source code
COPY ./backend .

RUN go mod download
RUN make container-build APP_BIN=${APPLICATION}

FROM node:20.15.1-alpine3.20 as FrontBuilder

WORKDIR /app
RUN apk update

COPY frontend/ ./

RUN yarn
RUN yarn run build

FROM debian:buster-slim

RUN mkdir /app && mkdir /app/Videos /app/Thumbnails
RUN apt update && apt install -y ffmpeg

WORKDIR /app

# FrontEnd Source code
COPY --from=FrontBuilder /app/dist/src/pages ./pages
COPY --from=FrontBuilder /app/dist/assets ./assets
COPY frontend/public ./public

# BackEnd Source code
COPY --from=BackendBuilder /app/bin/* ./

# Object storage 
ENV MINIO_ACCESSKEY=
ENV MINIO_SECRET_ACCESS_KEY=
ENV MINIO_END_POINT=

# Database
# ENV POSTGRES_USER=
ENV POSTGRES_PASSWORD=
ENV POSTGRES_PORT=

# Application
ENV PORT=19622

EXPOSE 19622

ENTRYPOINT [ "./phalanx" ]