FROM alpine:3.24
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

# The Terraform provider is compiled into the provider binary and runs
# in-process (no-fork), so the image has no Terraform CLI or provider binary.

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
