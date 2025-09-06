import {Clock, ExternalLink} from "lucide-react";
import {Carousel, CarouselContent, CarouselItem} from "~/components/ui/carousel";
import Autoplay from "embla-carousel-autoplay";
import {Card, CardContent, CardHeader} from "~/components/ui/card";
import {Badge} from "~/components/ui/badge";
import React from "react";
import type {TibiaArticleData} from "~/routes/home/types";
import {Link} from "react-router";

interface HomeTibiaNewsProps {
    news: TibiaArticleData[]
}

export function HomeTibiaNews({ news }: HomeTibiaNewsProps) {
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
                        return (
                            <CarouselItem key={index} className="pl-1 md:pl-4 basis-full md:basis-1/2 max-w-md">
                                <Card className="bg-card border-border hover:border-border/60 transition-all duration-200 hover:shadow-lg hover:shadow-black/20 group">
                                    <CardHeader className="pb-3">
                                        <div className="flex items-start justify-between gap-3">
                                            <div className="flex items-center gap-2 min-w-0 flex-1">
                                                <div className="p-1.5 rounded-md bg-muted/50">
                                                    <img src={getIcon(article.category)} className="w-4 h-4"  alt="category_icon"/>
                                                </div>
                                                <h3 className="font-semibold text-foreground text-lg leading-tight truncate" title={article.title}>
                                                    {article.title}
                                                </h3>
                                            </div>
                                            <button className="p-1.5 rounded-md hover:bg-muted/50 transition-colors opacity-60 group-hover:opacity-100 cursor-pointer">
                                                <Link
                                                    to={article.url}
                                                    aria-label="visit tibia.com"
                                                    target="_blank"
                                                    rel="noreferrer"
                                                >
                                                    <ExternalLink className="h-4 w-4 text-muted-foreground" />
                                                </Link>
                                            </button>
                                        </div>
                                    </CardHeader>
                                    <CardContent className="pt-0">
                                        <div className="flex items-center justify-between">
                                            <Badge variant="outline" className={getImpactColor(article.category)}>
                                                {article.category}
                                            </Badge>
                                            <div className="flex items-center gap-1.5 text-sm text-muted-foreground">
                                                <Clock className="h-3.5 w-3.5" />
                                                <span>{article.date}</span>
                                            </div>
                                        </div>
                                    </CardContent>
                                </Card>
                            </CarouselItem>
                        );
                    })}
                </CarouselContent>
            </Carousel>
        </section>
    )
}