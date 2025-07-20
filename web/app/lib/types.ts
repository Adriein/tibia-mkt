export type ApiResponse<T> = {ok: boolean, data?: T, error?: string}

export type YAxisTick = { price: number, createdAt: string }

export type ChartMetadata = {  yAxisTick: YAxisTick[], xAxisTick: string[] }

export type Price = { buyPrice: number, sellPrice: number, createdAt: string, world: string };

export type PriceChartData = { wikiLink: string, prices: Price[], chartMetadata: ChartMetadata, pagePosition: number }