FROM scratch
EXPOSE 8080
COPY am-go /
COPY content/ content/
ENTRYPOINT ["/am-go"]
