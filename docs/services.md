# Service

## Table of Contents

- [Service](#service)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
    - [Must Service](#must-service)
      - [1. Google (Gmail / Drive / Calendar )](#1-google-gmail--drive--calendar-)
      - [2. Meta ( Facebook / instagram / Whatsapp )](#2-meta--facebook--instagram--whatsapp-)
      - [3. Spotify](#3-spotify)
      - [4. Github](#4-github)
      - [5. OpenWeatherMap](#5-openweathermap)
      - [6. Timer](#6-timer)
    - [Maybe Service](#maybe-service)
      - [1. Dropbox](#1-dropbox)
      - [2. Youtube](#2-youtube)
      - [3. RSS flux](#3-rss-flux)
      - [4. X](#4-x)
      - [5. Microsoft (OneDrive / Outlook / Agenda )](#5-microsoft-onedrive--outlook--agenda-)
      - [6. Deezer](#6-deezer)

## Main Document

[main documentation](../README.md)

### Must Service

if we implement a service we will at least do:

- 1 action
- 1 reaction

List the service we will serve

#### 1. Google (Gmail / Drive / Calendar )

- Gmail:
  - [Gmail API](https://developers.google.com/gmail)
  - Action:
    - Receipt of a message
    - Receipt of a message from a user X
    - Receipt of a message whose title contains the word X
  - Reaction:
    - The user sends a message M to a recipient D
      - [API Guide](https://developers.google.com/gmail/api/guides/sending)
      - [API Reference](https://developers.google.com/gmail/api/reference/rest/v1/users.messages/send)
- Drive:
  - [Drive API](https://developers.google.com/drive)
  - Action:
    - A new file is present
    - A new file is present in the X directory
    - A user shares a file
  - Reaction:
    - The user adds the file F in the drive
      - [API Reference](https://developers.google.com/drive/api/reference/rest/v3/files/create)
    - The user shares the F file with another U user
- Calendar:
  - [Calendar API](https://developers.google.com/calendar)
  - Action:
    - Start of an event
    - End of an event
    - Add of an event
  - Reaction:
    - Add event
      - [insert API Reference](https://developers.google.com/calendar/api/v3/reference/events/insert)

#### 2. Meta ( Facebook / instagram / Whatsapp )

- Facebook:

  - Action:
    - A new message is posted in group G
    - A new message containing a #hashtag is posted
    - A new private message is received by the user
    - One of the user’s messages gets a like
    - The user gains a Follower
  - Reaction:
    - The user posts a message in group G
    - The user is a new person P

- Instagram:

  - Action:
    - A new message is posted in group G
    - A new message containing a #hashtag is posted
    - A new private message is received by the user
    - One of the user’s messages gets a like
    - The user gains a Follower
  - Reaction:
    - The user posts a message in group G
    - The user is a new person P

- Whatsapp:
  - [Whatsapp API](https://developers.facebook.com/docs/whatsapp) LIMITED
  - Action:
    - A new message is posted in group G
    - A new private message is received by the user
  - Reaction:
    - The user posts a message in group G
    - The user is a new person P

#### 3. Spotify

- [Spotify API](https://developer.spotify.com/documentation/web-api)
- Action:
  - User current-playback music
    - [get-information-about-the-users-current-playback API](https://developer.spotify.com/documentation/web-api/reference/get-information-about-the-users-current-playback)
- Reaction:
  - Start music
    - [start-a-users-playback API](https://developer.spotify.com/documentation/web-api/reference/start-a-users-playback)
  - Stop music
    - [pause-a-users-playback API](https://developer.spotify.com/documentation/web-api/reference/pause-a-users-playback)
  - Give information on the music/artist
  - Change volume of the music

#### 4. Github

- [Github API](https://docs.github.com/en/rest)
- **Action:**
  - **Create a new repository R**
  - Create a new branch in the repo R
  - Commit a file in the repo R in the branch B
- **Reaction:**
  - **Create a new repository R**
  - Create a new branch in the repo R
    - [get-a-branch API](https://docs.github.com/en/rest/branches/branches?apiVersion=2022-11-28#get-a-branch)
  - Commit a file in the repo R
    - [create-or-update-file-contents API](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-or-update-file-contents)

#### 5. OpenWeatherMap

- [OpenWeatherMap API](http://api.openweathermap.org/data/2.5/weather)
- [OpenWeatherMap API](https://openweathermap.org/api) (there are a lot of API)
  - Can put params like “q: cityName, appid: apiKey, units: 'metric'”
- **Action:**
  - Temperature lower or greater than T
  - Weather forecast
- **Reaction:**
  - Show current weather
    - [weather_overview API](https://openweathermap.org/api/one-call-3#weather_overview)

#### 6. Timer

- [WorldTimeApi API](https://worldtimeapi.org/pages/examples)
- Action:
  - Specific Time T
- Reaction:
  - Show time

### Maybe Service

List the service we MAYBE will serve

#### 1. Dropbox

- [Dropbox API](https://www.dropbox.com/developers/documentation/http/documentation)
- Action:
  - New file
- Reaction:
  - Add File T
    - [files-upload API](https://www.dropbox.com/developers/documentation/http/documentation#files-upload)
  - Download file
    - [files-download API](https://www.dropbox.com/developers/documentation/http/documentation#files-download)

#### 2. Youtube

- [Youtube API](https://developers.google.com/youtube/v3/docs)
- Action:
  - New video in subscribe
- Reaction:
  - Get info of a video
    - [videos API](https://developers.google.com/youtube/v3/docs/videos)

#### 3. RSS flux

- Action:
  - A new article is available
  - An article is added to their favorites by the user
- Reaction:
  - get the latest article

#### 4. X

- [X API](https://developer.x.com/en/docs)
- Action:
  - A new message is posted in group G
  - A new message containing a #hashtag is posted
  - A new private message is received by the user
  - One of the user’s messages gets a like
  - The user gains a Follower
- Reaction:
  - The user posts a message in group G
  - The user is a new person P
  - Get user info
    - [api-reference API](https://developer.x.com/en/docs/x-api/users/lookup/api-reference)

#### 5. Microsoft (OneDrive / Outlook / Agenda )

- API: haven’t found it

#### 6. Deezer

- API: can’t access it because I didn’t win the CAPTCHA
- Action:
  - User start music
  - User stop music
- Reaction:
  - Start music
  - Stop music
  - Give informations on the music/artist
  - Change volume of the music
