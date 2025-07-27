import type {PriceChartData} from "~/lib/types";

export type DetailPageData = { prices: DetailPagePricesData, statistics: DetailPageStatisticsData };

export type DetailPagePricesData = { [key: string]: PriceChartData } & { [key: string]: any};

export type DetailPageStatisticsData = {
    sellOffersMean: number;
    sellOffersStdDeviation: number;
    sellOffersMedian: number;
    buyOffersMean: number;
    buyOffersStdDeviation: number;
    buyOffersMedian: number;
};
