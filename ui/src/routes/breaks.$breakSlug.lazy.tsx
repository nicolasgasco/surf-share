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

    const {description, imageUrls, name, region, country, videoUrl} = data;

    return (
        <div className="max-w-3xl flex flex-col items-center justify-center text-center">
            <h1 className="mb-6">Your break info for {name} ({region}, {country})</h1>
            <p className="mb-12 text-md">{description}</p>

            {videoUrl && (
                <iframe src={videoUrl} title={`Webcam feed for ${name}`}
                        className="w-full h-full max-w-[800px] aspect-video" allowFullScreen/>
            )}

            {imageUrls && imageUrls.length > 0 && (
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mt-12"
                     aria-label={`Webcam images for ${name}`}>
                    {imageUrls.map((url) => (
                        <a href={url} rel="noopener noreferrer" target="_blank" aria-hidden="true">
                            <img key={url} src={url} className=" w-full h-auto rounded-lg shadow-md" alt=""/>
                        </a>
                    ))}
                </div>
            )}
        </div>
    )
}
