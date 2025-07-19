import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDate} from "~/lib/utils";
import type {Price, PriceChartData} from "~/home/types";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";
import {Button} from "~/components/ui/button";
import {Book, Eye} from "lucide-react";
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";

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

function PriceOverview({good, data}: PriceOverviewProps) {
    return (
        <Card className="w-full">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>{presentTimeSpan(data.prices)}</CardDescription>
                <CardAction>
                    <Tooltip>
                        <TooltipTrigger asChild>
                            <Button variant="secondary" size="icon" className="size-8">
                                <Eye/>
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>View details</p>
                        </TooltipContent>
                    </Tooltip>
                    <Tooltip>
                        <TooltipTrigger asChild>
                            <Button variant="secondary" size="icon" className="size-8">
                                <Book/>
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>Go to tibiawiki.com</p>
                        </TooltipContent>
                    </Tooltip>
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
    PriceOverview,
}