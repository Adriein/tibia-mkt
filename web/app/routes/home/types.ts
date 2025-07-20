export type YAxisTick = { price: number, createdAt: string }

export type ChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[] }

export type Price = { buyPrice: number, sellPrice: number, createdAt: string, world: string };

export type PriceChartData = { wikiLink: string, prices: Price[], chartMetadata: ChartMetadata, pagePosition: number }

export type HomePageData = { [key: string]: PriceChartData }