import React from "react";
import type {Route} from "@/.react-router/types/app/routes/detail/+types/detail";
import type {HomePageData} from "~/routes/home/types";
import {English, type HomeTranslations, loc} from "~/locale/loc";
import type {ApiResponse} from "~/lib/types";
import {fetchPrices, getRelevantPrices, orderByPagePosition} from "~/routes/home/routeFunctions";
import {beautifyCamelCase} from "~/lib/utils";

export function meta({params}: Route.MetaArgs) {
    return [
        { title: beautifyCamelCase(params.good) },
        { name: "description", content: "Detail" },
    ];
}

export async function loader(): Promise<{data: HomePageData, t: HomeTranslations}> {
    const prices: ApiResponse<HomePageData> = await fetchPrices();

    const orderedPrices: HomePageData = orderByPagePosition(prices);

    return {data: getRelevantPrices(orderedPrices), t: loc(English, "Home")};
}

export default function PriceDetail({loaderData}: Route.ComponentProps): React.ReactElement {
    return (
        <div>Detail</div>
    );
}