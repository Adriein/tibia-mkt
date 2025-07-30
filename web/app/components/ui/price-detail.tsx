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
import {useIsMobile} from "~/lib/hooks";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "~/components/ui/tabs"

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
    isMobile: boolean;
}

type PriceDetailChartProps = {
    good: string;
    data: PriceChartData;
    type: typeof SELL_CHART | typeof BUY_CHART
    t: DetailTranslations;
    isMobile: boolean;
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

function PriceDetailChart({good, type, data, t, isMobile}: PriceDetailChartProps): React.ReactElement {
    const [timeRange, setTimeRange] = React.useState(isMobile? "7d" : "90d");

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
        <Card className="w-full @container/card">
            <CardHeader>
                <CardTitle>{beautifyCamelCase(good)}</CardTitle>
                <CardDescription>{type === BUY_CHART? t.buyOffer : t.sellOffer}</CardDescription>
                <CardAction className="flex gap-3">
                    <ToggleGroup
                        type="single"
                        value={timeRange}
                        onValueChange={setTimeRange}
                        variant="outline"
                        className="hidden *:data-[slot=toggle-group-item]:!px-4 @[767px]/card:flex"
                    >
                        <ToggleGroupItem value="90d">{t.timeSpan3Months}</ToggleGroupItem>
                        <ToggleGroupItem value="30d">{t.timeSpan30Days}</ToggleGroupItem>
                        <ToggleGroupItem value="7d">{t.timeSpan7ays}</ToggleGroupItem>
                    </ToggleGroup>
                    <Select value={timeRange} onValueChange={setTimeRange}>
                        <SelectTrigger
                            className="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate @[767px]/card:hidden"
                            size="sm"
                            aria-label="Select a value"
                        >
                            <SelectValue placeholder="Last 3 months" />
                        </SelectTrigger>
                        <SelectContent className="rounded-xl">
                            <SelectItem value="90d" className="rounded-lg">
                                Last 3 months
                            </SelectItem>
                            <SelectItem value="30d" className="rounded-lg">
                                Last 30 days
                            </SelectItem>
                            <SelectItem value="7d" className="rounded-lg">
                                Last 7 days
                            </SelectItem>
                        </SelectContent>
                    </Select>
                </CardAction>
            </CardHeader>
            <CardContent>
                <ChartContainer config={chartConfig} className="aspect-auto h-[150px] sm:h-[200px] w-full">
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
        <Card>
            <CardContent>
                <div className="flex justify-center align-center gap-24">
                    <div className="flex flex-col items-center gap-2">
                        <span>{type === BUY_CHART? t.buyOffersMean : t.sellOffersMean}</span>
                        <span>
                            {type === BUY_CHART?
                                transformValueNumberToLocale(data.buyOffersMean) :
                                transformValueNumberToLocale(data.sellOffersMean)
                            }
                        </span>
                    </div>
                    <div className="flex flex-col align-center items-center gap-2">
                        <span>{type === BUY_CHART? t.buyOffersMedian : t.sellOffersMedian}</span>
                        <span>
                            {type === BUY_CHART?
                                transformValueNumberToLocale(data.buyOffersMedian) :
                                transformValueNumberToLocale(data.sellOffersMedian)
                            }
                        </span>
                    </div>
                    <div className="flex flex-col align-center items-center gap-2">
                        <span>{type === BUY_CHART? t.buyOfferStdDeviation : t.sellOfferStdDeviation}</span>
                        <span>
                            {type === BUY_CHART?
                                transformValueNumberToLocale(data.buyOffersStdDeviation) :
                                transformValueNumberToLocale(data.sellOffersStdDeviation)
                            }
                        </span>
                    </div>
                </div>
            </CardContent>
        </Card>
    );
}


function PriceDetail({good, prices, statistics, t, isMobile}: PriceDetailProps) {
    return isMobile? (
            <div className="w-full">
                <Tabs defaultValue="chart">
                    <TabsList>
                        <TabsTrigger value="chart">Charts</TabsTrigger>
                        <TabsTrigger value="info">Info</TabsTrigger>
                    </TabsList>
                    <TabsContent value="chart" className="flex flex-col gap-3">
                        <PriceDetailChart good={good} type={BUY_CHART} data={prices} t={t} isMobile={isMobile}/>
                        <PriceDetailChart good={good} type={SELL_CHART} data={prices} t={t} isMobile={isMobile}/>
                    </TabsContent>
                    <TabsContent value="info" className="flex flex-col gap-3">
                        <PriceDetailStatistics data={statistics} t={t} type={BUY_CHART}/>
                        <PriceDetailStatistics data={statistics} t={t} type={SELL_CHART}/>
                    </TabsContent>
                </Tabs>
            </div>) : (
            <div className="flex flex-col gap-3 w-full">
                <PriceDetailStatistics data={statistics} t={t} type={BUY_CHART}/>
                <PriceDetailChart good={good} type={BUY_CHART} data={prices} t={t} isMobile={isMobile}/>
                <PriceDetailStatistics data={statistics} t={t} type={SELL_CHART}/>
                <PriceDetailChart good={good} type={SELL_CHART} data={prices} t={t} isMobile={isMobile}/>
            </div>
        );
}

export {
    PriceDetail,
}