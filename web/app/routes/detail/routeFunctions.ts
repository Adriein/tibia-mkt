import type {ApiResponse} from "~/lib/types";
import type {DetailPageData} from "~/routes/detail/types";

export async function fetchGoodPrices(good: string): Promise<ApiResponse<DetailPageData>> {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/prices?" +
        "world=Secura&" +
        `good=${good}`
    );

    const response = await fetch(homeRequest);
    return await response.json();
}