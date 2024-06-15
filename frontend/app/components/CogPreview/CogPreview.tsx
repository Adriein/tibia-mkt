import { AreaChart } from '@mantine/charts';
import {Anchor, Badge, Flex, Space, Text, Title} from '@mantine/core';
import TibiaCoinGif from '~/assets/tibia-coin.gif';
import { formatDate } from "~/shared/util";
import { TibiaCoinCog } from "~/shared/types";

const tibiaServer = (data: TibiaCoinCog[]): string => data[0].world;

const xAxisDateFormatter = (value: string): string => formatDate(new Date(value));
const yAxisNumberFormatter = (value: string): string => new Intl.NumberFormat('en-US').format(value);

const xAxisTick = (data: TibiaCoinCog[]): string[] => {
    const SHOW_DATES: string[] = ['01', '10', '20', '30', '31'];
    const result: string[] = [];

    for (let i: number = 0; i < data.length; i++) {
        const point: TibiaCoinCog = data[i];
        const day: string = point.date.split("-")[2];

        if (!SHOW_DATES.includes(day)) {
            continue;
        }

        result.push(point.date)
    }

    return result;
}

const yAxisTick = (data: TibiaCoinCog[]): number[] => {
    const SHOW_NUMBERS: number[] = [30000, 35000, 40000, 45000, 50000];

    return SHOW_NUMBERS;
}

export function CogPreview({data}) {
    return (
        <>
            <Flex align="center" gap="md">
                <img src={TibiaCoinGif as string} alt="Tibia Coin"/>
                <Anchor href="https://tibia.fandom.com/wiki/Tibia_Coins" target="_blank">
                    <Title order={2}>Tibia Coin</Title>
                </Anchor>
                <Badge color="indigo">{tibiaServer(data)}</Badge>
            </Flex>
            <Space h="xl" />
            <AreaChart
                h={400}
                data={data}
                dataKey="date"
                series={[
                    {name: 'buyPrice', label: "Buy price", color: 'indigo.6'},
                    {name: 'sellPrice', label: "Sell price", color: 'teal.6'},
                ]}
                curveType="linear"
                legendProps={{verticalAlign: 'bottom'}}
                xAxisProps={{interval: "preserveStartEnd", tickFormatter: xAxisDateFormatter, ticks: xAxisTick(data)}}
                yAxisProps={{domain: [40000, 50000], ticks: yAxisTick(data)}}
                valueFormatter={yAxisNumberFormatter}
            />
        </>
    )
}