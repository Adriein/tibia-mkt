export type Cog = { buyPrice: number, sellPrice: number, date: string, world: string };

export type ChartMetadata = {  yAxisTick: number[], xAxisTick: string[] }

export type CogChart = { wiki: string, cog: Cog[], chartMetadata: ChartMetadata }

export type HomePageData = { [key: string]: CogChart }
