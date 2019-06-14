#FROM golang:1.10-stretch as build
#
#WORKDIR /go/src/danlu-gpu-pci-device-plugin
#COPY . .
#
#RUN export CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' && \
#    go install -ldflags="-s -w" -v danlu-gpu-pci-device-plugin


FROM golang:1.10-stretch

#ENV NVIDIA_VISIBLE_DEVICES=all
#ENV NVIDIA_DRIVER_CAPABILITIES=utility

COPY ./danlu-gpu-pci-device-plugin /usr/bin/danlu-gpu-pci-device-plugin
RUN chmod +x /usr/bin/danlu-gpu-pci-device-plugin

CMD ["danlu-gpu-pci-device-plugin"]
