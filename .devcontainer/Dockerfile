FROM mcr.microsoft.com/devcontainers/go:0-1.20

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root
ARG SERF_VERSION="0.8.2"
RUN sudo apt-get update \ 
    && sudo apt-get install -y --no-install-recommends protobuf-compiler libprotobuf-dev \
    && sudo rm -rf /var/lib/apt/lists/* \
    && sudo curl -L -o serf.zip https://releases.hashicorp.com/serf/${SERF_VERSION}/serf_${SERF_VERSION}_linux_amd64.zip \
    && sudo unzip serf.zip -d /usr/local/serf \
    && sudo rm -f serf.zip
ENV PATH="$PATH:/usr/local/serf"

# go getは、Googleのサーバーへの接続がブロックされたため
# RUN go install google.golang.org/grpc@v1.53.0
RUN git clone https://github.com/grpc/grpc-go.git /go/src/google.golang.org/grpc

RUN go install github.com/sqs/goreturns@latest

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
RUN go install google.golang.org/protobuf/...@v1.28.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
