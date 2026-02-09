export interface HourlyStats {
    time: string[];
    waveHeight: number[];
    wavePeriod: number[];
    waveDirection: number[];
    seaSurfaceTemperature: number[];
    seaLevelHeightMsl: number[];
}

export interface ForecastStats {
    latitude: number;
    longitude: number;
    generationtime_ms: number;
    timezone: string;
    elevation: number;
    hourlyUnits: {
        time: string;
        waveHeight: string;
        wavePeriod: string;
        waveDirection: string;
        seaSurfaceTemperature: string;
        seaLevelHeightMsl: string;
    };
    dailyUnits: {
        time: string;
        waveHeightMax: string;
        waveDirectionDominant: string;
        wavePeriodMax: string;
    };
    hourly: HourlyStats;
    daily: {
        time: string[];
        waveHeightMax: number[];
        waveDirectionDominant: number[];
        wavePeriodMax: number[];
    };
}
