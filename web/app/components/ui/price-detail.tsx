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
import {Badge} from "~/components/ui/badge";
import {TrendingUp, TrendingDown} from "lucide-react";

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

type PriceDetailStatsCardProps = {
    title: string;
    trend?: "up" | "down";
    change?: number;
    value: number | string;
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

function PriceDetailStatsCard({title, trend, change, value, t}: PriceDetailStatsCardProps) {
    return (
        <Card className="relative">
            <CardHeader className="pb-2">
                <CardTitle className="text-sm font-medium flex items-center justify-between">
                    {title}
                    {trend && (
                        <div
                            className={`flex items-center gap-1 text-xs font-medium ${
                                trend === "up"
                                    ? "text-green-600 dark:text-green-400"
                                    : trend === "down"
                                        ? "text-red-600 dark:text-red-400"
                                        : "text-gray-600 dark:text-gray-400"
                            }`}
                        >
                            {trend === "up" ? <TrendingUp className="w-3 h-3" /> : <TrendingDown className="w-3 h-3" />}
                            {change && `${change > 0 ? "+" : ""}${change}%`}
                        </div>
                    )}
                </CardTitle>
            </CardHeader>
            <CardContent>
                <div className="text-2xl font-bold">{transformValueNumberToLocale(value)}</div>
            </CardContent>
        </Card>
    );
}


function PriceDetail({good, prices, statistics, t, isMobile}: PriceDetailProps) {
    return (
        <div className="min-h-screen p-6">
            <div className="max-w-7xl mx-auto space-y-6">
                {/*Header Section */}
                <div className="flex items-center justify-between">
                    <div>
                        <h1 className="text-3xl font-bold">Market Analytics</h1>
                        <p className="mt-1">Honeycomb Trading Data</p>
                    </div>
                    <Badge
                        variant="secondary"
                        className="flex items-center gap-2 bg-green-500/10 text-green-600 dark:text-green-400 border-green-500/20 dark:border-green-500/30"
                    >
                        <div className="w-2 h-2 bg-green-500 dark:bg-green-400 rounded-full animate-pulse"></div>
                        Live Data
                    </Badge>
                </div>

                {/* Buy Offers Section */}
                <div className="space-y-4">
                    <div className="flex items-center gap-3">
                        <div className="w-3 h-3 bg-[var(--chart-theme-1)] rounded-full shadow-lg shadow-amber-500/50"></div>
                        <h2 className="text-xl font-semibold">Buy Offers</h2>
                        <Badge
                            variant="outline"
                            className="text-amber-600 dark:text-amber-400 border-amber-500/30 dark:border-amber-500/50"
                        >
                            Trending Up
                        </Badge>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <PriceDetailStatsCard
                            title="Average Price"
                            value={statistics.buyOffersMean}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                        <PriceDetailStatsCard
                            title="Median Price"
                            value={statistics.buyOffersMedian}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                        <PriceDetailStatsCard
                            title="Price Volatility"
                            value={statistics.buyOffersStdDeviation}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                    </div>

                    <PriceDetailChart good={good} type={BUY_CHART} data={prices} t={t} isMobile={isMobile}/>
                </div>

                {/* Sell Offers Section */}
                <div className="space-y-4">
                    <div className="flex items-center gap-3">
                        <div className="w-3 h-3 bg-[var(--chart-theme-2)] rounded-full shadow-lg shadow-blue-500/50"></div>
                        <h2 className="text-xl font-semibold">Sell Offers</h2>
                        <Badge
                            variant="outline"
                            className="text-red-600 dark:text-red-400 border-red-500/30 dark:border-red-500/50"
                        >
                            Trending Down
                        </Badge>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <PriceDetailStatsCard
                            title="Average Price"
                            value={statistics.sellOffersMean}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                        <PriceDetailStatsCard
                            title="Median Price"
                            value={statistics.sellOffersMedian}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                        <PriceDetailStatsCard
                            title="Price Volatility"
                            value={statistics.sellOffersStdDeviation}
                            trend={"up"}
                            change={10}
                            t={t}
                        />
                    </div>
                    <PriceDetailChart good={good} type={SELL_CHART} data={prices} t={t} isMobile={isMobile}/>
                </div>
            </div>
        </div>
    );
}

export {
    PriceDetail,
}