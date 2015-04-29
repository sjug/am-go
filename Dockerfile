FROM scratch
EXPOSE 8080
COPY am-go /
ENTRYPOINT ["/am-go"]
