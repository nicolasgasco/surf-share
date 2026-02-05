import {useEffect, useState} from "react";
import type {Breaks} from "../types/breaks.tsx";

export function useFetchBreaks(): {
    breaks: Breaks[];
} {
    const [breaks, setBreaks] = useState<Breaks[]>([]);

    useEffect(() => {
        async function fetchBreaks() {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/breaks`);

            if (!response.ok) {
                console.error('Failed to fetch breaks:', response.statusText);
                return {
                    breaks: []
                }
            }

            const {breaks}: {
                breaks: Breaks[];
                count: number;
            } = await response.json();

            setBreaks(breaks);
        }

        fetchBreaks();
    }, []);

    return {breaks};
}