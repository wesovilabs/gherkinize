FROM golang:onbuild
RUN mkdir -p /app ; mkdir -p /var/gherkinize/scenarios ; mkdir -p /var/gherkinize/config
VOLUME  ["/var/gherkinize/scenarios", "/var/gherkinize/config"]
WORKDIR /app
ADD . /app
RUN make clean ; make gherkinize
CMD ["dist/gherkinize","-i", "/var/gherkinize/scenarios", "-c", "/var/gherkinize/config/gherkin.toml", "validate"]