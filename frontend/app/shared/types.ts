export type Cog = { buyPrice: number, sellPrice: number, date: string, world: string };

export type YAxisTick = { price: number, date: string }

export type ReferenceLine = { offerType: string, average: number }

export type ChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[] }

export type DetailChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[], referenceLine: ReferenceLine }

export type CogChart = { wiki: string, cog: Cog[], chartMetadata: ChartMetadata, pagePosition: number }

export type HomePageData = { [key: string]: CogChart }

export type DetailPageData = {
    wiki: string,
    cog: Cog[],
    sellOfferChart: DetailChartMetadata,
    buyOfferChart: DetailChartMetadata
}

export type RemixMetaFunc = {params: { item: string }};
