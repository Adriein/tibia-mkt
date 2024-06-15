export type TibiaCoinCog = { buyPrice: number, sellPrice: number, date: string, world: string };

export type ChartMetadata = {  yAxisTick: number[], xAxisTick: string[] }

export type HomePageData = { cogs: TibiaCoinCog[], chartMetadata: ChartMetadata }
