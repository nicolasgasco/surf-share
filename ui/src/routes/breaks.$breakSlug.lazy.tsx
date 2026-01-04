import {createLazyFileRoute, useParams} from '@tanstack/react-router'
import {useFetchBreakBySlug} from "../hooks/useFetchBreakBySlug.tsx";

export const Route = createLazyFileRoute('/breaks/$breakSlug')({
    component: RouteComponent,
})

function RouteComponent() {
    const {breakSlug} = useParams({
        from: '/breaks/$breakSlug',
    })
    const {isLoading, data} = useFetchBreakBySlug(breakSlug);
    console.log(data);

    if (isLoading) {
        return (<div>
            <p>Loading...</p>
        </div>)
    }

    if (!data) {
        return <div>
            <h1>Something went wrong. Please try again later.</h1>
        </div>
    }

    const {description, name, region, country} = data;

    return (
        <div className="max-w-3xl flex flex-col items-center justify-center text-center">
            <h1 className="mb-6">Your break info for {name} ({region}, {country})</h1>
            <p>{description}</p>
        </div>
    )
}
