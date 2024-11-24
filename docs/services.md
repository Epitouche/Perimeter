# Service

## Table of Contents

- [Service](#service)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
    - [Must Service](#must-service)
      - [1. Gmail (Google service)](#1-gmail-google-service)
      - [2. Google Drive (Google service)](#2-google-drive-google-service)
      - [3. Google Calendar (Google service)](#3-google-calendar-google-service)
      - [4. Facebook (Meta service)](#4-facebook-meta-service)
      - [5. Instagram (Meta service)](#5-instagram-meta-service)
      - [5. Whatsapp (Meta service)](#5-whatsapp-meta-service)
      - [6. Spotify (Spotify Service)](#6-spotify-spotify-service)
      - [7. Github (Github Service)](#7-github-github-service)
      - [8. OpenWeatherMap (OpenWeatherMap Service)](#8-openweathermap-openweathermap-service)
      - [9. Timer (Timer Service)](#9-timer-timer-service)
    - [Maybe Service](#maybe-service)
      - [10. Dropbox (Dropbox Service)](#10-dropbox-dropbox-service)
      - [11. Youtube (Google service)](#11-youtube-google-service)
      - [12. RSS flux (RSS Service)](#12-rss-flux-rss-service)
      - [13. X (X Service)](#13-x-x-service)
      - [14. Outlook (Microsoft Service)](#14-outlook-microsoft-service)
      - [15. OneDrive (Microsoft Service)](#15-onedrive-microsoft-service)
      - [16. Outlook calendar (Microsoft Service)](#16-outlook-calendar-microsoft-service)
      - [17. Deezer (Deezer Service)](#17-deezer-deezer-service)

## Main Document

[main documentation](../README.md)

### Must Service

if we implement a service we will at least do:

- 1 action
- 1 reaction

List the service we will serve

#### 1. Gmail (Google service)

[Gmail API](https://developers.google.com/gmail)
Action:

- Receipt of a message
- Receipt of a message from a user X
- Receipt of a message whose title contains the word X

Reaction:

- The user sends a message M to a recipient D
  - [API Guide](https://developers.google.com/gmail/api/guides/sending)
  - [API Reference](https://developers.google.com/gmail/api/reference/rest/v1/users.messages/send)

#### 2. Google Drive (Google service)

[Google Drive API](https://developers.google.com/drive)
Action:

- A new file is present
- A new file is present in the X directory
- A user shares a file

Reaction:

- The user adds the file F in the drive
  - [API Reference](https://developers.google.com/drive/api/reference/rest/v3/files/create)
- The user shares the F file with another U user

#### 3. Google Calendar (Google service)

[Google Calendar API](https://developers.google.com/calendar)
Action:

- Start of an event
- End of an event
- Add of an event

Reaction:

- Add event
  - [insert API Reference](https://developers.google.com/calendar/api/v3/reference/events/insert)

#### 4. Facebook (Meta service)

[Facebook API](https://developers.facebook.com/docs/facebook-login/)

Action:

- A new message is posted in group G
- A new message containing a #hashtag is posted
- A new private message is received by the user
- One of the user’s messages gets a like
- The user gains a Follower

Reaction:

- The user posts a message in group G
- The user is a new person P

#### 5. Instagram (Meta service)

[Instagram API](https://developers.facebook.com/docs/instagram-platform)

Action:

- A new message is posted in group G
- A new message containing a #hashtag is posted
- A new private message is received by the user
- One of the user’s messages gets a like
- The user gains a Follower

Reaction:

- The user posts a message in group G
- The user is a new person P

#### 5. Whatsapp (Meta service)

[Whatsapp API](https://developers.facebook.com/docs/whatsapp) LIMITED

Action:

- A new message is posted in group G
- A new private message is received by the user

Reaction:

- The user posts a message in group G
- The user is a new person P

#### 6. Spotify (Spotify Service)

[Spotify API](https://developer.spotify.com/documentation/web-api)
Action:

- User current-playback music
  - [get-information-about-the-users-current-playback API](https://developer.spotify.com/documentation/web-api/reference/get-information-about-the-users-current-playback)

Reaction:

- Start music
  - [start-a-users-playback API](https://developer.spotify.com/documentation/web-api/reference/start-a-users-playback)
- Stop music
  - [pause-a-users-playback API](https://developer.spotify.com/documentation/web-api/reference/pause-a-users-playback)
- Give information on the music/artist
- Change volume of the music

#### 7. Github (Github Service)

[Github API](https://docs.github.com/en/rest)
Action:

- **Create a new repository R**
- Create a new branch in the repo R
- Commit a file in the repo R in the branch B

Reaction:

- **Create a new repository R**
- Create a new branch in the repo R
  - [get-a-branch API](https://docs.github.com/en/rest/branches/branches?apiVersion=2022-11-28#get-a-branch)
- Commit a file in the repo R
  - [create-or-update-file-contents API](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-or-update-file-contents)

#### 8. OpenWeatherMap (OpenWeatherMap Service)

[OpenWeatherMap API](http://api.openweathermap.org/data/2.5/weather)
[OpenWeatherMap API](https://openweathermap.org/api) (there are a lot of API)

- Can put params like “q: cityName, appid: apiKey, units: 'metric'”

Action:

- Temperature lower or greater than T
- Weather forecast

Reaction:

- Show current weather
  - [weather_overview API](https://openweathermap.org/api/one-call-3#weather_overview)

#### 9. Timer (Timer Service)

[WorldTimeApi API](https://worldtimeapi.org/pages/examples)
Action:

- Specific Time T

Reaction:

- Show time

### Maybe Service

List the service we MAYBE will serve

#### 10. Dropbox (Dropbox Service)

[Dropbox API](https://www.dropbox.com/developers/documentation/http/documentation)
Action:

- New file

Reaction:

- Add File T
  - [files-upload API](https://www.dropbox.com/developers/documentation/http/documentation#files-upload)
- Download file
  - [files-download API](https://www.dropbox.com/developers/documentation/http/documentation#files-download)

#### 11. Youtube (Google service)

[Youtube API](https://developers.google.com/youtube/v3/docs)
Action:

- New video in subscribe

Reaction:

- Get info of a video
  - [videos API](https://developers.google.com/youtube/v3/docs/videos)

#### 12. RSS flux (RSS Service)

Action:

- A new article is available
- An article is added to their favorites by the user

Reaction:

- get the latest article

#### 13. X (X Service)

[X API](https://developer.x.com/en/docs)
Action:

- A new message is posted in group G
- A new message containing a #hashtag is posted
- A new private message is received by the user
- One of the user’s messages gets a like
- The user gains a Follower

Reaction:

- The user posts a message in group G
- The user is a new person P
- Get user info
  - [api-reference API](https://developer.x.com/en/docs/x-api/users/lookup/api-reference)

#### 14. Outlook (Microsoft Service)

[Outlook API](https://learn.microsoft.com/en-us/outlook/rest/reference)

#### 15. OneDrive (Microsoft Service)

[OneDrive API](https://learn.microsoft.com/en-us/onedrive/developer/rest-api/getting-started/?view=odsp-graph-online)

#### 16. Outlook calendar (Microsoft Service)

[Outlook calendar API](https://learn.microsoft.com/en-us/graph/outlook-calendar-concept-overview)

#### 17. Deezer (Deezer Service)

[API](https://developers.deezer.com/api)
Action:

- User start music
- User stop music

Reaction:

- Start music
- Stop music
- Give informations on the music/artist
- Change volume of the music
