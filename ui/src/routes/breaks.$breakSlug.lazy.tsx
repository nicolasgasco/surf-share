import {createLazyFileRoute, useParams} from '@tanstack/react-router'
import {useFetchBreakBySlug} from "../hooks/useFetchBreakBySlug.tsx";
import {useFetchForecastBySlug} from "../hooks/useFetchForecastBySlug.tsx";
import {ForecastTable} from "../components/ForecastTable.tsx";
import {WebcamGrid} from "../components/WebcamGrid.tsx";
import {WebcamVideoPreview} from "../components/WebcamVideoPreview.tsx";
import {useFetchForecastStatsBySlug} from "../hooks/useFetchForecastStatsBySlug.tsx";
import {Accordion, AccordionItem} from "@heroui/accordion";

export const Route = createLazyFileRoute('/breaks/$breakSlug')({
    component: RouteComponent,
})

function RouteComponent() {
    const {breakSlug} = useParams({
        from: '/breaks/$breakSlug',
    })
    const {isLoading: isLoadingBreak, data: breakData} = useFetchBreakBySlug(breakSlug);
    const {isLoading: isLoadingForecast, data: forecastData} = useFetchForecastBySlug(breakSlug);
    const {isLoading: isLoadingStats, data: statsdata} = useFetchForecastStatsBySlug(breakSlug);

    if (isLoadingBreak || isLoadingForecast || isLoadingStats) {
        return (<div>
            <p>Loading...</p>
        </div>)
    }

    if (!breakData || !forecastData || !statsdata) {
        return <div>
            <h1 className="title-1">Something went wrong. Please try again later.</h1>
        </div>
    }

    const {description, imageUrls, name, region, country, videoUrl} = breakData;
    const {hourly: hourlyData} = forecastData;
    const {hourly: hourlyStats, daily: dailyStats} = statsdata;
    console.log('Stats data:', hourlyStats, dailyStats);

    return (
        <div className="max-w-3xl flex flex-col items-center justify-center text-center">
            <h1 className="mb-4 title-1">Your break info for {name} ({region}, {country})</h1>
            <p className="mb-6 text-sm">{description}</p>

            <Accordion>
                <AccordionItem key="1" aria-label="Webcam" title="Webcam">
                    <div className="py-4">
                        {videoUrl && <WebcamVideoPreview videoUrl={videoUrl} name={name}/>}

                        {imageUrls && imageUrls.length > 0 && <WebcamGrid imageUrls={imageUrls} name={name}/>}
                    </div>
                </AccordionItem>

                {forecastData && (
                    <AccordionItem key="2" aria-label="Forecast (3 days)" title="Forecast (3 days)">
                        <div className="py-4">
                            <ForecastTable hourlyData={hourlyData}/>
                        </div>
                    </AccordionItem>
                )}
            </Accordion>
        </div>
    )
}
