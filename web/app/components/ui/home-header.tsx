import {Bell, Search, TrendingUp} from "lucide-react";
import {Input} from "~/components/ui/input";
import {Select, SelectContent, SelectItem, SelectTrigger, SelectValue} from "~/components/ui/select";
import {Button} from "~/components/ui/button";
import React, {useState} from "react";
import {Avatar, AvatarImage} from "~/components/ui/avatar";
import HoneycombGif from "~/assets/honeycomb.gif";
import {useSearchParams} from "react-router";
import {BeautyLocale} from "~/locale/loc";

export function HomeHeader() {
    const [searchParams, setSearchParams] = useSearchParams();
    const [currentLang, setState] = useState(searchParams.get("lang") || BeautyLocale.English);

    const handleLanguageChange = (newLang: string): void => {
        setState(newLang);
        setSearchParams({ lang: newLang });
    };

    return (
        <header className="border-b border-border bg-card">
            <div className="container mx-auto px-4 py-4">
                <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-3">
                        <div className="flex items-center justify-center w-10 h-10 bg-primary rounded-lg">
                            {/*<TrendingUp className="w-6 h-6 text-primary-foreground" />*/}
                            <Avatar className="w-10 h-10">
                                <AvatarImage src={HoneycombGif} />
                            </Avatar>
                        </div>
                        <div>
                            <h1 className="text-2xl font-bold text-foreground">Tibia Mkt</h1>
                            <p className="text-sm text-muted-foreground">Making trade easy</p>
                        </div>
                    </div>

                    <div className="flex items-center space-x-6">
                        {/* Search section */}
                        <div className="relative hidden md:block">
                            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 w-4 h-4 text-muted-foreground" />
                            <Input placeholder="Search items..." className="pl-10 w-64" />
                        </div>

                        <Select defaultValue={currentLang} onValueChange={handleLanguageChange}>
                            <SelectTrigger className="w-16 h-9 text-sm font-medium">
                                <SelectValue />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value={BeautyLocale.English}>🇺🇸 EN</SelectItem>
                                <SelectItem value={BeautyLocale.Spanish}>🇪🇸 ES</SelectItem>
                                <SelectItem value={BeautyLocale.Portuguese}>🇧🇷 PT</SelectItem>
                                <SelectItem value={BeautyLocale.Polish}>🇵🇱 PL</SelectItem>
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