import {Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {CartesianGrid, Line, LineChart, XAxis} from "recharts";
import {beautifyCamelCase, formatDateToElegantForm, formatDateToShortForm} from "~/lib/utils";
import type {NameType, Payload, ValueType} from "recharts/types/component/DefaultTooltipContent";
import React from "react";
import type {Price, PriceChartData} from "~/lib/types";
import {ToggleGroup, ToggleGroupItem} from "~/components/ui/toggle-group";
import type {DetailTranslations} from "~/locale/loc";
import type {DetailPageEventsData, DetailPageStatisticsData} from "~/routes/detail/types";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {Badge} from "~/components/ui/badge";
import {TrendingUp, TrendingDown, CircleQuestionMark, ChartNoAxesCombined, CalendarClock, TriangleAlert, Scale} from "lucide-react";
import {Avatar, AvatarImage} from "~/components/ui/avatar";
import HoneycombGif from "~/assets/honeycomb.gif";
import {Tooltip, TooltipContent, TooltipTrigger} from "~/components/ui/tooltip";

const SELL_CHART = "sell";
const BUY_CHART = "buy";

const chartConfig = {
    unitPrice: {
        label: "Price",
    },
} satisfies ChartConfig

type PriceDetailProps = {
    good: string;
    prices: PriceChartData;
    statistics: DetailPageStatisticsData;
    events: DetailPageEventsData[];
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

const MARKET_STATUS_COLORS = {
    "Stable": "bg-green-500/10 text-green-400 border-green-500/20",
    "Risky": "bg-amber-500/10 text-amber-400 border-amber-500/20",
    "Volatile": "bg-red-500/10 text-red-400 border-red-500/20",
} as Record<string, string>;

const getStatusBarWith = (percentage: number): string => {
    if (percentage >= 95) return 'w-full';
    if (percentage >= 80) return 'w-4/5'; // or w-[80%]
    if (percentage >= 75) return 'w-3/4';
    if (percentage >= 60) return 'w-3/5'; // or w-[60%]
    if (percentage >= 50) return 'w-1/2';
    if (percentage >= 40) return 'w-2/5'; // or w-[40%]
    if (percentage >= 25) return 'w-1/4';
    if (percentage >= 20) return 'w-1/5'; // or w-[20%]
    if (percentage > 0) return 'w-1/12'; // Smallest non-zero, or custom like w-[5%]
    return 'w-0'; // For 0%
};

const marketPressure = (percentage: number): string => {
    if (percentage >= 80) return 'Strong';
    if (percentage >= 60) return 'High';
    if (percentage >= 40) return 'Moderate';
    return 'Low'
}

const getMarketTendencyIcon = (tendency: string) => {
    if (tendency.includes("Exhaustion") || tendency.includes("Pullback")) {
        return (
            <Badge
                variant="outline"
                className="text-amber-400 border-amber-500/40 bg-amber-500/10 text-sm font-medium px-3 py-1"
            >
                <TriangleAlert className="w-3 h-3 mr-1"/>
                {tendency}
            </Badge>
        );
    }

    if (tendency.includes("Bull")) {
        return (
            <Badge
            variant="outline"
            className="text-green-400 border-green-500/40 bg-green-500/10 text-sm font-medium px-3 py-1"
            >
                <TrendingUp className="w-3 h-3 mr-1"/>
                {tendency}
            </Badge>
        );
    }

    if (tendency.includes("Bear")) {
        return (
            <Badge
                variant="outline"
                className="text-red-400 border-red-500/40 bg-red-500/10 text-sm font-medium px-3 py-1"
            >
                <TrendingDown className="w-3 h-3 mr-1"/>
                {tendency}
            </Badge>
        );
    }

    return (
        <Badge
            variant="outline"
            className="text-blue-400 border-blue-500/40 bg-blue-500/10 text-sm font-medium px-3 py-1"
        >
            <Scale className="w-3 h-3 mr-1"/>
            {tendency}
        </Badge>
    );
}

const labelFormatter = (label: string, _: Array<Payload<ValueType, NameType>>): React.ReactNode => {
    return <span>{formatDateToShortForm(label)}</span>
}

const transformNumberToLocale = (value: number|string): string => {
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
                            tickFormatter={formatDateToShortForm}
                        />
                        <ChartTooltip
                            cursor={false}
                            content={
                                <ChartTooltipContent
                                    indicator="line"
                                    labelFormatter={labelFormatter}
                                    valueFormatter={transformNumberToLocale}
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

function PriceDetailStatsCard({title, value, info}: PriceDetailStatsCardProps) {
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
                </CardTitle>
            </CardHeader>
            <CardContent>
                <div className="text-2xl font-bold">{transformNumberToLocale(value)}</div>
            </CardContent>
        </Card>
    );
}


function PriceDetail({good, prices, statistics, events, t, isMobile}: PriceDetailProps) {
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
                                <h1 className="text-3xl font-bold">Honeycomb Detail</h1>
                                <p className="text-sm mt-1">
                                    Data series from {formatDateToElegantForm(prices.sellOffer.at(0)!.createdAt)} to {formatDateToElegantForm(prices.sellOffer.at(-1)!.createdAt)}
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
                        <CardContent className="relative space-y-6 flex flex-col flex-1">
                            {/* Key Metrics Row */}
                            <div className="grid grid-cols-2 gap-6">
                                <div className="space-y-1">
                                    <p className="text-xs uppercase tracking-wide">Bid-Ask Spread</p>
                                    <p className="text-xl font-bold">
                                        {transformNumberToLocale(statistics.overview.buySellSpread)}
                                    </p>
                                    <p className="text-xs text-[var(--primary)] font-medium">
                                        {statistics.overview.spreadPercentage}% of sell price
                                    </p>
                                </div>
                                <div className="space-y-1">
                                    <p className="text-xs uppercase tracking-wide">24h Volume</p>
                                    <p className="text-xl font-bold">
                                        {transformNumberToLocale(statistics.overview.lastTwentyFourHoursVolume)}
                                    </p>
                                    <p className="text-xs text-[var(--chart-theme-2)] font-medium">
                                        {statistics.overview.marketVolumePercentageTendency}% from yesterday
                                    </p>
                                </div>
                            </div>

                            {/* Secondary Metrics */}
                            <div className="grid grid-cols-2 gap-6 pt-4 border-t border-[var(--secondary)]">
                                <div className="flex justify-between items-center">
                                    <span className="text-sm">Market Cap</span>
                                    <span className="font-semibold">
                                        {transformNumberToLocale(statistics.overview.marketCap)}
                                    </span>
                                </div>
                                <div className="flex justify-between items-center">
                                    <span className="text-sm">Goods in market</span>
                                    <span className="font-semibold">{statistics.overview.totalGoodsBeingSold}</span>
                                </div>
                            </div>

                            {/* Market Status */}
                            <div className="flex flex-1 items-center justify-between pt-4 border-t border-[var(--secondary)]">
                                <div className="flex items-center gap-2">
                                    <span className="text-sm">Market Status</span>
                                </div>
                                <Badge
                                    variant={"secondary"}
                                    className={`text-xs ${MARKET_STATUS_COLORS[statistics.insights.marketStatus]}`}
                                >
                                    {`${statistics.insights.marketStatus} Trading`}
                                </Badge>
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
                            {/* Market Tendency - Featured Row */}
                            <div className="p-2 bg-gradient-to-r from-[var(--card)] to-[color-mix(in_oklab,var(--card) 90%, var(--primary) 10%)] rounded-lg border border-[var(--border)]">
                                <div className="flex items-center justify-between">
                                    <div className="flex items-center gap-3">
                                        <div>
                                            <span className="text-sm font-medium">Market Tendency</span>
                                            <p className="text-xs text-gray-300 mt-0.5">Overall market direction</p>
                                        </div>
                                    </div>
                                    <div className="flex items-center gap-3">
                                        {getMarketTendencyIcon(statistics.insights.marketType)}
                                    </div>
                                </div>
                            </div>
                            <div className="space-y-3">
                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--primary)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Buy Pressure</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className={`${getStatusBarWith(statistics.insights.buyPressure)} h-full bg-[var(--primary)] rounded-full`}/>
                                        </div>
                                        <span className="text-xs text-[var(--primary)] font-medium w-16 text-right">
                                            {marketPressure(statistics.insights.buyPressure)}
                                        </span>
                                    </div>
                                </div>

                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--chart-theme-2)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Sell Pressure</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className={`${getStatusBarWith(statistics.insights.sellPressure)} h-full bg-[var(--chart-theme-2)] rounded-full`}/>
                                        </div>
                                        <span className="text-xs text-[var(--chart-theme-2)] font-medium w-16 text-right">
                                            {marketPressure(statistics.insights.sellPressure)}
                                        </span>
                                    </div>
                                </div>

                                <div className="flex items-center justify-between p-3 bg-[var(--secondary)] rounded-lg">
                                    <div className="flex items-center gap-2 min-w-0 flex-1">
                                        <div className="w-2 h-2 bg-[var(--chart-2)] rounded-full flex-shrink-0"></div>
                                        <span className="text-sm">Liquidity</span>
                                    </div>
                                    <div className="flex items-center gap-2 min-w-[100px]">
                                        <div className="w-16 h-2 bg-[var(--foreground)] rounded-full overflow-hidden">
                                            <div className={`${getStatusBarWith(statistics.insights.liquidity)} h-full bg-[var(--chart-2)] rounded-full`}></div>
                                        </div>
                                        <span className="text-xs text-[var(--chart-2)] font-medium w-16 text-right">
                                            {marketPressure(statistics.insights.liquidity)}
                                        </span>
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
                            {events.map((event: DetailPageEventsData) => {
                                return (
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
                                );
                            })}
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

                        <div className="mt-4 pt-4 flex items-center justify-between border-t border-[var(--secondary)]">
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
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <PriceDetailStatsCard
                            title={t.averagePrice}
                            value={statistics.stats.buyOffersMean}
                            info={t.averagePriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.medianPrice}
                            value={statistics.stats.buyOffersMedian}
                            info={t.medianPriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.stdDeviation}
                            value={statistics.stats.buyOffersStdDeviation}
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
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                        <PriceDetailStatsCard
                            title={t.averagePrice}
                            value={statistics.stats.sellOffersMean}
                            info={t.averagePriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.medianPrice}
                            value={statistics.stats.sellOffersMedian}
                            info={t.medianPriceInfo}
                        />
                        <PriceDetailStatsCard
                            title={t.stdDeviation}
                            value={statistics.stats.sellOffersStdDeviation}
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