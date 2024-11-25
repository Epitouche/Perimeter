# POC Angular

## Table Of Contents

- [POC Angular](#poc-angular)
  - [Table Of Contents](#table-of-contents)
  - [Other Documents](#other-documents)
  - [**Angular Weather App**](#angular-weather-app)
    - [**Installation and Setup**](#installation-and-setup)
      - [**1. Prerequisites**](#1-prerequisites)
      - [**2. Clone the Repository**](#2-clone-the-repository)
      - [**3. Install Dependencies**](#3-install-dependencies)
      - [**4. Start the Application**](#4-start-the-application)
    - [**Environment Configuration**](#environment-configuration)
      - [**Modifying the Environment**](#modifying-the-environment)
    - [**Using the OpenWeatherMap API**](#using-the-openweathermap-api)
    - [**Limitations of the API**](#limitations-of-the-api)

## Other Documents

[Main Documentation](../../../README.md)
[Comparative Study](../README.md)

---

## **Angular Weather App**

This project is a simple Angular-based POC, as a web application that fetches and displays weather data for a given city using the OpenWeatherMap API.

---

### **Installation and Setup**

Follow the steps below to set up and run the application on your local machine.

#### **1. Prerequisites**
- **Node.js**: Install [Node.js](https://nodejs.org/) (LTS recommended).
- **Angular CLI**: Install Angular CLI globally:
  ```bash
  npm install -g @angular/cli
  ```

#### **2. Clone the Repository**
Clone the repository using the following command:
```bash
git clone <repository-url>
```
Replace `<repository-url>` with the actual repository URL.

#### **3. Install Dependencies**
Navigate to the project folder and install the dependencies:
```bash
cd <project-folder>
npm install
```

#### **4. Start the Application**
To start the Angular application, run:
```bash
ng serve -o
```
This will host the app on `http://localhost:4200/`.

---

### **Environment Configuration**

The application uses an environment file to store sensitive data such as the API key for the OpenWeatherMap API.

#### **Modifying the Environment**
1. Open the environment configuration file located at:
   - `src/environments/environment.ts` (for development)
   - `src/environments/environment.prod.ts` (for production)

2. Update the file with your OpenWeatherMap API key:
   ```typescript
   export const environment = {
     production: false,
     weatherApiBaseUrl: 'https://api.openweathermap.org/data/2.5/weather',
     OpenWeatherMapApiKey: 'YOUR_API_KEY', // Replace with your actual API key
   };
   ```

---

### **Using the OpenWeatherMap API**

This app integrates the OpenWeatherMap API to fetch weather data. The API is queried using the following parameters:
- **`q`**: The city name.
- **`lang`**: The language for the results (default: `'FR'`).
- **`units`**: Units for temperature (default: `'metric'`).
- **`appid`**: The API key.

Example API call:
```plaintext
https://api.openweathermap.org/data/2.5/weather?q=Paris&lang=FR&units=metric&appid=YOUR_API_KEY
```

---

### **Limitations of the API**
1. **Rate Limits**:
   - Free tier allows **60 API calls per minute**.
   - Exceeding the limit will result in HTTP 429 (Too Many Requests).

2. **Data Accuracy**:
   - Weather data might not always reflect real-time conditions accurately, as it's based on various data models.

3. **City Search**:
   - The city name must match exactly. Ambiguous or misspelled city names may lead to incorrect results or errors.

4. **Free Tier Restrictions**:
   - Some advanced features (e.g., hourly forecasts, historical data) require a paid plan.
