import {createLazyFileRoute, useNavigate} from '@tanstack/react-router'
import {useFetchBreaks} from "../hooks/useFetchBreaks.tsx";

export const Route = createLazyFileRoute('/')({
    component: RouteComponent,
})

function RouteComponent() {
    const {breaks} = useFetchBreaks()
    const navigate = useNavigate()

    const handleChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        const selectedSlug = event.target.value;
        navigate({
            to: `/breaks/${selectedSlug}`,
        })
    }

    return (
        <>
            <h1 className="mb-4">Welcome to SurfShare</h1>
            <p className="mb-6">Start by choosing one of the surf breaks below to find more about them.</p>

            <select name="breaks"
                    className="text-md bg-white disabled:opacity-50 disabled:cursor-not-allowed text-black p-2 rounded shadow-lg font-medium min-w-[200px] text-center"
                    disabled={breaks.length === 0}
                    defaultValue="default"
                    onChange={handleChange}
            >
                <option disabled value="default">Choose a break</option>
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
