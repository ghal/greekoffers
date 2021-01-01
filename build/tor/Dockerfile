FROM alpine
RUN apk update && apk add tor
RUN cp /etc/tor/torrc.sample /etc/tor/torrc && echo "SocksPort 0.0.0.0:9050" > /etc/tor/torrc
EXPOSE 9050
CMD ["/usr/bin/tor"]