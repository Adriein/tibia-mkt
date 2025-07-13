import type {Route} from "./+types/home";
import {fetchPrices, getLast6MonthsPreview, orderHomePageData} from "~/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData} from "~/home/types";
import {PriceOverview} from "~/components/ui/price-overview";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Tibia Mkt" },
    { name: "description", content: "Welcome to Tibia Mkt!" },
  ];
}

export async function loader(): Promise<{data: HomePageData}> {
  const prices: ApiResponse<HomePageData> = await fetchPrices();

  return {data: orderHomePageData(prices)};
}

export default function Home({loaderData}: Route.ComponentProps) {
  const { data } = loaderData;

  return (
      <main className="flex flex-col items-center w-screen h-screen">
        <header className="flex items-center justify-center w-full h-24 text-center text-2xl font-bold text-foreground">
          <h1>Welcome to Tibia Mkt</h1>
        </header>
        <div className="flex w-full border-amber-600 border-2">
          <PriceOverview good={"honeycomb"} data={getLast6MonthsPreview(data.honeycomb)}/>
        </div>
      </main>
  );
}
