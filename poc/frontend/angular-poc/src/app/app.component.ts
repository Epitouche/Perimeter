import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CommonModule } from '@angular/common';
import { WeatherService } from './services/weather.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    CommonModule,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {
  constructor(private weatherService: WeatherService) {

  }
  ngOnInit(): void {
    this.weatherService.getWeatherData('bordeaux')
    .subscribe({
      next: (response) => {
        console.log(response);
      }
    })
  }
}
