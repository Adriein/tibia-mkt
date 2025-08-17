import type {ApiResponse, Price, PriceChartData} from "~/lib/types";
import type {HomePageData, HomePagePriceDataPoint, MergedHomePageData} from "~/routes/home/types";


export async function fetchPrices(): Promise<ApiResponse<HomePageData>> {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/prices?" +
        "world=Secura&" +
        "good=tibiaCoin&good=honeycomb&good=swamplingWood&good=brokenShamanicStaff"
    );

    const response = await fetch(homeRequest);
    return await response.json();
}

export function orderByPagePosition(unOrderedResults: HomePageData): HomePageData {
    const pricesMap: Map<number, string> = Object.keys(unOrderedResults)
        .reduce((result: Map<number, string>, goodName: string): Map<number, string> => {
            const price: PriceChartData = unOrderedResults[goodName];

            return result.set(price.pagePosition, goodName);
        }, new Map<number, string>());

    let result: HomePageData = {};

    for (let i: number = 0; i < Object.keys(unOrderedResults).length; i++) {
        const goodName: string = pricesMap.get(i + 1)!;

        result = {...result, [goodName]: unOrderedResults[goodName]};
    }

    return result;
}

export function mergeSellAndBuyOffers(orderedResults: HomePageData): MergedHomePageData {
    return Object.keys(orderedResults).reduce((acc: MergedHomePageData, good: string): MergedHomePageData => {
        const chartData: PriceChartData = orderedResults[good];

        const specificGoodMergedPrices: HomePagePriceDataPoint[] = chartData
            .buyOffer
            .map((bo: Price): HomePagePriceDataPoint => {
                const sellOffer: Price|undefined = chartData
                    .sellOffer
                    .find((so: Price): boolean => so.createdAt === bo.createdAt);

                if (!sellOffer) {
                    throw new Error();
                }

                return {
                    sellPrice: sellOffer.unitPrice,
                    buyPrice: bo.unitPrice,
                    createdAt: bo.createdAt,
                    world: bo.world
                }
            });

        return {
            ...acc,
            [good]: {
                wikiLink: chartData.wikiLink,
                dataPoints: specificGoodMergedPrices,
                pagePosition: chartData.pagePosition
            }
        };
    }, {} as MergedHomePageData)
}

export function getRelevantPrices(data: MergedHomePageData): MergedHomePageData {
    return Object.keys(data).reduce((acc: MergedHomePageData, goodName: string): MergedHomePageData => {
        const ticks: string[] = getTimeTicks(data[goodName].dataPoints);

        const prices: HomePagePriceDataPoint[] = data[goodName]
            .dataPoints
            .filter((p: HomePagePriceDataPoint): boolean => ticks.includes(p.createdAt));

        acc[goodName] = {...data[goodName], dataPoints: prices};

        return acc;
    }, {})
}

function getTimeTicks(data: HomePagePriceDataPoint[], desiredTicks = 16): string[] {
    if (data.length <= desiredTicks) {
        return data.map((p: HomePagePriceDataPoint) => p.createdAt);
    }

    const start = new Date(data[0].createdAt);
    const end = new Date(data[data.length - 1].createdAt);

    const totalTime: number = end.getTime() - start.getTime();
    const interval: number = totalTime / (desiredTicks - 1);

    const ticks: string[] = [];
    for (let i: number = 0; i < desiredTicks; i++) {
        const date = new Date(start.getTime() + i * interval);

        const year: number = date.getFullYear();
        const month: string = String(date.getMonth() + 1).padStart(2, '0');
        const day: string = String(date.getDate()).padStart(2, '0');

        ticks.push(`${year}-${month}-${day}`);
    }

    return ticks;
}