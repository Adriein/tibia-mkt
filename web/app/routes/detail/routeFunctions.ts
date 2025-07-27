import type {ApiResponse} from "~/lib/types";
import type {DetailPageData, DetailPagePricesData, DetailPageStatisticsData} from "~/routes/detail/types";

export async function fetchDetailData(good: string): Promise<ApiResponse<DetailPageData>> {
    const [pricesRes, statRes] = await Promise.all([
        fetchGoodPrices(good),
        fetchStatisticsData(good)
    ]);

    if (!pricesRes.ok || !pricesRes.data || !statRes.ok || !statRes.data) {
        return {
            ok: false,
            error: "Error while fetching data",
        }
    }

    return {
        ok: true,
        data: {
            prices: pricesRes.data,
            statistics: statRes.data
        }
    }
}

async function fetchGoodPrices(good: string): Promise<ApiResponse<DetailPagePricesData>> {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/prices?" +
        "world=Secura&" +
        `good=${good}`
    );

    const response = await fetch(homeRequest);
    return await response.json();
}

async function fetchStatisticsData(good: string): Promise<ApiResponse<DetailPageStatisticsData>> {
    const req: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/details?" +
        "world=Secura&" +
        `good=${good}`
    );

    const response = await fetch(req);

    return await response.json();
}