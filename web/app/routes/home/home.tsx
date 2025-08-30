import {
    fetchPrices,
    fetchTibiaNews,
    getRelevantPrices,
    mergeSellAndBuyOffers,
    orderByPagePosition
} from "~/routes/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {PricesHomePageData, MergedHomePageData, HomePageData} from "~/routes/home/types";
import {PriceOverview} from "~/components/ui/price-overview";
import {BeautyLocale, type HomeTranslations, loc} from "~/locale/loc";
import type {Route} from "@/.react-router/types/app/routes/home/+types/home";
import React from "react";
import {Server} from "lucide-react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {HomeHeader} from "~/components/ui/home-header";
import {HomeTibiaNews} from "~/components/ui/home-news";

export function meta({}: Route.MetaArgs) {
    return [
        { title: "Tibia Mkt" },
        { name: "description", content: "Welcome to Tibia Mkt!" },
    ];
}

export async function loader({ request }: Route.LoaderArgs): Promise<{data: HomePageData, t: HomeTranslations}> {
    const url = new URL(request.url);
    const language: string = url.searchParams.get('lang') || BeautyLocale.English;

    const [news, prices] = await Promise.all([await fetchTibiaNews(), await fetchPrices()]);

    //const prices: ApiResponse<HomePageData> = await fetchPrices();

    if (!prices.ok || !prices.data) {
        return {data: {news, prices: {}}, t: loc(language, "Home")};
    }

    const orderedPrices: PricesHomePageData = orderByPagePosition(prices.data);

    const results: MergedHomePageData = mergeSellAndBuyOffers(orderedPrices);

    return {
        data: {
            prices: getRelevantPrices(results),
            news,
        },
        t: loc(language, "Home")
    };
}

export default function Home({loaderData}: Route.ComponentProps) {
  const { data, t } = loaderData;

  return (
      <div className="min-h-screen">
          <HomeHeader />
          <main className="container mx-auto px-4 py-8 space-y-8">
              <HomeTibiaNews news={data.news} />
              <section>
                  <div className="flex flex-col space-y-4 mb-6 md:flex-row md:items-center md:justify-between md:space-y-0">
                      <div>
                          <h2 className="text-2xl font-bold text-foreground">Market Overview</h2>
                          <p className="text-muted-foreground">Track the latest prices and trends</p>
                      </div>

                      <div className="flex items-center space-x-3 self-start md:self-auto">
                          <div className="flex items-center space-x-2 text-muted-foreground">
                              <Server className="w-4 h-4" />
                              <span className="text-sm font-medium">Server:</span>
                          </div>
                          <Select defaultValue="secura">
                              <SelectTrigger className="w-32 md:w-40">
                                  <SelectValue />
                              </SelectTrigger>
                              <SelectContent>
                                  <SelectItem value="antica">Antica</SelectItem>
                                  <SelectItem value="astera">Astera</SelectItem>
                                  <SelectItem value="belobra">Belobra</SelectItem>
                                  <SelectItem value="bombra">Bombra</SelectItem>
                                  <SelectItem value="celesta">Celesta</SelectItem>
                                  <SelectItem value="damora">Damora</SelectItem>
                                  <SelectItem value="dibra">Dibra</SelectItem>
                                  <SelectItem value="epoca">Epoca</SelectItem>
                                  <SelectItem value="ferobra">Ferobra</SelectItem>
                                  <SelectItem value="gladera">Gladera</SelectItem>
                                  <SelectItem value="harmonia">Harmonia</SelectItem>
                                  <SelectItem value="honbra">Honbra</SelectItem>
                                  <SelectItem value="impulsa">Impulsa</SelectItem>
                                  <SelectItem value="jacabra">Jacabra</SelectItem>
                                  <SelectItem value="kalibra">Kalibra</SelectItem>
                                  <SelectItem value="lobera">Lobera</SelectItem>
                                  <SelectItem value="menera">Menera</SelectItem>
                                  <SelectItem value="monza">Monza</SelectItem>
                                  <SelectItem value="nefera">Nefera</SelectItem>
                                  <SelectItem value="noctera">Noctera</SelectItem>
                                  <SelectItem value="ombra">Ombra</SelectItem>
                                  <SelectItem value="pacera">Pacera</SelectItem>
                                  <SelectItem value="peloria">Peloria</SelectItem>
                                  <SelectItem value="quintera">Quintera</SelectItem>
                                  <SelectItem value="refugia">Refugia</SelectItem>
                                  <SelectItem value="secura">Secura</SelectItem>
                                  <SelectItem value="solidera">Solidera</SelectItem>
                                  <SelectItem value="talera">Talera</SelectItem>
                                  <SelectItem value="tornera">Tornera</SelectItem>
                                  <SelectItem value="unica">Unica</SelectItem>
                                  <SelectItem value="venebra">Venebra</SelectItem>
                                  <SelectItem value="vita">Vita</SelectItem>
                                  <SelectItem value="wintera">Wintera</SelectItem>
                                  <SelectItem value="yonabra">Yonabra</SelectItem>
                                  <SelectItem value="zunera">Zunera</SelectItem>
                              </SelectContent>
                          </Select>
                      </div>
                  </div>

                  <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                      {Object.keys(data.prices).map((good: string, index: number) => {
                          return (
                              <PriceOverview key={index} good={good} data={data.prices[good]}/>
                          );
                      })}
                  </div>
              </section>
          </main>
      </div>
  );
}
