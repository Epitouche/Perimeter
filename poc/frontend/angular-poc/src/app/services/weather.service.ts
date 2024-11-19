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

  getWeatherData(cityName: string, lang: string = 'FR'): Observable<WeatherData> {
    const url = `${environment.weatherApiBaseUrl}/city/${cityName}/${lang}`;

    const headers = {
      'x-rapidapi-key': environment.XRapidAPIKeyHeaderValue,
      'x-rapidapi-host': environment.XRapidAPIUrlHostHeaderValue,
    };

    return this.http.get<WeatherData>(url, { headers });
  }
}
