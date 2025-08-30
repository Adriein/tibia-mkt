import {AlertTriangle, Clock, TrendingUp, Zap} from "lucide-react";
import {Carousel, CarouselContent, CarouselItem} from "~/components/ui/carousel";
import Autoplay from "embla-carousel-autoplay";
import {Card, CardContent, CardHeader} from "~/components/ui/card";
import {Badge} from "~/components/ui/badge";
import React from "react";
import type {TibiaArticleData} from "~/routes/home/types";

interface HomeTibiaNewsProps {
    news: TibiaArticleData[]
}

export function HomeTibiaNews({ news }: HomeTibiaNewsProps) {
    /*const news = [
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
    ]*/
    const getImpactColor = (category: string) => {
        switch (category) {
            case "community":
                return "bg-yellow-500/10 text-yellow-500 border-yellow-500/20"
            case "development":
                return "bg-blue-500/10 text-blue-500 border-blue-500/20"
            default:
                return "bg-muted text-muted-foreground"
        }
    }

    const getIconColor = (category: string) => {
        switch (category) {
            case "community":
                return "text-yellow-500"
            case "development":
                return "text-blue-500"
            default:
                return "text-muted-foreground"
        }
    }

    const getIcon = (category: string): string => {
        switch (category) {
            case "community":
                return "https://static.tibia.com/images/global/content/newsicon_community_big.png";
            case "development":
                return "https://static.tibia.com/images/global/content/newsicon_development_big.png";
            case "cipsoft":
                return "https://static.tibia.com/images/global/content/newsicon_cipsoft_big.png";
            case "support":
                return "https://static.tibia.com/images/global/content/newsicon_support_big.png";
            default:
                return "https://static.tibia.com/images/global/content/newsicon_technical_big.png"
        }
    }

    return (
        <section className="relative">
            <Carousel
                opts={{
                    align: "start",
                    loop: true,
                }}
                plugins={[
                    Autoplay({
                        delay: 6000,
                    }),
                ]}
                className="w-full mx-auto"
            >
                <CarouselContent className="-ml-1 md:-ml-4">
                    {news.map((article: TibiaArticleData, index) => {
                        const Icon = getIcon(article.category) as unknown as any
                        return (
                            <CarouselItem key={index} className="pl-1 md:pl-4 basis-full md:basis-1/2 max-w-md">
                                <Card className="bg-card border-border hover:bg-muted/50 transition-colors cursor-pointer h-full">
                                    <CardHeader className="pb-3">
                                        <div className="flex items-start justify-between gap-3">
                                            <div className="flex-1">
                                                <h3 className="font-semibold text-foreground leading-tight mb-3 text-sm md:text-base">
                                                    {article.title}
                                                </h3>
                                                <div className="flex items-center gap-2 mb-2">
                                                    <Badge variant="outline" className={getImpactColor(article.category)}>
                                                        {article.category}
                                                    </Badge>
                                                    <div className="flex items-center gap-1 text-xs text-muted-foreground">
                                                        <Clock className="w-3 h-3" />
                                                        {article.date}
                                                    </div>
                                                </div>
                                            </div>
                                            <div className={`p-2 rounded-lg bg-muted/50`}>
                                                <img src={getIcon(article.category)} className={`w-8 h-8 md:w-8 md:h-8`}  alt="category_icon"/>
                                            </div>
                                        </div>
                                    </CardHeader>
                                </Card>
                            </CarouselItem>
                        )
                    })}
                </CarouselContent>
            </Carousel>
        </section>
    )
}