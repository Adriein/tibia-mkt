import {CogChart, HomePageData} from "~/shared/types";
import {HomeResponse} from "~/routes/_index/types";

const fetchHomeData = async (): Promise<HomeResponse<HomePageData>> => {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/home?" +
        "item=tibiaCoin&item=honeycomb&item=swamplingWood&item=brokenShamanicStaff"
    );

    const response = await fetch(homeRequest);
    return await response.json();
}

const fetchSearchInputGoods = async (): Promise<HomeResponse<string[]>> => {
    const searchGoodRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/goods`
    );

    const response = await fetch(searchGoodRequest);
    return await response.json();
}

export const fetchData = async (): Promise<[HomeResponse<HomePageData>, HomeResponse<string[]>]> => {
    return await Promise.all([fetchHomeData(), fetchSearchInputGoods()]);
}

export const orderHomePageData = (unOrderedResults: HomeResponse<HomePageData>): HomePageData => {
    const cogOrderMap: Map<number, string> = Object.keys(unOrderedResults.data)
        .reduce((result: Map<number, string>, cogName: string) => {
            const cog: CogChart = unOrderedResults.data[cogName];

            return result.set(cog.pagePosition, cogName);
        }, new Map<number, string>());

    let result: HomePageData = {};

    for (let i: number = 0; i < Object.keys(unOrderedResults.data).length; i++) {
        const cogName: string = cogOrderMap.get(i + 1)!;

        result = {...result, [cogName]: unOrderedResults.data[cogName]};
    }

    return result;
}