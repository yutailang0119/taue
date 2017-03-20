# Taue is reporting developer activity (GitHub, GitLab, etc...) to Slack

* Go 1.8~
* It is assumed that work on Heroku

## Usage
1. Add user information to `taue/resources/Users.json`
    * id int (required): user id
    * name string (required): user name (sample Twitter)
    * slackName string (required): Slack user name
    * githubName string (required): GitHub user name (https://github.com/xxxxxxx)
    * githubToken string (optional): [GitHub Personal access token](https://github.com/settings/tokens)
    * gitlabId int (optional): GitLab user id
    * gitlabToken string (optional): [GitLab Personal access token](https://gitlab.com/profile/personal_access_tokens)

```javascript
[
  {
    "id": 0,
    "name": "yutailang0119",
    "slackName": "yutaro",
    "githubName": "Yu-taro",
    "githubToken": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "gitlabId": 01234567,
    "gitlabToken": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  },
  {
    "name": "d_date",
    "slackName": "d_date",
    "githubName": "d-date",
  }
]
```

2. Add [Slack WebHook URL](https://api.slack.com/incoming-webhooks)Endpoint (https://hooks.slack.com/services/xxxxxxxxxxxxxxxx) to env
**Example for Heroku**

```bash
$ heroku config:set SLACK_WEBHOOK_ENDPOINT="xxxxxxxxxxxxxxxx" --app "app_name"
```

3. Set a scheduler

## License
[taue](https://github.com/Yu-taro/taue) is released under the [Apache License 2.0](LICENSE).


