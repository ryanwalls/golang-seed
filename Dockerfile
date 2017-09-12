FROM golang:1.9

RUN apt-get update && apt-get install -y git

ENV APPLICATION_NAME=golang-seed

# Install glide
WORKDIR $GOPATH/src/github.com/Masterminds
RUN git clone https://github.com/Masterminds/glide.git
RUN cd glide && git checkout v0.12.3 && make bootstrap-dist && make install

RUN mkdir -p $GOPATH/src/github.com/3dsim/$APPLICATION_NAME

COPY . $GOPATH/src/github.com/3dsim/$APPLICATION_NAME

WORKDIR $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
RUN glide install --strip-vendor
RUN go test $(glide novendor)
# Must run after glide installs 3rd party (vendor) dependencies.  Installs binary on path.
RUN go install

ENTRYPOINT ["golang-seed"]
