import type {PriceChartData} from "~/lib/types";

export type PricesHomePageData = { [key: string]: PriceChartData }

export type HomePagePriceDataPoint = {
    sellPrice: number,
    buyPrice: number,
    createdAt: string,
    world: string
}

export type HomePriceChartData = { wikiLink: string, dataPoints: HomePagePriceDataPoint[], pagePosition: number }

export type MergedHomePageData = {[key: string]: HomePriceChartData}

export type HomePageData = {prices: MergedHomePageData, news: TibiaArticleData[]}

export type TibiaArticleRes = {
    category: string,
    date: string,
    id: number,
    news: string,
    type: string,
    url: string,
    url_api: string
};

export type LatestTibiaNewsRes = {
    news: TibiaArticleRes[]
};

export type TibiaArticleData = {
    title: string,
    date: string,
    category: string,
};