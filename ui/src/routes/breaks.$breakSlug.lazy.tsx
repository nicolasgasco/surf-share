import {createLazyFileRoute, useParams} from '@tanstack/react-router'
import {useFetchBreakBySlug} from "../hooks/useFetchBreakBySlug.tsx";
import {useFetchForecastBySlug} from "../hooks/useFetchForecastBySlug.tsx";
import {ForecastTable} from "../components/ForecastTable.tsx";
import {WebcamGrid} from "../components/WebcamGrid.tsx";
import {WebcamVideoPreview} from "../components/WebcamVideoPreview.tsx";

export const Route = createLazyFileRoute('/breaks/$breakSlug')({
    component: RouteComponent,
})

function RouteComponent() {
    const {breakSlug} = useParams({
        from: '/breaks/$breakSlug',
    })
    const {isLoading: isLoadingBreak, data: breakData} = useFetchBreakBySlug(breakSlug);
    const {isLoading: isLoadingForecast, data: forecastData} = useFetchForecastBySlug(breakSlug);

    if (isLoadingBreak || isLoadingForecast) {
        return (<div>
            <p>Loading...</p>
        </div>)
    }

    if (!breakData || !forecastData) {
        return <div>
            <h1>Something went wrong. Please try again later.</h1>
        </div>
    }

    const {description, imageUrls, name, region, country, videoUrl} = breakData;
    const {hourly: hourlyData} = forecastData;

    return (
        <div className="max-w-3xl flex flex-col items-center justify-center text-center">
            <h1 className="mb-4 title-1">Your break info for {name} ({region}, {country})</h1>
            <p className="mb-6 text-sm">{description}</p>

            <section className="mb-12">
                <h2 className="mb-4 title-2">Webcam</h2>

                {videoUrl && <WebcamVideoPreview videoUrl={videoUrl} name={name}/>}

                {imageUrls && imageUrls.length > 0 && <WebcamGrid imageUrls={imageUrls} name={name}/>}

            </section>
            {forecastData && (
                <section className="w-full">
                    <h2 className="mb-4 title-2">Forecast (3 days)</h2>
                    <ForecastTable hourlyData={hourlyData}/>
                </section>
            )}
        </div>
    )
}
