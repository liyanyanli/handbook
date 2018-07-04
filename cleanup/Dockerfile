FROM debian:stable-slim

ADD cleanUp.sh .
RUN chmod +x cleanUp.sh

CMD ["./cleanUp.sh","/data","30","messages-*","1"]
