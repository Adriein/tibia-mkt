import React from "react";
import type {Route} from "@/.react-router/types/app/routes/detail/+types/detail";
import {BeautyLocale, type DetailTranslations, languageConverter, loc} from "~/locale/loc";
import type {ApiResponse, PriceChartData} from "~/lib/types";
import {camelCaseToTitle} from "~/lib/utils";
import {fetchDetailData} from "~/routes/detail/routeFunctions";
import type {DetailPageData, DetailPageEventsData, DetailPageStatisticsData} from "~/routes/detail/types";
import {redirect} from "react-router";
import {PriceDetail} from "~/components/ui/price-detail";

type LoaderData = {
    t: DetailTranslations;
    prices: PriceChartData;
    statistics: DetailPageStatisticsData;
    events: DetailPageEventsData[];
    isMobile: boolean;
    loc: string;
};

export function meta({params}: Route.MetaArgs) {
    return [
        { title: `Tibia Mkt | ${camelCaseToTitle(params.good)}` },
        { name: "description", content: "Detail" },
    ];
}

export async function loader({params, request}: Route.LoaderArgs): Promise<LoaderData | Response> {
    const url = new URL(request.url);
    const language: string = url.searchParams.get('lang') || BeautyLocale.English;

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
        t: loc(languageConverter(language), "Detail"),
        loc: language,
        isMobile,
    };
}

export default function Detail({loaderData, params}: Route.ComponentProps): React.ReactElement {
    const { prices, statistics, events, t, isMobile, loc } = loaderData;
    return (
        <main>
            <PriceDetail
                good={params.good}
                prices={prices}
                statistics={statistics}
                events={events}
                t={t}
                isMobile={isMobile}
                loc={loc}
            />
        </main>
    );
}