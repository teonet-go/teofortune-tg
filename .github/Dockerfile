
# Copyright 2022 Kirill Scherba <kirill@scherba.ru>.  All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Teonet teofortune-tg server application docker file
#
# Docker build (included private repositories):
# --------------------------------------------
#
#   docker build -t teofortune-tg -f ./.github/Dockerfile .
#
# Docker test run:
# ---------------
#
#   docker run --network host -it --rm teofortune-tg teofortune-tg
#
# Publish to github:
# -----------------
#
#   docker login ghcr.io -u USERNAME -p TOKEN_FOR_PACKAGES
#   #
#   docker tag teocfg ghcr.io/teonet-go/teofortune-tg:0.0.0
#   docker push ghcr.io/teonet-go/teofortune-tg:0.0.0
#   #
#   docker tag teofortune-tg ghcr.io/teonet-go/teofortune-tg:0.0.0
#   docker push ghcr.io/teonet-go/teofortune-tg:0.0.0
#
# Run from gihub repo:
# -------------------
#
#   docker run --network=host -it --restart=always --name teofortune-tg \
#     -v $HOME/.config/teonet/teofortune-tg:/root/.config/teonet/teofortune-tg \
#     ghcr.io/teonet-go/teofortune-tg:0.0.1 teofortune-tg -loglevel=debug
#

# Build stage
#
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .

# Add from arguments
ARG version
ENV version=$version

RUN go get
RUN go install ./

# Final stage
#
FROM alpine:latest
RUN apk --no-cache add ca-certificates fortune
COPY --from=builder /go/bin/* /usr/local/bin/
CMD "teofortune-tg"
LABEL Name=teofortune-tg Version=$(version)
