import type {Break} from "../types/breaks.tsx";
import {useEffect, useState} from "react";
import camelize from "camelize";

export function useFetchBreakBySlug(slug: string): { data: Break | null, isLoading: boolean } {
    const [breakData, setBreakData] = useState<Break | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);

    useEffect(() => {
        async function fetchBreak() {
            setIsLoading(true);
            const response = await fetch(`http://localhost:8080/breaks/${slug}`);

            if (!response.ok) {
                console.error('Failed to fetch break:', response.statusText);
                setBreakData(null);
                return;
            }

            const breakInfo = await response.json();
            setBreakData(camelize(breakInfo));

            setIsLoading(false);
        }

        fetchBreak();
    }, [slug]);

    return {
        data: breakData,
        isLoading,
    }
}