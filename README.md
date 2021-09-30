gmail-downloader is a simple command line utility to download attachments from a Gmail mailbox
related to a label.


## Build

```
go build
```

## Requirements

You need `credentials.json` to use the application.
Visit the [Gmail API page](https://developers.google.com/gmail/api/quickstart/go)
and turn on le Gmail API to have it.

Save `credentials.json` in the root folder of the project or where you want (in this case remember
to set the `GDOWN_HOME` env var).

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

Use the env var `GDOWN_HOME` to set the home folder for the application (where `credentials.json` is saved).

## License

This software is under GPLv3 license.

Every contribution will go under the same license.

## Credits

This utility takes inspiration from [attachment-downloader](https://github.com/munir131/attachment-downloader)
of [munir131](https://github.com/munir131) to understand which APIs of Gmail to use to reach the goal.
