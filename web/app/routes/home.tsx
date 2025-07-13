import type {Route} from "./+types/home";
import {fetchPrices, orderHomePageData} from "~/home/routeFunctions";
import type {ApiResponse} from "~/lib/types";
import type {HomePageData} from "~/home/types";
import { CartesianGrid, LabelList, Line, LineChart, XAxis } from "recharts"
import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "~/components/ui/card";
import {type ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent} from "~/components/ui/chart";
import {TrendingUp} from "lucide-react";
import {formatDate} from "~/lib/utils";

const chartConfig = {
  buyOffer: {
    label: "Buy Offer",
    color: "var(--chart-1)",
  },
  sellOffer: {
    label: "Sell Offer",
    color: "var(--chart-2)",
  },
} satisfies ChartConfig

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
        <header className="flex items-center justify-center w-full h-48 text-center text-2xl font-bold text-foreground">
          <h1>Welcome to Tibia Mkt</h1>
        </header>
        <div className="flex w-full">
          <Card className="w-full">
            <CardHeader>
              <CardTitle>Line Chart - Label</CardTitle>
              <CardDescription>January - June 2024</CardDescription>
            </CardHeader>
            <CardContent>
              <ChartContainer config={chartConfig}>
                <LineChart
                    accessibilityLayer
                    data={data.honeycomb.prices}
                    margin={{
                      top: 20,
                      left: 12,
                      right: 12,
                    }}
                >
                  <CartesianGrid vertical={false} />
                  <XAxis
                      dataKey="createdAt"
                      tickLine={false}
                      axisLine={false}
                      tickMargin={8}
                      tickFormatter={formatDate}
                  />
                  <ChartTooltip
                      cursor={false}
                      content={<ChartTooltipContent indicator="line" />}
                  />
                  <Line
                      dataKey="buyOffer"
                      type="natural"
                      stroke="var(--primary)"
                      strokeWidth={2}
                      dot={{
                        fill: "var(--primary)",
                      }}
                      activeDot={{
                        r: 6,
                      }}
                  >
                    <LabelList
                        position="top"
                        offset={12}
                        className="fill-foreground"
                        fontSize={12}
                    />
                  </Line>
                </LineChart>
              </ChartContainer>
            </CardContent>
            <CardFooter className="flex-col items-start gap-2 text-sm">
              <div className="flex gap-2 leading-none font-medium">
                Trending up by 5.2% this month <TrendingUp className="h-4 w-4" />
              </div>
              <div className="text-muted-foreground leading-none">
                Showing total visitors for the last 6 months
              </div>
            </CardFooter>
          </Card>
        </div>
      </main>
  );
}
