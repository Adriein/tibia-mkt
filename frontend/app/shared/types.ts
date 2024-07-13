export type Cog = { buyPrice: number, sellPrice: number, date: string, world: string };

export type YAxisTick = { price: number, date: string }

export type ChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[] }

export type CogChart = { wiki: string, cog: Cog[], chartMetadata: ChartMetadata, pagePosition: number }

export type HomePageData = { [key: string]: CogChart }
