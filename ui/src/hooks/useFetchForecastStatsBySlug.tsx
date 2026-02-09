import type {ForecastStats} from "../types/forecastStats.tsx";
import {useEffect, useState} from "react";
import camelize from "camelize";

export function useFetchForecastStatsBySlug(slug: string): {
    data: ForecastStats | null,
    isLoading: boolean,
    error: string | null
} {
    const [statsData, setStatsData] = useState<ForecastStats | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function fetchStats() {
            setIsLoading(true);
            setError(null);

            try {
                const response = await fetch(`${import.meta.env.VITE_API_URL}/forecast/stats/${slug}`);

                if (!response.ok) {
                    const errorMsg = `Failed to fetch forecast stats: ${response.statusText}`;
                    console.error(errorMsg);
                    setError(errorMsg);
                    setStatsData(null);
                    return;
                }

                const statsInfo = await response.json();
                setStatsData(camelize(statsInfo));
                setError(null);
            } catch (err) {
                const errorMsg = err instanceof Error ? err.message : 'An error occurred';
                console.error('Forecast stats fetch error:', errorMsg);
                setError(errorMsg);
                setStatsData(null);
            } finally {
                setIsLoading(false);
            }
        }

        fetchStats();
    }, [slug]);

    return {
        data: statsData,
        isLoading,
        error,
    }
}
