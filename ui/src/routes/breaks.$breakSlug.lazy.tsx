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
    console.log('statsdata', statsdata);
    const highestWaves = dailyStats.waveHeightMax.sort((a, b) => b - a).map((waveHeight, index) => {
        const date = dailyStats.time[index];
        return {height: waveHeight, date};
    });

    const minSeaSurfaceTemp = hourlyStats.seaSurfaceTemperature.map((temp, index) => {
        const date = hourlyStats.time[index];
        return {temp, date};
    }).sort((a, b) => a.temp - b.temp);
    const maxSeaSurfaceTemp = hourlyStats.seaSurfaceTemperature.map((temp, index) => {
        const date = hourlyStats.time[index];
        return {temp, date};
    }).sort((a, b) => b.temp - a.temp);

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

                {statsdata && (
                    <AccordionItem key="3" aria-label="Stats (last 3 months)" title="Stats (last 3 months)">
                        <dl className="py-4 flex flex-col gap-2 w-full">
                            <dt className="font-bold">Highest waves:</dt>
                            <dd className="mb-2">
                                <ul>
                                    {highestWaves.slice(0, 3).map((wave, index) => (
                                        <li key={index}>
                                            {wave.height} m on {new Date(wave.date).toLocaleDateString(
                                            'en-US',
                                            {month: 'short', day: 'numeric', year: 'numeric'}
                                        )}
                                        </li>
                                    ))}
                                </ul>
                            </dd>

                            <dt className="font-bold">Lowest sea surface temperature:</dt>
                            <dd className="mb-2">
                                {minSeaSurfaceTemp[0].temp} °C
                                on {new Date(minSeaSurfaceTemp[0].date).toLocaleDateString(
                                'en-US',
                                {month: 'short', day: 'numeric', year: 'numeric'}
                            )}
                            </dd>

                            <dt className="font-bold">Highest sea surface temperature:</dt>
                            <dd className="mb-2">
                                {maxSeaSurfaceTemp[0].temp} °C
                                on {new Date(maxSeaSurfaceTemp[0].date).toLocaleDateString(
                                'en-US',
                                {month: 'short', day: 'numeric', year: 'numeric'}
                            )}
                            </dd>
                        </dl>
                    </AccordionItem>
                )}
            </Accordion>
        </div>
    )
}
