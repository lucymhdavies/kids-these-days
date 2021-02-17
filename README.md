# Kids These Days

Code based on https://jonaylor89.hashnode.dev/creating-a-twitter-bot-in-go

## Credentials

You need a Twitter dev account.

Sign up here:
https://developer.twitter.com/en/apply-for-access

Once you have that set up (could take a few days), follow tutorial here to generate creds:
https://dev.to/stratospher/many-bot-accounts-using-1-twitter-developer-account-17ff
https://gist.github.com/moonmilk/035917e668872013c1bd

(Remember to edit the values at the top of `authorize.js`)


Store your creds in a .env file. The first two are from your Twitter app. The second two are generated from the `authorize.js`

```
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_SECRET=
```