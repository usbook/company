FROM scratch
WORKDIR $GOPATH/src/company
COPY . $GOPATH/src/company
EXPOSE 8020
CMD ["./company"]
