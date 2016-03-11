FROM scratch

ADD bin/aws-cloudwatch-daemon /

CMD ["/aws-cloudwatch-daemon"]
