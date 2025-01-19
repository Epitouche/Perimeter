# Services

## Table of Contents

- [Services](#services)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Must Services](#must-services)
    - [1. Spotify (Spotify Service)](#1-spotify-spotify-service)
    - [2. OpenWeatherMap (OpenWeatherMap Service)](#2-openweathermap-openweathermap-service)
    - [3. Timer (Timer Service)](#3-timer-timer-service)
    - [4. Gmail (Google Service)](#4-gmail-google-service)
    - [5. GitHub (GitHub Service)](#5-github-github-service)
    - [6. Dropbox (Dropbox Service)](#6-dropbox-dropbox-service)
  - [Optional Services](#optional-services)
    - [7. Outlook (Microsoft Service)](#7-outlook-microsoft-service)
    - [8. YouTube (Google Service)](#8-youtube-google-service)
    - [9. Facebook (Meta Service)](#9-facebook-meta-service)
    - [10. Instagram (Meta Service)](#10-instagram-meta-service)
    - [11. RSS Feed (RSS Service)](#11-rss-feed-rss-service)
    - [12. X (X Service)](#12-x-x-service)
    - [13. Google Drive (Google Service)](#13-google-drive-google-service)
    - [14. Google Calendar (Google Service)](#14-google-calendar-google-service)
    - [15. OneDrive (Microsoft Service)](#15-onedrive-microsoft-service)

---

## Main Document

Refer to the [main documentation](../../README.md).

---

## Must Services

For each service we implement, the minimum requirements are:

- **One action**
- **One reaction**

the bold action / reaction are the one we take

### 1. Spotify (Spotify Service)

[Spotify API Documentation](https://developer.spotify.com/documentation/web-api)

API limit:

- Free plan: 100,000 requests per day for the Web API (requests to retrieve information about tracks, albums, etc.).
- Limited to 5,000 requests per user per day for integrated Spotify applications.
- For actions like playing or controlling playback, it may be limited to 100 calls per second.

**Actions:**

- [x] **Get user’s current playback music**
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/get-information-about-the-users-current-playback)

**Reactions:**

- [x] **Next Music**
- [x] **Previous Music**
- [ ] **Start playback**
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/start-a-users-playback)
- [ ] **Pause playback**
  - [API Reference](https://developer.spotify.com/documentation/web-api/reference/pause-a-users-playback)
- [ ] Change volume
- [ ] **Retrieve music/artist details**

### 2. OpenWeatherMap (OpenWeatherMap Service)

[OpenWeatherMap API Documentation](https://openweathermap.org/api)

API limit: 1,000 requests per day to obtain current weather data, forecasts, etc.

**Actions:**

- [x] **Temperature above T**
- [x] **Temperature below T**
- [x] **Temperature equal T**
- [x] **Fetch weather forecast**

**Reactions:**

- [x] **Display current weather**
  - [API Reference](https://openweathermap.org/api/one-call-3#weather_overview)
- [x] **Display current temperature**

---

### 3. Timer (Timer Service)

[TimeAPI Documentation](https://www.timeapi.io/swagger/index.html)

API limit: we don't find it

**Actions:**

- [x] **Trigger at specific time T**

**Reactions:**

- [x] **Get current time**

---

### 4. Gmail (Google Service)

[Gmail API Documentation](https://developers.google.com/gmail)

API limit: up to 10 000 calls API per day for app using OAuth2

**Actions:**

- [x] **Receive a message**
- [ ] Receive a message from user X
- [ ] Receive a message with a subject containing the word X

**Reactions:**

- [x] **Send message M to recipient D**
  - [API Guide](https://developers.google.com/gmail/api/guides/sending)
  - [API Reference](https://developers.google.com/gmail/api/reference/rest/v1/users.messages/send)

---

### 5. GitHub (GitHub Service)

[GitHub API Documentation](https://docs.github.com/en/rest)

API limit: All of these requests count towards your personal rate limit of 5,000 requests per hour.

**Actions:**

- [ ] Create a new repository
- [ ] **Create a branch in repository R**
- [x] new Commit a file in repository R
- [x] new pull request in repository R
- [x] new workflow run in repository R

**Reactions:**

- [ ] **Create a branch in repository R**
  - [Branch API](https://docs.github.com/en/rest/branches/branches?apiVersion=2022-11-28#get-a-branch)
- [x] **Get Commit in repository R**
  - [File Contents API](https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-or-update-file-contents)
- [x] Get workflow run in repository R

---

### 6. Dropbox (Dropbox Service)

[Dropbox API Documentation](https://www.dropbox.com/developers/documentation/http/documentation)

API limit: no obvious limit, but get code `429` when limit reach.

**Actions:**

- [x] **New file added in folder F**

**Reactions:**

- [ ] **Upload file**
  - [API Reference](https://www.dropbox.com/developers/documentation/http/documentation#files-upload)
- [ ] **Download file**
  - [API Reference](https://www.dropbox.com/developers/documentation/http/documentation#files-download)
- [x] Download file content in dropbox file

---

## Optional Services

### 7. Outlook (Microsoft Service)

[Outlook API Documentation](https://learn.microsoft.com/en-us/outlook/rest/reference)

**Actions:**

- [x] **Receive a message**

**Reactions:**

- [x] **Send message M to recipient D**

---

### 8. YouTube (Google Service)

[YouTube API Documentation](https://developers.google.com/youtube/v3/docs)

**Actions:**

- [ ] New video in a subscription feed

**Reactions:**

- [ ] Retrieve video information
  - [API Reference](https://developers.google.com/youtube/v3/docs/videos)

---

### 9. Facebook (Meta Service)

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

### 10. Instagram (Meta Service)

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

### 11. RSS Feed (RSS Service)

**Actions:**

- [ ] New article available
- [ ] Article marked as favorite

**Reactions:**

- [ ] Retrieve latest article

---

### 12. X (X Service)

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

### 13. Google Drive (Google Service)

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

### 14. Google Calendar (Google Service)

[Google Calendar API Documentation](https://developers.google.com/calendar)

**Actions:**

- [ ] Event starts
- [ ] Event ends
- [ ] Add an event

**Reactions:**

- [ ] Add an event
  - [API Reference](https://developers.google.com/calendar/api/v3/reference/events/insert)

---

### 15. OneDrive (Microsoft Service)

[OneDrive API Documentation](https://learn.microsoft.com/en-us/onedrive/developer/rest-api/getting-started/?view=odsp-graph-online)
