FROM alpine

LABEL maintainer="Sherlock Holo sherlockya@gmail.com"

ADD coredns /

RUN chmod +x /coredns

CMD /coredns -conf /Corefile
