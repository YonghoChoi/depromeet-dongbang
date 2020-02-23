FROM debian:jessie

RUN apt-get update -y && apt-get install -y vim dos2unix

RUN mkdir /opt/dongbang
COPY ./output/linux/dongbang-api /opt/dongbang/dongbang-api
ADD ./cmd/api/conf /opt/dongbang/
RUN chmod +x /opt/dongbang/dongbang-api

WORKDIR /opt/dongbang
CMD /opt/dongbang/dongbang-api