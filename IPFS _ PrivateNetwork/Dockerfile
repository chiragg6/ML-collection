FROM ubuntu
MAINTAINER Chirag Gupta "chiragg6@gmail.com"

FROM golang:1.10.1 AS builder
RUN go version


RUN apt-get update
RUN apt -y install wget && apt -y install sudo

RUN apt install git
RUN wget https://dist.ipfs.io/go-ipfs/v0.4.4/go-ipfs_v0.4.4_linux-amd64.tar.gz
RUN tar xvfz go-ipfs_v0.4.4_linux-amd64.tar.gz
RUN cd go-ipfs && ./install.sh


RUN ipfs init
RUN ipfs bootstrap rm --all
RUN ipfs config Addresses.Gateway /ip4/0.0.0.0/tcp/8080
RUN ipfs config Addresses.API /ip4/0.0.0.0/tcp/5001
RUN export LIBP2P_FORCE_PNET=1
RUN go get github.com/Kubuxu/go-ipfs-swarm-key-gen/ipfs-swarm-key-gen
RUN ipfs-swarm-key-gen > ~/.ipfs/swarm.key
RUN ipfs daemon > ipfs.log &






