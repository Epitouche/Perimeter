import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment.prod';
import { WeatherData } from '../models/weather.models';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WeatherService {

  constructor(private http: HttpClient) { }

  // Calling API
  getWeatherData(cityName: string): Observable<WeatherData> {
    const url = `${environment.weatherApiBaseUrl}`;

    const params = {
      q: cityName,
      appid: environment.OpenWeatherMapKeyValue,
      units: 'metric'
    };

    return this.http.get<WeatherData>(url, { params });
  }
}
