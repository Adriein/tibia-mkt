import React from "react";
import type {Route} from "@/.react-router/types/app/routes/detail/+types/detail";
import {English, type HomeTranslations, loc} from "~/locale/loc";
import type {ApiResponse} from "~/lib/types";
import {beautifyCamelCase} from "~/lib/utils";
import {fetchGoodPrices} from "~/routes/detail/routeFunctions";
import type {DetailPageData} from "~/routes/detail/types";
import {redirect} from "react-router";
import {PriceDetail} from "~/components/ui/price-detail";

export function meta({params}: Route.MetaArgs) {
    return [
        { title: beautifyCamelCase(params.good) },
        { name: "description", content: "Detail" },
    ];
}

export async function loader({params}: Route.LoaderArgs): Promise<{ data: any; t: HomeTranslations } | Response> {
    const prices: ApiResponse<DetailPageData> = await fetchGoodPrices(params.good);

    if (!prices.ok || !prices.data) {
        return redirect("/404");
    }

    return {data: prices.data[params.good], t: loc(English, "Home")};
}

export default function Detail({loaderData, params}: Route.ComponentProps): React.ReactElement {
    const { data, t } = loaderData;
    return (
        <main className="flex flex-col items-center w-screen h-screen">
            <div className="flex w-full">
                <PriceDetail good={params.good} data={data}/>
            </div>
        </main>
    );
}