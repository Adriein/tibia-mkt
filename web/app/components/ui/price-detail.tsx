import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDate} from "~/lib/utils";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";
import type {Price, PriceChartData} from "~/lib/types";
import {ToggleGroup, ToggleGroupItem} from "~/components/ui/toggle-group";
import type {DetailTranslations} from "~/locale/loc";
import type {DetailPageStatisticsData} from "~/routes/detail/types";

const SELL_CHART = "sell";
const BUY_CHART = "buy";

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
    prices: PriceChartData;
    statistics: DetailPageStatisticsData;
    t: DetailTranslations;
}

type PriceDetailChartProps = {
    good: string;
    data: PriceChartData;
    type: typeof SELL_CHART | typeof BUY_CHART
    t: DetailTranslations;
}

type PriceDetailStatisticsProps = {
    type: typeof SELL_CHART | typeof BUY_CHART
    data: DetailPageStatisticsData
    t: DetailTranslations;
}

const labelFormatter = (label: string, _: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDate(label)}</span>
}

const transformValueNumberToLocale = (value: number|string): string => {
    return Intl.NumberFormat("es-Es").format(value as number).toString()
};

function PriceDetailChart({good, type, data, t}: PriceDetailChartProps) {
    const [timeRange, setTimeRange] = React.useState("90d");

    const filteredData: Price[] = data.prices.filter((item: Price): boolean => {
        const date = new Date(item.createdAt);

        let daysToSubtract = 90;

        if (timeRange === "30d") {
            daysToSubtract = 30
        } else if (timeRange === "7d") {
            daysToSubtract = 7
        }

        const startDate = new Date(data.prices.at(-1)?.createdAt!);

        startDate.setDate(startDate.getDate() - daysToSubtract)

        return date >= startDate
    });

    return (
        <Card className="w-full">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>{type === BUY_CHART? t.buyOffer : t.sellOffer}</CardDescription>
                <CardAction className="flex gap-3">
                    <ToggleGroup
                        type="single"
                        value={timeRange}
                        onValueChange={setTimeRange}
                        variant="outline"
                        className="*:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
                    >
                        <ToggleGroupItem value="90d">{t.timeSpan3Months}</ToggleGroupItem>
                        <ToggleGroupItem value="30d">{t.timeSpan30Days}</ToggleGroupItem>
                        <ToggleGroupItem value="7d">{t.timeSpan7ays}</ToggleGroupItem>
                    </ToggleGroup>
                </CardAction>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="aspect-auto h-[200px] w-full">
                    <LineChart
                        accessibilityLayer
                        data={filteredData}
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
                        {type === BUY_CHART? <Line
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
                            /> :
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
                            />}
                    </LineChart>
                </ChartContainer>
            </CardContent>
        </Card>
    );
}

function PriceDetailStatistics({data, type, t}: PriceDetailStatisticsProps) {
    return (
        <Card className="w-full">
            <CardContent>
                <div className="flex justify-center align-center gap-24">
                    <div className="flex flex-col items-center gap-2">
                        <span>Buy offer mean</span>
                        <span>{transformValueNumberToLocale(data.buyOffersMean)}</span>
                    </div>
                    <div className="flex flex-col align-center items-center gap-2">
                        <span>Buy offer median</span>
                        <span>{transformValueNumberToLocale(data.buyOffersMedian)}</span>
                    </div>
                    <div className="flex flex-col align-center items-center gap-2">
                        <span>Buy offer std deviation</span>
                        <span>{transformValueNumberToLocale(data.buyOffersStdDeviation)}</span>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}


function PriceDetail({good, prices, statistics, t}: PriceDetailProps) {
    return (
        <div className="flex flex-col gap-3 w-full">
            <PriceDetailStatistics data={statistics} t={t} type={BUY_CHART}/>
            <PriceDetailChart good={good} type={BUY_CHART} data={prices} t={t}/>
            <PriceDetailStatistics data={statistics} t={t} type={SELL_CHART}/>
            <PriceDetailChart good={good} type={SELL_CHART} data={prices} t={t}/>
        </div>
    );
}

export {
    PriceDetail,
}