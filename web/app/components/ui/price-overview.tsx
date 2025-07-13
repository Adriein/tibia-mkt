import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDate} from "~/lib/utils";
import type {PriceChartData} from "~/home/types";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";

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

const labelFormatter = (label: string, payload: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDate(label)}</span>
}

const transformValueNumberToLocale = (value: number|string): string => {
    return Intl.NumberFormat("es-Es").format(value as number).toString()
};

function PriceOverview({good, data}: PriceOverviewProps) {
    return (
        <Card className="w-full">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>January - June 2024</CardDescription>
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
                            //ticks={xAxisTick(data.prices, data.chartMetadata.xAxisTick)}
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