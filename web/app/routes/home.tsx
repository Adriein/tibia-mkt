import type {Route} from "./+types/home";
import {fetchPrices, getRelevantPrices, orderByPagePosition} from "~/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData} from "~/home/types";
import {PriceOverview} from "~/components/ui/price-overview";
import {English, type HomeTranslations, loc} from "~/locale/loc";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Tibia Mkt" },
    { name: "description", content: "Welcome to Tibia Mkt!" },
  ];
}

export async function loader(): Promise<{data: HomePageData, t: HomeTranslations}> {
  const prices: ApiResponse<HomePageData> = await fetchPrices();

  const orderedPrices: HomePageData = orderByPagePosition(prices);

  return {data: getRelevantPrices(orderedPrices), t: loc(English, "Home")};
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
