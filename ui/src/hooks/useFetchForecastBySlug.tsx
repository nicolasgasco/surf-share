import type {Forecast} from "../types/forecast.tsx";
import {useEffect, useState} from "react";
import camelize from "camelize";

export function useFetchForecastBySlug(slug: string): {
    data: Forecast | null,
    isLoading: boolean,
    error: string | null
} {
    const [forecastData, setForecastData] = useState<Forecast | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function fetchForecast() {
            setIsLoading(true);
            setError(null);

            try {
                const response = await fetch(`${import.meta.env.VITE_API_URL}/forecast/${slug}`);

                if (!response.ok) {
                    const errorMsg = `Failed to fetch forecast: ${response.statusText}`;
                    console.error(errorMsg);
                    setError(errorMsg);
                    setForecastData(null);
                    return;
                }

                const forecastInfo = await response.json();
                setForecastData(camelize(forecastInfo));
                setError(null);
            } catch (err) {
                const errorMsg = err instanceof Error ? err.message : 'An error occurred';
                console.error('Forecast fetch error:', errorMsg);
                setError(errorMsg);
                setForecastData(null);
            } finally {
                setIsLoading(false);
            }
        }

        fetchForecast();
    }, [slug]);

    return {
        data: forecastData,
        isLoading,
        error,
    }
}
