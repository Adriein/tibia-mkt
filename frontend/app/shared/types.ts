export type Cog = { buyPrice: number, sellPrice: number, date: string, world: string };

export type YAxisTick = { price: number, date: string }

export type ReferenceLine = { offerType: string, average: number }

export type ChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[] }

export type DetailChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[], referenceLine: ReferenceLine }

export type CogChart = { wiki: string, cogs: Cog[], chartMetadata: ChartMetadata, pagePosition: number }

export type HomePageData = { [key: string]: CogChart }

export type DetailCreature = {name: string, dropRate: number, killStatistic: number}

export type DetailPageData = {
    wiki: string,
    creatures: DetailCreature[],
    sellOfferHistoricData: SellOfferHistoricData[],
    sellOfferProbability: SellOfferProbability,
    cogs: Cog[],
    sellOfferChart: DetailChartMetadata,
    buyOfferChart: DetailChartMetadata
}

export type RemixMetaFunc = {params: { item: string }};

export type SellOfferFrequency = { range: string, occurrences: number, frequency: number }

export type SellOfferProbability = { median: number, mean: number, stdDeviation: number, sellOfferFrequency: SellOfferFrequency[]}

export type SellOfferHistoricData = { id: string, cog: string,  mean: number, stdDeviation: number, createdAt: string}