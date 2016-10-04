FROM       alpine:3.3
MAINTAINER Kelsey Hightower <kelsey.hightower@gmail.com>
ADD        journal-2-logentries journal-2-logentries
ENTRYPOINT ["/journal-2-logentries"]
