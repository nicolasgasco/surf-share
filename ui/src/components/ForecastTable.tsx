import {degreeToDirection} from "../utils/degreeToDirection.ts";
import type {HourlyForecast} from "../types/forecast.tsx";

interface ForecastTableProps {
    hourlyData: HourlyForecast;
}

interface ForecastRowProps {
    label: string;
    data: number[];
    format: (value: number) => string;
}

function ForecastRow({label, data, format}: ForecastRowProps) {
    return (
        <tr className="border-b border-dotted border-gray-300 [&>td:nth-child(even)]:bg-gray-100/5">
            <td className="text-[0.6rem] font-normal py-1">{label}</td>
            {data.map((value, index) => (
                <td key={index} className="text-xs font-normal p-2">
                    {format(value)}
                </td>
            ))}
        </tr>
    );
}

function DateHeader({dateGroups}: { dateGroups: Record<string, number> }) {
    return (
        <tr>
            <th></th>
            {Object.entries(dateGroups).map(([date, count]) => (
                <th key={date} colSpan={count}
                    className="text-xs font-normal p-2 border-r border-gray-300 text-left">
                    {new Date(date).toLocaleDateString('en-GB', {
                        year: 'numeric',
                        month: 'long',
                        day: 'numeric'
                    })}
                </th>
            ))}
        </tr>
    );
}

function TimeHeader({times}: { times: string[] }) {
    return (
        <tr className="border-b-2 border-gray-300">
            <th></th>
            {times.map((time, index) => (
                <th key={index} className="text-xs font-normal p-2">
                    {time.split('T')[1]}
                </th>
            ))}
        </tr>
    );
}

export function ForecastTable({hourlyData}: ForecastTableProps) {
    const dateGroups = hourlyData.time.reduce((acc, time) => {
        const date = time.split('T')[0];
        if (!acc[date]) acc[date] = 0;
        acc[date]++;
        return acc;
    }, {} as Record<string, number>);

    return (
        <div className="overflow-x-auto w-full mb-8">
            <table className="table-auto border-collapse w-full">
                <thead>
                <DateHeader dateGroups={dateGroups}/>
                <TimeHeader times={hourlyData.time}/>
                </thead>
                <tbody>
                <ForecastRow label="Wave Height (m)" data={hourlyData.waveHeight} format={(v) => v.toFixed(2)}/>
                <ForecastRow label="Wave period (s)" data={hourlyData.wavePeriod} format={(v) => v.toFixed(0)}/>
                <tr className="border-b border-dotted border-gray-300 [&>td:nth-child(even)]:bg-gray-100/5">
                    <td className="text-[0.6rem] font-normal py-1">Wave direction</td>
                    {hourlyData.waveDirection.map((value, index) => (
                        <td key={index} className="text-xs font-normal p-2">
                            {degreeToDirection(value)}
                        </td>
                    ))}
                </tr>
                <ForecastRow label="Water level (m)" data={hourlyData.seaLevelHeightMsl} format={(v) => v.toFixed(2)}/>
                <ForecastRow label="Water Temperature (°C)" data={hourlyData.seaSurfaceTemperature}
                             format={(v) => v.toFixed(1)}/>
                </tbody>
            </table>
        </div>
    );
}
