FROM scratch
COPY customer-app /

EXPOSE 8080
ENTRYPOINT ["/customer-app"]
