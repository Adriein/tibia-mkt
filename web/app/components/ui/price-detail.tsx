import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDate} from "~/lib/utils";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";
import {Button} from "~/components/ui/button";
import {Book, Eye} from "lucide-react";
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";
import {Link} from "react-router";
import type {Price, PriceChartData} from "~/lib/types";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {ToggleGroup, ToggleGroupItem} from "~/components/ui/toggle-group";

const chartConfig = {
    buyOffer: {
        label: "Buy Offer",
        color: "var(--chart-1)",
    },
    sellOffer: {
        label: "Sell Offer",
        color: "var(--chart-2)",
    },
} satisfies ChartConfig

type PriceDetailProps = {
    good: string;
    data: PriceChartData;
}

const labelFormatter = (label: string, _: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDate(label)}</span>
}

const transformValueNumberToLocale = (value: number|string): string => {
    return Intl.NumberFormat("es-Es").format(value as number).toString()
};

function presentTimeSpan(data: Price[]): string {
    const start: string = Intl
        .DateTimeFormat('es-ES', {year: "numeric", month: "short"})
        .format(new Date(data[0].createdAt));

    const end: string = Intl
        .DateTimeFormat('es-ES', {year: "numeric", month: "short"})
        .format(new Date(data[data.length - 1].createdAt));

    return `${start} - ${end}`;
}

function PriceDetail({good, data}: PriceDetailProps) {
    const [timeRange, setTimeRange] = React.useState("90d")

    /*const filteredData = chartData.filter((item) => {
        const date = new Date(item.date)
        const referenceDate = new Date("2024-06-30")
        let daysToSubtract = 90
        if (timeRange === "30d") {
            daysToSubtract = 30
        } else if (timeRange === "7d") {
            daysToSubtract = 7
        }
        const startDate = new Date(referenceDate)
        startDate.setDate(startDate.getDate() - daysToSubtract)
        return date >= startDate
    })*/

    return (
        <Card className="w-full">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>{presentTimeSpan(data.prices)}</CardDescription>
                <CardAction className="flex gap-3">
                    <ToggleGroup
                        type="single"
                        value={timeRange}
                        onValueChange={setTimeRange}
                        variant="outline"
                        className="*:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
                    >
                        <ToggleGroupItem value="90d">Last 3 months</ToggleGroupItem>
                        <ToggleGroupItem value="30d">Last 30 days</ToggleGroupItem>
                        <ToggleGroupItem value="7d">Last 7 days</ToggleGroupItem>
                    </ToggleGroup>
                </CardAction>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="aspect-auto h-[250px] w-full">
                    <LineChart
                        accessibilityLayer
                        data={data.prices}
                        margin={{
                            top: 20,
                            left: 12,
                            right: 12,
                        }}
                    >
                        <CartesianGrid vertical={false} />
                        <XAxis
                            dataKey="createdAt"
                            tickLine={false}
                            axisLine={false}
                            tickMargin={8}
                            interval="preserveStartEnd"
                            tickFormatter={formatDate}
                        />
                        <ChartTooltip
                            cursor={false}
                            content={
                                <ChartTooltipContent
                                    indicator="line"
                                    labelFormatter={labelFormatter}
                                    valueFormatter={transformValueNumberToLocale}
                                    className="w-[150px]"/>
                            }
                        />
                        <Line
                            dataKey="buyOffer"
                            type="natural"
                            stroke="var(--chart-theme-1)"
                            strokeWidth={2}
                            dot={{
                                fill: "var(--chart-theme-1)",
                            }}
                            activeDot={{
                                r: 6,
                            }}
                        />
                        <Line
                            dataKey="sellOffer"
                            type="natural"
                            stroke="var(--chart-theme-2)"
                            strokeWidth={2}
                            dot={{
                                fill: "var(--chart-theme-2)",
                            }}
                            activeDot={{
                                r: 6,
                            }}
                        />
                    </LineChart>
                </ChartContainer>
            </CardContent>
        </Card>
    );
}

export {
    PriceDetail,
}