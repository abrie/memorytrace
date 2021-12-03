FROM golang:1.17 as backend
WORKDIR /go/src
COPY backend/src/ .
RUN CGO_ENABLED=1 go install -ldflags '-extldflags "-static"' -tags sqlite_omit_load_extension,netgo backend/cmd/server

FROM node:latest as frontend
ARG GITHUB_RUN_NUMBER
WORKDIR /usr/src/app
COPY frontend/ ./
RUN yarn
RUN REACT_APP_CI_RUN_NUMBER=$GITHUB_RUN_NUMBER yarn build

FROM gcr.io/distroless/base
ARG BUILD_TAG
COPY --from=backend /go/bin/server /server
COPY --from=frontend /usr/src/app/build /frontend
ENV BUILD_TAG=$BUILD_TAG

ENTRYPOINT ["/server", "-store", "/data", "-static", "/frontend", "-addr", ":80"]
