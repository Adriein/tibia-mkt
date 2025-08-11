import type {PriceChartData} from "~/lib/types";

export type DetailPageData = {
    prices: DetailPagePricesData,
    statistics: DetailPageStatisticsData,
    events: DetailPageEventsData[]
};

export type Stats = {
    sellOffersMean: number;
    sellOffersStdDeviation: number;
    sellOffersMedian: number;
    buyOffersMean: number;
    buyOffersStdDeviation: number;
    buyOffersMedian: number;
}

export type Overview = {
    buySellSpread: number;
    spreadPercentage: number;
    marketCap: number;
    lastTwentyFourHoursVolume: number;
    marketVolumePercentageTendency: number;
    totalGoodsBeingSold: number;
}

export type Insights = {
    marketStatus: string;
    marketType: string;
    buyPressure: number;
    sellPressure: number;
    liquidity: number;
}

export type DetailPageEventsData = {
    id: number;
    name: string;
    goodName: string;
    world: string;
    description: string;
    occurredAt: string;
}

export type DetailPagePricesData = { [key: string]: PriceChartData } & { [key: string]: any};

export type DetailPageStatisticsData = {
    stats: Stats;
    overview: Overview;
    insights: Insights;
};
