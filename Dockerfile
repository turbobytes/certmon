FROM scratch

ADD bin/certmon /bin/
ADD bin/ca-certificates.crt /etc/ssl/certs/
ADD assets /assets

CMD ["certmon", "-config", "/config.yaml", "-listen", ":8082", "-ui", "/assets/index.html"]
