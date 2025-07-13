import type {ApiResponse} from "~/lib/types";
import type {HomePageData, Price, PriceChartData} from "~/home/types";


export const fetchPrices: () => Promise<ApiResponse<HomePageData>> = async (): Promise<ApiResponse<HomePageData>> => {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/prices?" +
        "world=Secura&" +
        "good=tibiaCoin&good=honeycomb&good=swamplingWood&good=brokenShamanicStaff"
    );

    const response = await fetch(homeRequest);
    return await response.json();
}

export const orderHomePageData = (unOrderedResults: ApiResponse<HomePageData>): HomePageData => {
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

export const getLast6MonthsPreview = (chartData: PriceChartData): PriceChartData => {
    const initialDataPoint = chartData.prices.length - (30 * 3);
    return {
        ...chartData,
        prices: chartData.prices.slice(initialDataPoint, chartData.prices.length),
    }
}