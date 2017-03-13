FROM golang:onbuild
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN make clean ; make gherkinize;
CMD ["./app"]