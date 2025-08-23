import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDateToShortForm} from "~/lib/utils";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";
import {Button} from "~/components/ui/button";
import {Book, Eye} from "lucide-react";
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";
import {Link, useSearchParams} from "react-router";
import type {Price, PriceChartData} from "~/lib/types";
import type {HomePagePriceDataPoint, HomePriceChartData} from "~/routes/home/types";
import * as sea from "node:sea";

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

type PriceOverviewProps = {
    good: string;
    data: HomePriceChartData;
}

const labelFormatter = (label: string, _: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDateToShortForm(label)}</span>
}

const transformValueNumberToLocale = (value: number|string): string => {
    return Intl.NumberFormat("es-Es").format(value as number).toString()
};

function presentTimeSpan(data: HomePagePriceDataPoint[]): string {
    const start: string = Intl
        .DateTimeFormat('es-ES', {year: "numeric", month: "short"})
        .format(new Date(data[0].createdAt));

    const end: string = Intl
        .DateTimeFormat('es-ES', {year: "numeric", month: "short"})
        .format(new Date(data[data.length - 1].createdAt));

    return `${start} - ${end}`;
}

function PriceOverview({good, data}: PriceOverviewProps) {
    const [searchParams] = useSearchParams();
    const lang: string|null = searchParams.get("lang");

    return (
        <Card className="w-full hover:bg-muted/50">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>{presentTimeSpan(data.dataPoints)}</CardDescription>
                <CardAction className="flex gap-3">
                    <Tooltip>
                        <TooltipTrigger asChild>
                            <Button asChild variant="secondary" size="icon" className="size-8">
                                <Link to={lang?`/${good}/detail?lang=${lang}` : `/${good}/detail`}>
                                    <Eye/>
                                </Link>
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>View details</p>
                        </TooltipContent>
                    </Tooltip>
                    <Tooltip>
                        <TooltipTrigger asChild>
                            <Button variant="secondary" size="icon" className="size-8">
                                <Link
                                    to={`https://tibia.fandom.com/wiki/${good}`}
                                    aria-label="Go to Tibia Wiki"
                                    target="_blank"
                                    rel="noreferrer"
                                >
                                    <Book/>
                                </Link>
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>Tibia Wiki</p>
                        </TooltipContent>
                    </Tooltip>
                </CardAction>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="aspect-auto h-[250px] w-full">
                    <LineChart
                        accessibilityLayer
                        data={data.dataPoints}
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
                            tickFormatter={formatDateToShortForm}
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
                            dataKey="sellPrice"
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
                            dataKey="buyPrice"
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
    PriceOverview,
}