import React from "react";
import type {Route} from "@/.react-router/types/app/routes/detail/+types/detail";
import {type DetailTranslations, English, loc} from "~/locale/loc";
import type {ApiResponse, PriceChartData} from "~/lib/types";
import {beautifyCamelCase} from "~/lib/utils";
import {fetchDetailData} from "~/routes/detail/routeFunctions";
import type {DetailPageData, DetailPageStatisticsData} from "~/routes/detail/types";
import {redirect} from "react-router";
import {PriceDetail} from "~/components/ui/price-detail";

type LoaderData = { t: DetailTranslations; prices: PriceChartData, statistics: DetailPageStatisticsData };

export function meta({params}: Route.MetaArgs) {
    return [
        { title: `Tibia Mkt | ${beautifyCamelCase(params.good)}` },
        { name: "description", content: "Detail" },
    ];
}

export async function loader({params}: Route.LoaderArgs): Promise<LoaderData | Response> {
    const res: ApiResponse<DetailPageData> = await fetchDetailData(params.good);

    if (!res.ok || !res.data) {
        return redirect("/404");
    }

    return {
        prices: res.data.prices[params.good],
        statistics: res.data.statistics,
        t: loc(English, "Detail")
    };
}

export default function Detail({loaderData, params}: Route.ComponentProps): React.ReactElement {
    const { prices, statistics, t } = loaderData;
    return (
        <main className="flex flex-col items-center w-screen h-screen">
            <div className="flex w-full">
                <PriceDetail good={params.good} data={prices} t={t}/>
            </div>
        </main>
    );
}