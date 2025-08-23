import {fetchPrices, getRelevantPrices, mergeSellAndBuyOffers, orderByPagePosition} from "~/routes/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData, MergedHomePageData} from "~/routes/home/types";
import {PriceOverview} from "~/components/ui/price-overview";
import {English, type HomeTranslations, loc} from "~/locale/loc";
import type {Route} from "@/.react-router/types/app/routes/home/+types/home";
import React from "react";
import {Activity, Bell, DollarSign, Globe, Search, Server, TrendingDown, TrendingUp} from "lucide-react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {Button} from "~/components/ui/button";
import {Input} from "~/components/ui/input";
import {Card, CardContent} from "~/components/ui/card";

export function Header() {
    return (
        <header className="border-b border-border bg-card">
            <div className="container mx-auto px-4 py-4">
                <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-3">
                        <div className="flex items-center justify-center w-10 h-10 bg-primary rounded-lg">
                            <TrendingUp className="w-6 h-6 text-primary-foreground" />
                        </div>
                        <div>
                            <h1 className="text-2xl font-bold text-foreground">Tibia Mkt</h1>
                            <p className="text-sm text-muted-foreground">Trading Market</p>
                        </div>
                    </div>

                    <div className="flex items-center space-x-4">
                        <div className="relative hidden md:block">
                            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-muted-foreground" />
                            <Input placeholder="Search items..." className="pl-10 w-64" />
                        </div>

                        <div className="flex items-center space-x-2">
                            <Globe className="w-4 h-4 text-muted-foreground" />
                            <Select defaultValue="en">
                                <SelectTrigger className="w-20 h-9 border-0 bg-transparent">
                                    <SelectValue />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="en">EN</SelectItem>
                                    <SelectItem value="es">ES</SelectItem>
                                    <SelectItem value="pt">PT</SelectItem>
                                    <SelectItem value="pl">PL</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                        <Button variant="ghost" size="icon">
                            <Bell className="w-5 h-5" />
                        </Button>
                    </div>
                </div>
            </div>
        </header>
    )
}

export function StatsOverview() {
    const stats = [
        {
            title: "Total Volume",
            value: "2.4M",
            change: "+12.5%",
            changeType: "positive" as const,
            icon: DollarSign,
        },
        {
            title: "Active Items",
            value: "1,247",
            change: "+3.2%",
            changeType: "positive" as const,
            icon: Activity,
        },
        {
            title: "Top Gainer",
            value: "Honeycomb",
            change: "+15.2%",
            changeType: "positive" as const,
            icon: TrendingUp,
        },
        {
            title: "Top Loser",
            value: "Swampling Wood",
            change: "-3.1%",
            changeType: "negative" as const,
            icon: TrendingDown,
        },
    ]

    return (
        <section>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                {stats.map((stat) => {
                    const Icon = stat.icon
                    return (
                        <Card key={stat.title} className="bg-card border-border">
                            <CardContent className="p-6">
                                <div className="flex items-center justify-between">
                                    <div>
                                        <p className="text-sm text-muted-foreground">{stat.title}</p>
                                        <p className="text-2xl font-bold text-foreground">{stat.value}</p>
                                        <p
                                            className={`text-sm font-medium ${
                                                stat.changeType === "positive" ? "text-green-500" : "text-red-500"
                                            }`}
                                        >
                                            {stat.change}
                                        </p>
                                    </div>
                                    <div
                                        className={`p-3 rounded-lg ${stat.changeType === "positive" ? "bg-green-500/10" : "bg-red-500/10"}`}
                                    >
                                        <Icon className={`w-6 h-6 ${stat.changeType === "positive" ? "text-green-500" : "text-red-500"}`} />
                                    </div>
                                </div>
                            </CardContent>
                        </Card>
                    )
                })}
            </div>
        </section>
    )
}

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
      <div className="min-h-screen">
          <Header />
          <main className="container mx-auto px-4 py-8 space-y-8">
              <StatsOverview />
              <section>
                  <div className="flex items-center justify-between mb-6">
                      <div>
                          <h2 className="text-2xl font-bold text-foreground">Market Overview</h2>
                          <p className="text-muted-foreground">Track the latest prices and trends</p>
                      </div>

                      <div className="flex items-center space-x-3">
                          <div className="flex items-center space-x-2 text-muted-foreground">
                              <Server className="w-4 h-4" />
                              <span className="text-sm font-medium">Server:</span>
                          </div>
                          <Select defaultValue="secura">
                              <SelectTrigger className="w-40">
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
                      {Object.keys(data).map((good: string, index: number) => {
                          return (
                              <PriceOverview good={good} data={data[good]}/>
                          );
                      })}
                  </div>
              </section>
          </main>
      </div>
  );
}
