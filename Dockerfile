FROM ubuntu:22.04 AS build-stage

USER root
ARG DEBIAN_FRONTEND=noninteractive
ARG GID=1000
ARG UID=1000
ENV LANG=C.UTF-8 LC_ALL=C.UTF-8
ENV SHELL="/bin/bash"
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN apt update -y > /dev/null && \
    apt install -y software-properties-common && \
    apt install -y bzip2 ca-certificates curl expect ftp git gnupg && \
    apt install -y pkg-config python3 python3-dev python3-pip python3-venv && \
    apt install -y sudo unzip upx vim wget xz-utils zip
RUN apt autoremove --purge -y > /dev/null && \
    apt autoclean -y > /dev/null && \
    rm -rf /var/lib/apt/lists/* && \
    rm -rf /var/log/* && \
    rm -rf /tmp/* \
RUN echo "alias pip=pip3" | tee --append /etc/bash.bashrc && \
    echo "alias python=python3" | tee --append /etc/bash.bashrc && \
    echo "StrictHostKeyChecking no" | tee --append /etc/ssh/ssh_config && \
    echo "craftslab ALL=(ALL) NOPASSWD: ALL" | tee --append /etc/sudoers && \
    echo "dash dash/sh boolean false" | debconf-set-selections && \
    DEBIAN_FRONTEND=noninteractive dpkg-reconfigure dash && \
    groupadd -g $GID craftslab && \
    useradd -d /home/craftslab -ms /bin/bash -g craftslab -u $UID craftslab

USER craftslab
WORKDIR /home/craftslab
ENV PATH==$PATH:/usr/local/go/bin
RUN curl -L https://go.dev/dl/go1.23.4.linux-amd64.tar.gz -o go.tar.gz && \
    sudo rm -rf /usr/local/go && \
    sudo tar -C /usr/local -xzf go.tar.gz && \
    rm -f go.tar.gz
RUN git clone https://github.com/ai-flowx/flowx.git flowx --depth=1 && \
    pushd flowx && \
    make build && \
    popd

FROM ubuntu:22.04 AS production-stage
WORKDIR /
RUN apt update -y && \
    apt install -y pkg-config python3 python3-dev
COPY --from=build-stage /home/craftslab/flowx/bin/* /
ENTRYPOINT ["/flowx"]
CMD ["--help"]
