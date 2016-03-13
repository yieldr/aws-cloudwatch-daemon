# aws-cloudwatch-daemon

Send disk and memory usage statistics to AWS CloudWatch.

## Usage

	aws-cloudwatch-daemon -memory-usage -disk-usage -disk-path="/"

Using the docker image:

	docker run -v /usr/share/ca-certificates:/etc/ssl/certs:ro yieldr/aws-cloudwatch-daemon

A common use case is to have the daemon run on a schedule such as cron or systemd timer. You can find an example in the `dist/systemd` folder.

## Credits

This repository is heavily based on [Allen Chen(a3linux)](a3linux)'s implementation with several parts revised.
