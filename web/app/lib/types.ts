export type ApiResponse<T> = {ok: boolean, data?: T, error?: string}

export type Price = { unitPrice: number, amount: number, createdAt: string, world: string };

export type PriceChartData = { wikiLink: string, buyOffer: Price[], sellOffer: Price[], pagePosition: number }