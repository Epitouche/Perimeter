# Services

## Table of Contents

- [Services](#services)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Must Services](#must-services)
    - [1. Gmail (Google Service)](#1-gmail-google-service)
    - [2. Google Drive (Google Service)](#2-google-drive-google-service)
    - [3. Google Calendar (Google Service)](#3-google-calendar-google-service)
    - [4. Facebook (Meta Service)](#4-facebook-meta-service)
    - [5. Instagram (Meta Service)](#5-instagram-meta-service)
    - [6. WhatsApp (Meta Service)](#6-whatsapp-meta-service)
    - [7. Spotify (Spotify Service)](#7-spotify-spotify-service)
    - [8. GitHub (GitHub Service)](#8-github-github-service)
    - [9. OpenWeatherMap (OpenWeatherMap Service)](#9-openweathermap-openweathermap-service)
    - [10. Timer (Timer Service)](#10-timer-timer-service)
  - [Optional Services](#optional-services)
    - [11. Dropbox (Dropbox Service)](#11-dropbox-dropbox-service)
    - [12. YouTube (Google Service)](#12-youtube-google-service)
    - [13. RSS Feed (RSS Service)](#13-rss-feed-rss-service)
    - [14. X (X Service)](#14-x-x-service)
    - [15. Outlook (Microsoft Service)](#15-outlook-microsoft-service)
    - [16. OneDrive (Microsoft Service)](#16-onedrive-microsoft-service)
    - [17. Outlook Calendar (Microsoft Service)](#17-outlook-calendar-microsoft-service)
    - [18. Deezer (Deezer Service)](#18-deezer-deezer-service)

---

## Main Document

Refer to the [main documentation](../README.md).

---

## Must Services

For each service we implement, the minimum requirements are:

- **One action**
- **One reaction**

### 1. Gmail (Google Service)

[Gmail API Documentation](https://developers.google.com/gmail)

**Actions:**

- [ ] Receive a message
- [ ] Receive a message from user X
- [ ] Receive a message with a subject containing the word X

**Reactions:**

- [ ] Send message M to recipient D
  - [API Guide](https://developers.google.com/gmail/api/guides/sending)
  - [API Reference](https://developers.google.com/gmail/api/reference/rest/v1/users.messages/send)

---

### 2. Google Drive (Google Service)

[Google Drive API Documentation](https://developers.google.com/drive)

**Actions:**

- [ ] Detect a new file
- [ ] Detect a new file in directory X
- [ ] Detect a file shared by a user

**Reactions:**

- [ ] Add file F to the drive
  - [API Reference](https://developers.google.com/drive/api/reference/rest/v3/files/create)
- [ ] Share file F with user U

---

### 3. Google Calendar (Google Service)

[Google Calendar API Documentation](https://developers.google.com/calendar)

**Actions:**

- [ ] Event starts
- [ ] Event ends
- [ ] Add an event

**Reactions:**

- [ ] Add an event
  - [API Reference](https://developers.google.com/calendar/api/v3/reference/events/insert)

---

### 4. Facebook (Meta Service)

[Facebook API Documentation](https://developers.facebook.com/docs/facebook-login/)

**Actions:**

- [ ] A new message is posted in group G
- [ ] A new message containing #hashtag
- [ ] Receive a private message
- [ ] User's message gets a like
- [ ] User gains a follower

**Reactions:**

- [ ] Post a message in group G
- [ ] Follow a new user P

---

### 5. Instagram (Meta Service)

[Instagram API Documentation](https://developers.facebook.com/docs/instagram-platform)

**Actions:**

- [ ] A new message is posted in group G
- [ ] A new message containing #hashtag
- [ ] Receive a private message
- [ ] User's message gets a like
- [ ] User gains a follower

**Reactions:**

- [ ] Post a message in group G
- [ ] Follow a new user P

---

### 6. WhatsApp (Meta Service)

[WhatsApp API Documentation](https://developers.facebook.com/docs/whatsapp)
_(Limited functionality)_

**Actions:**

- [ ] New message in group G
- [ ] Receive a private message

**Reactions:**

- [ ] Post a message in group G
- [ ] Follow a new user P

---

### 7. Spotify (Spotify Service)

[Spotify API Documentation](https://developer.spotify.com/documentation/web-api)

**Actions:**

- [ ] Get user’s current playback information
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/get-information-about-the-users-current-playback)

**Reactions:**

- [ ] Start playback
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/start-a-users-playback)
- [ ] Pause playback
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/pause-a-users-playback)
- [ ] Change volume
- [ ] Retrieve music/artist details

---

### 8. GitHub (GitHub Service)

[GitHub API Documentation](https://docs.github.com/en/rest)

**Actions:**

- [ ] Create a new repository
- [ ] Create a branch in repository R
- [ ] Commit a file in branch B of repository R

**Reactions:**

- [ ] Create a branch in repository R
  - [Branch API](https://docs.github.com/en/rest/branches/branches?apiVersion=2022-11-28#get-a-branch)
- [ ] Commit a file in branch B of repository R
  - [File Contents API](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-or-update-file-contents)

---

### 9. OpenWeatherMap (OpenWeatherMap Service)

[OpenWeatherMap API Documentation](https://openweathermap.org/api)

**Actions:**

- [ ] Temperature below or above T
- [ ] Fetch weather forecast

**Reactions:**

- [ ] Display current weather
  - [API Reference](https://openweathermap.org/api/one-call-3#weather_overview)

---

### 10. Timer (Timer Service)

[WorldTimeAPI Documentation](https://worldtimeapi.org/pages/examples)

**Actions:**

- [ ] Trigger at specific time T

**Reactions:**

- [ ] Get current time

---

## Optional Services

### 11. Dropbox (Dropbox Service)

[Dropbox API Documentation](https://www.dropbox.com/developers/documentation/http/documentation)

**Actions:**

- [ ] New file added

**Reactions:**

- [ ] Upload file
  - [API Reference](https://www.dropbox.com/developers/documentation/http/documentation#files-upload)
- [ ] Download file
  - [API Reference](https://www.dropbox.com/developers/documentation/http/documentation#files-download)

---

### 12. YouTube (Google Service)

[YouTube API Documentation](https://developers.google.com/youtube/v3/docs)

**Actions:**

- [ ] New video in a subscription feed

**Reactions:**

- [ ] Retrieve video information
  - [API Reference](https://developers.google.com/youtube/v3/docs/videos)

---

### 13. RSS Feed (RSS Service)

**Actions:**

- [ ] New article available
- [ ] Article marked as favorite

**Reactions:**

- [ ] Retrieve latest article

---

### 14. X (X Service)

[X API Documentation](https://developer.x.com/en/docs)

**Actions:**

- [ ] A new message is posted in group G
- [ ] A new message containing #hashtag
- [ ] Receive a private message
- [ ] User's message gets a like
- [ ] User gains a follower

**Reactions:**

- [ ] Post a message in group G
- [ ] Follow a new user P

---

### 15. Outlook (Microsoft Service)

[Outlook API Documentation](https://learn.microsoft.com/en-us/outlook/rest/reference)

---

### 16. OneDrive (Microsoft Service)

[OneDrive API Documentation](https://learn.microsoft.com/en-us/onedrive/developer/rest-api/getting-started/?view=odsp-graph-online)

---

### 17. Outlook Calendar (Microsoft Service)

[Outlook Calendar API Documentation](https://learn.microsoft.com/en-us/graph/outlook-calendar-concept-overview)

---

### 18. Deezer (Deezer Service)

[Deezer API Documentation](https://developers.deezer.com/api)

Application registration

```text
We're not accepting new application creation at this time. Please check again later.
```

**Actions:**

- [ ] Start music playback
- [ ] Stop playback
- [ ] Get user’s current playback information

**Reactions:**

- [ ] Start playback
- [ ] Pause playback
- [ ] Change volume
- [ ] Retrieve music/artist details
