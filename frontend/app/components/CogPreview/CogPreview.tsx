import classes from "./CogPreview.module.css";
import {AreaChart} from '@mantine/charts';
import {Anchor, Badge, Card, ActionIcon, Space, Title, Image, Tooltip} from '@mantine/core';
import TibiaWikiIcon from '~/assets/tibia-wiki.png';
import {gif, beautifyCamelCase, camelCaseToSnakeCase} from "~/shared/util";
import {Cog, CogChart, YAxisTick} from "~/shared/types";
import {DEFAULT_WORLD} from "~/shared/constants";
import {IconEye} from '@tabler/icons-react';
import {NavigateFunction, useNavigate} from "react-router";
import {xAxisDateFormatter, xAxisTick, yAxisNumberFormatter} from "~/shared/chart-util";

interface CogPreviewProps {
    name: string;
    wikiLink: string;
    data: CogChart;
}

const tibiaServer = (data: Cog[]): string => data?.length? data[0].world : DEFAULT_WORLD;

export function CogPreview({ name, wikiLink, data }: CogPreviewProps) {
    const navigate: NavigateFunction = useNavigate();

    return (
        <Card withBorder shadow="sm" radius="md">
            <Card.Section withBorder inheritPadding py="xs">
                <div className={classes.chartHeader}>
                    <div className={classes.worldBadge}>
                        <Title order={2}>{beautifyCamelCase(name)}</Title>
                        <Image src={gif(name)} alt={name}/>
                    </div>
                    <Tooltip label="Details" openDelay={300}>
                        <ActionIcon
                            variant="default"
                            aria-label="Details"
                            onClick={() => navigate(`/${camelCaseToSnakeCase(name)}/detail`)}
                        >
                            <IconEye className={classes.eyeIconButton} />
                        </ActionIcon>
                    </Tooltip>
                    <Tooltip label="Go to TibiaWiki" openDelay={300}>
                        <Anchor href={wikiLink} target="_blank">
                            <ActionIcon variant="default" aria-label="Tibia Wiki">
                                <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20} w={20}/>
                            </ActionIcon>
                        </Anchor>
                    </Tooltip>
                    <Badge color="indigo">{tibiaServer(data.cogs)}</Badge>
                </div>
            </Card.Section>
            <Space h="xl"/>
            <AreaChart
                withLegend
                h={450}
                data={data.cogs}
                dataKey="date"
                tooltipAnimationDuration={200}
                series={[
                    {name: 'sellOffer', label: "Sell Offer", color: 'teal.6'},
                    {name: 'buyOffer', label: "Buy Offer", color: 'indigo.6'},
                ]}
                curveType="linear"
                legendProps={{verticalAlign: 'bottom'}}
                xAxisProps={{
                    interval: "preserveStartEnd",
                    tickFormatter: xAxisDateFormatter,
                    ticks: xAxisTick(data.cogs, data.chartMetadata.xAxisTick)
                }}
                yAxisProps={{
                    domain: data.chartMetadata.yAxisTick.map((tick: YAxisTick) => tick.price)
                }}
                valueFormatter={yAxisNumberFormatter}
            />
        </Card>
    )
}