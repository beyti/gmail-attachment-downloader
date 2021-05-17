gmail-downloader is a simple command line utility to download attachments from a Gmail mailbox
related to a label.


## Build

1. Install the dependencies

```
go get -u -v google.golang.org/api/gmail/v1
go get -u -v golang.org/x/oauth2/google
```

2. Build the project

```
go build
```

## Requirements

You need `credentials.json` to use the application.
Visit the [Gmail API page](https://developers.google.com/gmail/api/quickstart/go)
and turn on le Gmail API to have it.

You have to put `credentials.json` in the root folder of the project

## Usage

```
./gmail-downloader -h
Usage of ./gmail-downloader:
  -label string
    	label to process
  -no-overwrite
    	not overwrite existent files
  -out string
    	output folder (default "./attachments")
```
## License

This software is under GPLv3 license.

Every contribution will go under the same license.

## Credits

This utility takes inspiration from [attachment-downloader](https://github.com/munir131/attachment-downloader)
of [munir131](https://github.com/munir131) to understand which APIs of Gmail to use to reach the goal.
