import {useFetchBreaks} from "../hooks/useFetchBreaks.tsx";

export function HomePage() {
    const {breaks} = useFetchBreaks()

    return (
        <>
            <h1 className="mb-2">Welcome to SurfShare</h1>
            <p className="mb-6">Start by choosing one of the surf breaks below</p>

            <select name="breaks" className="text-md bg-white text-black p-2 rounded shadow-lg font-medium">
                {
                    breaks.map((breakSummary) => (
                        <option key={breakSummary.id} value={breakSummary.slug}>
                            {breakSummary.name}
                        </option>
                    ))
                }
            </select>
        </>
    )
}