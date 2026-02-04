export interface HourlyForecast {
    time: string[];
    waveHeight: number[];
    wavePeriod: number[];
    waveDirection: number[];
    seaSurfaceTemperature: number[];
    seaLevelHeightMsl: number[];
    temperature2m: number[];
    windSpeed10m: number[];
    windDirection10m: number[];
}

export interface Forecast {
    latitude: number;
    longitude: number;
    hourly: HourlyForecast;
}
