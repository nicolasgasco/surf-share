export interface HourlyForecast {
    time: string[];
    waveHeight: number[];
    wavePeriod: number[];
    waveDirection: number[];
    seaSurfaceTemperature: number[];
    seaLevelHeightMsl: number[];
}

export interface Forecast {
    latitude: number;
    longitude: number;
    hourly: HourlyForecast;
}
