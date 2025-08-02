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
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {Badge} from "~/components/ui/badge";
import {TrendingUp, TrendingDown, CircleQuestionMark, ChartNoAxesCombined, CalendarClock} from "lucide-react";
import {Avatar, AvatarImage} from "~/components/ui/avatar";
import HoneycombGif from "~/assets/honeycomb.gif";
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";

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
    info: string;
}

const labelFormatter = (label: string, _: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDate(label)}</span>
}

const transformValueNumberToLocale = (value: number|string): string => {
    return Intl.NumberFormat("es-Es").format(value as number).toString()
};

function PriceDetailChart({good, type, data, t, isMobile}: PriceDetailChartProps): React.ReactElement {
    const [timeRange, setTimeRange] = React.useState(isMobile? "7d" : "90d");

    let filteredData: Price[];

    if (type === SELL_CHART) {
        filteredData = data.sellOffer.filter((item: Price): boolean => {
            const date = new Date(item.createdAt);

            let daysToSubtract = 90;

            if (timeRange === "30d") {
                daysToSubtract = 30
            } else if (timeRange === "7d") {
                daysToSubtract = 7
            }

            const startDate = new Date(data.sellOffer.at(-1)?.createdAt!);

            startDate.setDate(startDate.getDate() - daysToSubtract);

            return date >= startDate
        });
    } else {
        filteredData = data.buyOffer.filter((item: Price): boolean => {
            const date = new Date(item.createdAt);

            let daysToSubtract = 90;

            if (timeRange === "30d") {
                daysToSubtract = 30
            } else if (timeRange === "7d") {
                daysToSubtract = 7
            }

            const startDate = new Date(data.buyOffer.at(-1)?.createdAt!);

            startDate.setDate(startDate.getDate() - daysToSubtract)

            return date >= startDate
        });
    }

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
                                dataKey="unitPrice"
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
                                dataKey="unitPrice"
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

function PriceDetailStatsCard({title, trend, change, value, info}: PriceDetailStatsCardProps) {
    return (
        <Card className="relative">
            <CardHeader className="pb-2">
                <CardTitle className="font-medium flex items-center justify-between">
                    <div className="flex items-center gap-2">
                        {title}
                        <Tooltip>
                            <TooltipTrigger asChild>
                                <CircleQuestionMark className="h-4 w-4" />
                            </TooltipTrigger>
                            <TooltipContent>
                                <p>{info}</p>
                            </TooltipContent>
                        </Tooltip>
                    </div>
                    {trend && (
                        <Badge className={`flex items-center gap-1 text-xs font-medium ${
                            trend === "up"
                                ? "text-green-600 dark:text-green-400"
                                : trend === "down"
                                    ? "text-red-600 dark:text-red-400"
                                    : "text-gray-600 dark:text-gray-400"
                        }`} variant="outline">
                            {trend === "up" ? <TrendingUp className="w-3 h-3" /> : <TrendingDown className="w-3 h-3" />}
                            {change && `${change > 0 ? "+" : ""}${change}%`}
                        </Badge>
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
                    <div className="flex items-center gap-4">
                        <div className="flex items-center gap-3">
                            <Avatar className="w-10 h-10">
                                <AvatarImage src={HoneycombGif} />
                            </Avatar>
                            <div>
                                <h1 className="text-3xl font-bold">Honeycomb Analytics</h1>
                                <p className="text-sm mt-1">
                                    Data series {formatDate(prices.sellOffer.at(0)!.createdAt)} - {formatDate(prices.sellOffer.at(-1)!.createdAt)}
                                </p>
                            </div>
                        </div>
                    </div>
                    <Badge
                        variant="secondary"
                        className="flex items-center gap-2 bg-green-500/10 text-green-400 border-green-500/30"
                    >
                        <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
                        Live Data
                    </Badge>
                </div>

                {/* Market Summary */}
                <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">

                    {/* Market Overview Card */}
                    <Card className="relative overflow-hidden">
                        <CardHeader>
                            <CardTitle className="relative flex items-center gap-2">
                                <ChartNoAxesCombined className="w-5 h-5" />
                                Market Overview
                            </CardTitle>
                        </CardHeader>
                        <CardContent className="relative space-y-4">
                            <div className="grid grid-cols-2 gap-4">
                                <div className="space-y-2">
                                    <div className="flex justify-between items-center">
                                        <span className="text-sm">Bid-Ask Spread:</span>
                                        <span className="font-medium">1000</span>
                                    </div>
                                    <div className="flex justify-between items-center">
                                        <span className="text-sm">Spread %:</span>
                                        <span className="font-medium">
                                            10%
                                        </span>
                                    </div>
                                </div>
                                <div className="space-y-2">
                                    <div className="flex justify-between items-center">
                                        <span className="text-sm">Market Cap:</span>
                                        <span className="font-medium">1250000</span>
                                    </div>
                                    <div className="flex justify-between items-center">
                                        <span className="text-sm">24h Volume:</span>
                                        <span className="font-medium">1250000</span>
                                    </div>
                                </div>
                            </div>

                            <div className="pt-2 border-t border-[var(--secondary)]">
                                <div className="flex items-center justify-between">
                                    <span className="text-sm">Market Status:</span>
                                    <div className="flex items-center gap-2">
                                        <Badge
                                            variant={"secondary"}
                                            className={`text-xs`}
                                        >
                                            {"Stable Trading"}
                                        </Badge>
                                    </div>
                                </div>
                            </div>
                        </CardContent>
                    </Card>

                    {/* Trading Insights Card */}
                    <Card>
                        <CardHeader>
                            <CardTitle className="flex items-center gap-2">
                                <TrendingUp className="w-5 h-5" />
                                Trading Insights
                            </CardTitle>
                        </CardHeader>
                        <CardContent className="space-y-4">
                            <div className="space-y-3">
                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--primary)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Buy Pressure</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className="w-3/4 h-full bg-[var(--primary)] rounded-full"></div>
                                        </div>
                                        <span className="text-xs text-[var(--primary)] font-medium w-16 text-right">Strong</span>
                                    </div>
                                </div>

                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--chart-theme-2)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Sell Pressure</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className="w-1/2 h-full bg-[var(--chart-theme-2)] rounded-full"></div>
                                        </div>
                                        <span className="text-xs text-[var(--chart-theme-2)] font-medium w-16 text-right">Moderate</span>
                                    </div>
                                </div>

                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--chart-2)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Liquidity</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className="w-5/6 h-full bg-[var(--chart-2)] rounded-full"></div>
                                        </div>
                                        <span className="text-xs text-[var(--chart-2)] font-medium w-16 text-right">High</span>
                                    </div>
                                </div>
                            </div>
                        </CardContent>
                    </Card>
                </div>

                {/* Market Activity Timeline */}
                <Card>
                    <CardHeader>
                        <CardTitle className="flex items-center gap-2">
                            <CalendarClock className="w-5 h-5" />
                            Recent Market Activity
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        <div className="space-y-3">
                            <div className="flex items-center gap-4 p-3 bg-[var(--secondary)] rounded-lg">
                                <div className="flex items-center gap-2">
                                    <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse"></div>
                                    <span className="text-xs">2 min ago</span>
                                </div>
                                <span className="text-sm">Large buy order executed at 1580</span>
                                <Badge variant="outline" className="text-green-400 border-green-500/30 text-xs">
                                    +2.1%
                                </Badge>
                            </div>

                            <div className="flex items-center gap-4 p-3 bg-[var(--secondary)] rounded-lg">
                                <div className="flex items-center gap-2">
                                    <div className="w-2 h-2 bg-blue-400 rounded-full"></div>
                                    <span className="text-xs">8 min ago</span>
                                </div>
                                <span className="text-sm text-gray-300">Market volatility decreased by 15%</span>
                                <Badge variant="outline" className="text-blue-400 border-blue-500/30 text-xs">
                                    Stable
                                </Badge>
                            </div>

                            <div className="flex items-center gap-4 p-3 bg-[var(--secondary)] rounded-lg">
                                <div className="flex items-center gap-2">
                                    <div className="w-2 h-2 bg-amber-400 rounded-full"></div>
                                    <span className="text-xs">15 min ago</span>
                                </div>
                                <span className="text-sm">Trading volume increased by 23%</span>
                                <Badge variant="outline" className="text-amber-400 border-amber-500/30 text-xs">
                                    Volume
                                </Badge>
                            </div>
                        </div>

                        <div className="mt-4 pt-4 flex items-center justify-between">
                            <span className="text-sm">Last data refresh:</span>
                            <div className="flex items-center gap-2">
                                <span className="text-sm">1 min ago</span>
                            </div>
                        </div>
                    </CardContent>
                </Card>

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
                            title={t.averagePrice}
                            value={statistics.buyOffersMean}
                            trend={"up"}
                            change={10}
                            info={t.averagePriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.medianPrice}
                            value={statistics.buyOffersMedian}
                            trend={"up"}
                            change={10}
                            info={t.medianPriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.stdDeviation}
                            value={statistics.buyOffersStdDeviation}
                            trend={"up"}
                            change={10}
                            info={t.stdDeviationInfo}
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
                            title={t.averagePrice}
                            value={statistics.sellOffersMean}
                            trend={"up"}
                            change={10}
                            info={t.averagePriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.medianPrice}
                            value={statistics.sellOffersMedian}
                            trend={"up"}
                            change={10}
                            info={t.medianPriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.stdDeviation}
                            value={statistics.sellOffersStdDeviation}
                            trend={"up"}
                            change={10}
                            info={t.stdDeviationInfo}
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