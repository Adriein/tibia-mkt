import { AreaChart } from '@mantine/charts';
import {Anchor, Badge, Flex, Space, Text, Title} from '@mantine/core';
import TibiaCoinGif from '~/assets/tibia-coin.gif';
import { formatDate } from "~/shared/util";
import {HomePageData, TibiaCoinCog} from "~/shared/types";
import {DEFAULT_WORLD} from "~/shared/constants";

interface CogPreviewProps {
    data: HomePageData
}

const tibiaServer = (data: TibiaCoinCog[]): string => data?.length? data[0].world : DEFAULT_WORLD;

const xAxisDateFormatter = (value: string): string => formatDate(new Date(value));
const yAxisNumberFormatter = (value: string): string => new Intl.NumberFormat('en-US').format(value);

const xAxisTick = (data: TibiaCoinCog[], xAxisDomain: string[]): string[] => {
    const SHOW_DATES: string[] = xAxisDomain;
    const result: string[] = [];

    for (let i: number = 0; i < data.length; i++) {
        const point: TibiaCoinCog = data[i];
        const day: string = point.date.split("-")[2];

        if (i == 0 || i == data.length - 1) {
            result.push(point.date)
        }

        if (!SHOW_DATES.includes(day)) {
            continue;
        }

        result.push(point.date)
    }

    return result;
}

export function CogPreview({ data }: CogPreviewProps) {
    return (
        <>
            <Flex align="center" gap="md">
                <Anchor href="https://tibia.fandom.com/wiki/Tibia_Coins" target="_blank">
                    <Title order={2}>Tibia Coin</Title>
                </Anchor>
                <img src={TibiaCoinGif as string} alt="Tibia Coin"/>
                <Badge color="indigo">{tibiaServer(data.cogs)}</Badge>
            </Flex>
            <Space h="xl"/>
            <AreaChart
                h={400}
                data={data.cogs}
                dataKey="date"
                tooltipAnimationDuration={200}
                series={[
                    {name: 'buyPrice', label: "Buy price", color: 'indigo.6'},
                    {name: 'sellPrice', label: "Sell price", color: 'teal.6'},
                ]}
                curveType="linear"
                legendProps={{verticalAlign: 'bottom'}}
                xAxisProps={{
                    interval: "preserveStartEnd",
                    tickFormatter: xAxisDateFormatter,
                    ticks: xAxisTick(data.cogs, data.chartMetadata.xAxisTick)
                }}
                yAxisProps={{
                    domain: data.chartMetadata.yAxisTick
                }}
                valueFormatter={yAxisNumberFormatter}
            />
        </>
    )
}