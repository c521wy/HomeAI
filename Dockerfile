FROM alpine AS build-stage

COPY . /usr/local/HomeAI/

WORKDIR /usr/local/HomeAI/

RUN \
set -x && \
sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
apk add go git && \
./build.sh && \
mv ./bin/HomeAI /usr/local/bin/HomeAI

FROM alpine

COPY --from=build-stage /usr/local/bin/HomeAI /usr/local/bin/HomeAI

CMD /usr/local/bin/HomeAI /etc/HomeAI.json
