FROM golang:buster 

ARG INOTIFY_TOOLS_VERSION=3.20.11.0

RUN apt-get -y update \
    && apt-get upgrade -y \
    && apt-get install -y make libtool g++ \
    # Install inotify-tools
    && wget -O- https://github.com/inotify-tools/inotify-tools/archive/refs/tags/${INOTIFY_TOOLS_VERSION}.tar.gz | tar xzC /tmp \
    && cd /tmp/inotify-tools-${INOTIFY_TOOLS_VERSION} \
    && ./autogen.sh && ./configure --prefix=/usr && make && su -c 'make install' \
    && rm -rf /tmp/inotify-tools-${INOTIFY_TOOLS_VERSION} \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* 

VOLUME /code
WORKDIR /code

CMD ["bash", "scripts/hot-reload-run.sh"]