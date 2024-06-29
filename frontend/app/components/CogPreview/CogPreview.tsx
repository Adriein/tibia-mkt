import classes from "./CogPreview.module.css";
import { AreaChart } from '@mantine/charts';
import {Anchor, Badge, Card, ActionIcon, Space, Title, Image, Tooltip} from '@mantine/core';
import TibiaWikiIcon from '~/assets/tibia-wiki.png';
import {formatDate, gif, beautifyCamelCase} from "~/shared/util";
import {Cog, CogChart} from "~/shared/types";
import {DEFAULT_WORLD} from "~/shared/constants";

interface CogPreviewProps {
    name: string;
    data: CogChart;
}

const tibiaServer = (data: Cog[]): string => data?.length? data[0].world : DEFAULT_WORLD;

const xAxisDateFormatter = (value: string): string => formatDate(new Date(value));
const yAxisNumberFormatter = (value: string): string => new Intl.NumberFormat('en-US').format(value);

const xAxisTick = (data: Cog[], xAxisDomain: string[]): string[] => {
    const SHOW_DATES: string[] = xAxisDomain;
    const result: string[] = [];

    for (let i: number = 0; i < data.length; i++) {
        const point: Cog = data[i];
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

export function CogPreview({ name, data }: CogPreviewProps) {
    return (
        <Card withBorder shadow="sm" radius="md">
            <Card.Section withBorder inheritPadding py="xs">
                <div className={classes.chartHeader}>
                    <div className={classes.worldBadge}>
                        <Title order={2}>{beautifyCamelCase(name)}</Title>
                        <Image src={gif(name)} alt="Tibia Coin"/>
                    </div>
                    <Tooltip label="Go to TibiaWiki" openDelay={300}>
                        <Anchor href="https://tibia.fandom.com/wiki/Tibia_Coins" target="_blank">
                            <ActionIcon variant="default" aria-label="Tibia Wiki">
                                <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20} w={20}/>
                            </ActionIcon>
                        </Anchor>
                    </Tooltip>
                    <Badge color="indigo">{tibiaServer(data.cog)}</Badge>
                </div>
            </Card.Section>
            <Space h="xl"/>
            <AreaChart
                h={400}
                data={data.cog}
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
                    ticks: xAxisTick(data.cog, data.chartMetadata.xAxisTick)
                }}
                yAxisProps={{
                    domain: data.chartMetadata.yAxisTick
                }}
                valueFormatter={yAxisNumberFormatter}
            />
        </Card>
    )
}