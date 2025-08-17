import {fetchPrices, getRelevantPrices, mergeSellAndBuyOffers, orderByPagePosition} from "~/routes/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData, MergedHomePageData} from "~/routes/home/types";
import {PriceOverview} from "~/components/ui/price-overview";
import {English, type HomeTranslations, loc} from "~/locale/loc";
import type {Route} from "@/.react-router/types/app/routes/home/+types/home";
import React from "react";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Tibia Mkt" },
    { name: "description", content: "Welcome to Tibia Mkt!" },
  ];
}

export async function loader(): Promise<{data: MergedHomePageData, t: HomeTranslations}> {
  const prices: ApiResponse<HomePageData> = await fetchPrices();

    if (!prices.ok || !prices.data) {
        return {data: {}, t: loc(English, "Home")};
    }

    const orderedPrices: HomePageData = orderByPagePosition(prices.data);

    const results: MergedHomePageData = mergeSellAndBuyOffers(orderedPrices);

  return {data: getRelevantPrices(results), t: loc(English, "Home")};
}

export default function Home({loaderData}: Route.ComponentProps) {
  const { data, t } = loaderData;

  return (
      <main className="flex flex-col items-center w-screen h-screen">
        <header className="flex items-center justify-center w-full h-24 text-center text-2xl font-bold text-foreground">
          <h1>{t.welcome}</h1>
        </header>
        <div className="flex w-full">
          <PriceOverview good={"honeycomb"} data={data.honeycomb}/>
        </div>
      </main>
  );
}
