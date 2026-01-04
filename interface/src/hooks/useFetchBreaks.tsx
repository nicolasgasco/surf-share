import {useEffect, useState} from "react";
import type {BreakSummary} from "../types/breakSummary.tsx";

export function useFetchBreaks(): {
    breaks: BreakSummary[];
} {
    const [breaks, setBreaks] = useState<BreakSummary[]>([]);

    useEffect(() => {
        async function fetchBreaks() {
            const response = await fetch('http://localhost:8080/breaks');
            const {breaks}: {
                breaks: BreakSummary[];
                count: number;
            } = await response.json();

            setBreaks(breaks);
        }

        fetchBreaks();
    }, []);

    return {breaks};
}