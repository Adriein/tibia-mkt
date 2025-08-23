import {fetchPrices, getRelevantPrices, mergeSellAndBuyOffers, orderByPagePosition} from "~/routes/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData, MergedHomePageData} from "~/routes/home/types";
import {PriceOverview} from "~/components/ui/price-overview";
import {English, type HomeTranslations, loc} from "~/locale/loc";
import type {Route} from "@/.react-router/types/app/routes/home/+types/home";
import React, {useEffect, useState} from "react";
import {
    AlertTriangle,
    Bell, ChevronLeft, ChevronRight, Clock,
    Search,
    Server,
    TrendingUp,
    Zap
} from "lucide-react";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {Button} from "~/components/ui/button";
import {Input} from "~/components/ui/input";
import {Card, CardContent, CardHeader} from "~/components/ui/card";
import {Badge} from "~/components/ui/badge";

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

                    <div className="flex items-center space-x-6">
                        {/* Search section */}
                        <div className="relative hidden md:block">
                            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-muted-foreground" />
                            <Input placeholder="Search items..." className="pl-10 w-64" />
                        </div>

                        <Select defaultValue="en">
                            <SelectTrigger className="w-16 h-9 text-sm font-medium">
                                <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="en">🇺🇸 EN</SelectItem>
                                <SelectItem value="es">🇪🇸 ES</SelectItem>
                                <SelectItem value="pt">🇧🇷 PT</SelectItem>
                                <SelectItem value="pl">🇵🇱 PL</SelectItem>
                            </SelectContent>
                        </Select>

                        {/* Notification section */}
                        <Button variant="outline" size="icon" className="relative">
                            <Bell className="w-5 h-5" />
                            <span className="absolute -top-1 -right-1 w-2 h-2 bg-primary rounded-full"></span>
                        </Button>
                    </div>
                </div>
            </div>
        </header>
    )
}

export function GameNews() {
    const [currentIndex, setCurrentIndex] = useState(0);

    const news = [
        {
            title: "Summer Update 2025 Released",
            summary: "New hunting grounds and rare items added. Expect price fluctuations on creature products.",
            date: "2 hours ago",
            category: "Update",
            impact: "high" as const,
            icon: Zap,
        },
        {
            title: "Double XP Weekend Announced",
            summary: "Increased demand for supplies and equipment expected this weekend.",
            date: "5 hours ago",
            category: "Event",
            impact: "medium" as const,
            icon: TrendingUp,
        },
        {
            title: "Server Maintenance Scheduled",
            summary: "Antica and Luminera will be offline for 2 hours. Trading may be affected.",
            date: "1 day ago",
            category: "Maintenance",
            impact: "low" as const,
            icon: AlertTriangle,
        },
        {
            title: "Rare Item Drop Rate Adjusted",
            summary: "Demon Horn and other rare drops have been rebalanced. Monitor price changes closely.",
            date: "2 days ago",
            category: "Balance",
            impact: "high" as const,
            icon: TrendingUp,
        },
        {
            title: "Summer Update 2025 Released",
            summary: "New hunting grounds and rare items added. Expect price fluctuations on creature products.",
            date: "2 hours ago",
            category: "Update",
            impact: "high" as const,
            icon: Zap,
        },
        {
            title: "Double XP Weekend Announced",
            summary: "Increased demand for supplies and equipment expected this weekend.",
            date: "5 hours ago",
            category: "Event",
            impact: "medium" as const,
            icon: TrendingUp,
        },
        {
            title: "Server Maintenance Scheduled",
            summary: "Antica and Luminera will be offline for 2 hours. Trading may be affected.",
            date: "1 day ago",
            category: "Maintenance",
            impact: "low" as const,
            icon: AlertTriangle,
        },
        {
            title: "Rare Item Drop Rate Adjusted",
            summary: "Demon Horn and other rare drops have been rebalanced. Monitor price changes closely.",
            date: "2 days ago",
            category: "Balance",
            impact: "high" as const,
            icon: TrendingUp,
        },
    ]

    const getImpactColor = (impact: string) => {
        switch (impact) {
            case "high":
                return "bg-red-500/10 text-red-500 border-red-500/20"
            case "medium":
                return "bg-yellow-500/10 text-yellow-500 border-yellow-500/20"
            case "low":
                return "bg-blue-500/10 text-blue-500 border-blue-500/20"
            default:
                return "bg-muted text-muted-foreground"
        }
    }

    const getIconColor = (impact: string) => {
        switch (impact) {
            case "high":
                return "text-red-500"
            case "medium":
                return "text-yellow-500"
            case "low":
                return "text-blue-500"
            default:
                return "text-muted-foreground"
        }
    }

    const nextSlide = () => {
        setCurrentIndex((prev) => (prev + 1) % Math.ceil(news.length / 4))
    }

    const prevSlide = () => {
        setCurrentIndex((prev) => (prev - 1 + Math.ceil(news.length / 4)) % Math.ceil(news.length / 4))
    }

    const goToSlide = (index: number) => {
        setCurrentIndex(index)
    }

    return (
        <section className="relative">
            <div className="flex items-center justify-between mb-4">
                <div className="flex gap-2">
                    {Array.from({ length: Math.ceil(news.length / 4) === 1 ? 2 : Math.ceil(news.length / 4) }).map((_, index) => (
                        <button
                            key={index}
                            onClick={(): void => goToSlide(index)}
                            className={`w-2 h-2 rounded-full transition-colors ${
                                currentIndex === index ? "bg-foreground" : "bg-muted-foreground/30"
                            }`}
                        />
                    ))}
                </div>
                <div className="flex gap-1">
                    <Button variant="ghost" size="sm" onClick={prevSlide} className="h-8 w-8 p-0 cursor-pointer">
                        <ChevronLeft className="w-4 h-4" />
                    </Button>
                    <Button variant="ghost" size="sm" onClick={nextSlide} className="h-8 w-8 p-0 cursor-pointer">
                        <ChevronRight className="w-4 h-4" />
                    </Button>
                </div>
            </div>

            <div className="overflow-hidden">
                <div
                    className="flex transition-transform duration-700 ease-out gap-4"
                    style={{ transform: `translateX(-${currentIndex * 100}%)` }}
                >
                    {Array.from({ length: Math.ceil(news.length / 4) }).map((_, slideIndex: number) => (
                        <div key={slideIndex} className="flex gap-4 min-w-full px-4">
                            {news.slice(slideIndex * 4, slideIndex * 4 + 4).map((article, index) => {
                                const Icon = article.icon
                                return (
                                    <Card
                                        key={slideIndex * 2 + index}
                                        className="bg-card border-border hover:bg-muted/50 transition-colors flex-1 max-w-md min-w-0"
                                    >
                                        <CardHeader className="pb-3">
                                            <div className="flex items-start justify-between gap-3">
                                                <div className="flex-1">
                                                    <h3 className="font-semibold text-foreground leading-tight mb-2">{article.title}</h3>
                                                    <div className="flex items-center gap-2 mb-2">
                                                        <Badge variant="outline" className={getImpactColor(article.impact)}>
                                                            {article.category}
                                                        </Badge>
                                                        <div className="flex items-center gap-1 text-xs text-muted-foreground">
                                                            <Clock className="w-3 h-3" />
                                                            {article.date}
                                                        </div>
                                                    </div>
                                                </div>
                                                <div className={`p-2 rounded-lg bg-muted/50`}>
                                                    <Icon className={`w-5 h-5 ${getIconColor(article.impact)}`} />
                                                </div>
                                            </div>
                                        </CardHeader>
                                        <CardContent className="pt-0">
                                            <p className="text-sm text-muted-foreground leading-relaxed">{article.summary}</p>
                                        </CardContent>
                                    </Card>
                                )
                            })}
                        </div>
                    ))}
                </div>
            </div>
        </section>
    )
}

export default function Home({loaderData}: Route.ComponentProps) {
  const { data, t } = loaderData;

  return (
      <div className="min-h-screen">
          <Header />
          <main className="container mx-auto px-4 py-8 space-y-8">
              <GameNews />
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
                      {Object.keys(data).map((good: string, index: number) => {
                          return (
                              <PriceOverview key={index} good={good} data={data[good]}/>
                          );
                      })}
                  </div>
              </section>
          </main>
      </div>
  );
}
