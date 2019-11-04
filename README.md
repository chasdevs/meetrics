# Meetrics

Basic Google Calendar meeting data tracking within a Google organization. This is the codebase behind [this blog post](https://engineering.videoblocks.com/analyzing-meeting-metrics-using-the-google-calendar-api-3c76c9f8ffea) which describes how we analyzed meeting data at [Storyblocks](https://www.storyblocks.com).

![Meeting Data Graph](https://miro.medium.com/max/3868/1*K1mHm1dBwsQGCvs9A1xs6A.png)

## Requirements

- A Google Cloud Service Account 
    - Needs the following scopes: _calendar.readonly_, _admin.directory.user.readonly_
    - Need a valid `service_account_key.json` file.
    - May require admin access to your organization's G Suite account.

## Setup

### Preparing Google Calendar

A "service account" is used to communicate with the Google Calendar API. Their [documentation](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) best describes how to do this. The basic steps I took were the following:

1. Create a project in the Google APIs [dashboard](https://console.developers.google.com/apis/dashboard) to hold the service account.
1. Create the service account in the [IAM & admin](https://console.developers.google.com/iam-admin/serviceaccounts) section of your project.
1. Delegate domain-wide authority to the service account.
1. Log into the G Suite [Admin](http://admin.google.com/) section (as an admin user) to authorize the service account for API scopes. For listing calendar information, you need the _calendar.readonly_ scope: `https://www.googleapis.com/auth/calendar.readonly`. For listing users in the domain, you need the _admin.directory.user.readonly_ scope: `https://www.googleapis.com/auth/admin.directory.user.readonly`.
1. Go to the [credentials page](https://console.developers.google.com/apis/credentials) in the Google APIs dashboard for your project, click **Create credentials > Service account key**, select your service account and JSON as the key type, and click **Create**. This will download your credentials JSON file containing your private key which will be used in Google API SDKs.

To connect, place your `service_account_key.json` file inside this application's _config/_ directory (it is git-ignored).

Lastly, you will need to choose a "subject" to query Google Calendar, which is the email address of a user in your organization who can see the calendars of your users (I used mine). This is set via the `google.subject` config variable.

### Configuration

You MUST put a valid Google _service_account_key.json_ file in the _config/_ directory.

The following configs are set via environment variables or _config/production.yml_ file.

- GOOGLE_DOMAIN
    - (required) The domain of the Google account. 
- GOOGLE_SUBJECT 
    - (required) The email of the account to query Google's Calendar API.

## Development

```
# Start mysql using docker-compose and initialize local database.
make init
```

### Verify Configuration

You can verify that the code is able to hit the Google APIs by running the tests for the `apis` package.
```
go test ./pkg/apis
```

### Docker

Build and push the docker image. Jenkins pulls the docker image each night and runs the script to populate the database with meeting information.

```bash
docker build -t chasdevs/meetrics .
```

You can run locally from the docker image:

```bash
docker run --rm -e DB_HOST="docker.for.mac.localhost" chasdevs/meetrics
```

Or against prod:
```bash
docker run --rm -e ENV=production chasdevs/meetrics
```