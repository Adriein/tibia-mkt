import React from "react";
import type {Route} from "@/.react-router/types/app/routes/detail/+types/detail";
import {type DetailTranslations, English, loc} from "~/locale/loc";
import type {ApiResponse, PriceChartData} from "~/lib/types";
import {beautifyCamelCase} from "~/lib/utils";
import {fetchDetailData} from "~/routes/detail/routeFunctions";
import type {DetailPageData, DetailPageEventsData, DetailPageStatisticsData} from "~/routes/detail/types";
import {redirect} from "react-router";
import {PriceDetail} from "~/components/ui/price-detail";

type LoaderData = {
    t: DetailTranslations;
    prices: PriceChartData;
    statistics: DetailPageStatisticsData;
    events: DetailPageEventsData[];
    isMobile: boolean
};

export function meta({params}: Route.MetaArgs) {
    return [
        { title: `Tibia Mkt | ${beautifyCamelCase(params.good)}` },
        { name: "description", content: "Detail" },
    ];
}

export async function loader({params, request}: Route.LoaderArgs): Promise<LoaderData | Response> {
    const userAgent: string = request.headers.get('user-agent') || '';
    const isMobile: boolean = /Mobi|Android|IPhone/i.test(userAgent);

    const res: ApiResponse<DetailPageData> = await fetchDetailData(params.good);

    if (!res.ok || !res.data) {
        return redirect("/404");
    }

    return {
        prices: res.data.prices[params.good],
        statistics: res.data.statistics,
        events: res.data.events,
        t: loc(English, "Detail"),
        isMobile,
    };
}

export default function Detail({loaderData, params}: Route.ComponentProps): React.ReactElement {
    const { prices, statistics, events, t, isMobile } = loaderData;
    return (
        <main>
            <PriceDetail
                good={params.good}
                prices={prices}
                statistics={statistics}
                events={events}
                t={t}
                isMobile={isMobile}
            />
        </main>
    );
}