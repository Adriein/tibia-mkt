import type {PriceChartData} from "~/lib/types";

export type HomePageData = { [key: string]: PriceChartData }

export type HomePagePriceDataPoint = {
    sellPrice: number,
    buyPrice: number,
    createdAt: string,
    world: string
}

export type HomePriceChartData = { wikiLink: string, dataPoints: HomePagePriceDataPoint[], pagePosition: number }

export type MergedHomePageData = {[key: string]: HomePriceChartData}