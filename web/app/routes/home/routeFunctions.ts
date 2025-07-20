import type {ApiResponse} from "~/lib/types";
import type {HomePageData, Price, PriceChartData} from "~/routes/home/types";


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

export function orderByPagePosition(unOrderedResults: ApiResponse<HomePageData>): HomePageData {
    if (!unOrderedResults.ok || !unOrderedResults.data) {
        return {};
    }

    const pricesMap: Map<number, string> = Object.keys(unOrderedResults.data)
        .reduce((result: Map<number, string>, goodName: string): Map<number, string> => {
            const price: PriceChartData = unOrderedResults.data![goodName];

            return result.set(price.pagePosition, goodName);
        }, new Map<number, string>());

    let result: HomePageData = {};

    for (let i: number = 0; i < Object.keys(unOrderedResults.data).length; i++) {
        const goodName: string = pricesMap.get(i + 1)!;

        result = {...result, [goodName]: unOrderedResults.data[goodName]};
    }

    return result;
}

export function getRelevantPrices(data: HomePageData): HomePageData {
    return Object.keys(data).reduce((acc: HomePageData, goodName: string): HomePageData => {
        const ticks: string[] = getTimeTicks(data[goodName].prices);

        const prices: Price[] = data[goodName].prices.filter((p: Price): boolean => ticks.includes(p.createdAt));

        acc[goodName] = {...data[goodName], prices};

        return acc;
    }, {})
}

function getTimeTicks(data: Price[], desiredTicks = 16): string[] {
    if (data.length <= desiredTicks) {
        return data.map((p: Price) => p.createdAt);
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