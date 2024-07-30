import {xAxisDateFormatter, xAxisTick, yAxisNumberFormatter} from "~/shared/chart-util";
import {Cog, DetailPageData, YAxisTick} from "~/shared/types";
import {LineChart} from "@mantine/charts";
import {Card, Group, SegmentedControl, Text} from '@mantine/core';
import {useEffect, useState} from "react";

interface CogDetailChartProps {
    data: DetailPageData;
}

const TIME_INTERVAL: string[] = ['Last Week', 'Last Month', 'Last 3 Months', 'Last 6 Months', 'Last Year', "All Series"];

const TIME_INTERVAL_DIC: Record<string, number> = {
    'Last Week': 7,
    'Last Month': 30,
    'Last 3 Months': 90,
    'Last 6 Months': 180,
    'Last Year': 365,
    'All Series': 0,
};

export function CogDetailChart({ data }: CogDetailChartProps) {
    const [timeInterval, setTimeInterval] = useState<string>('Last Week');
    const [cogs, setCogs] = useState<Cog[]>([]);

    useEffect(() => {
        const chartData: Cog[] = data.cog;

        const result: Cog[] = [];

        for (let i: number = 0; i < chartData.length; i++) {
            if (TIME_INTERVAL_DIC[timeInterval] === 0) {
                result.push(...data.cog);

                break;
            }

            if (TIME_INTERVAL_DIC[timeInterval] === i + 1) {
                result.push(chartData[i]);

                break;
            }

            result.push(chartData[i]);
        }

        setCogs(result);
    }, [timeInterval, data.cog]);

    return (
        <>
            <SegmentedControl
                value={timeInterval}
                onChange={setTimeInterval}
                mb="lg"
                withItemsBorders={false}
                data={TIME_INTERVAL}
            />
            <Card withBorder radius="md" mb="lg">
                <Card.Section withBorder inheritPadding py="xs">
                    <Group justify="space-between">
                        <Text fw={500}>Sell Offers</Text>
                    </Group>
                </Card.Section>
                <Card.Section mt="sm" p="sm">
                    <LineChart
                        h={200}
                        data={cogs}
                        dataKey="date"
                        series={[{name: 'sellOffer', label: "Sell Offer", color: 'teal.6'}]}
                        xAxisProps={{
                            interval: "preserveStartEnd",
                            tickFormatter: xAxisDateFormatter,
                            ticks: xAxisTick(data.cog, data.sellOfferChart.xAxisTick)
                        }}
                        yAxisProps={{
                            domain: data.sellOfferChart.yAxisTick.map((tick: YAxisTick) => tick.price)
                        }}
                        valueFormatter={yAxisNumberFormatter}
                        lineChartProps={{ syncId: 'offer' }}
                    />
                </Card.Section>
            </Card>
            <Card withBorder shadow="sm" radius="md">
                <Card.Section withBorder inheritPadding py="xs">
                    <Group justify="space-between">
                        <Text fw={500}>Buy Offers</Text>
                    </Group>
                </Card.Section>
                <Card.Section mt="sm" p="sm">
                    <LineChart
                        h={200}
                        data={cogs}
                        dataKey="date"
                        series={[{name: 'buyOffer', label: "Buy Offer", color: 'indigo.6'}]}
                        xAxisProps={{
                            interval: "preserveStartEnd",
                            tickFormatter: xAxisDateFormatter,
                            ticks: xAxisTick(data.cog, data.buyOfferChart.xAxisTick)
                        }}
                        yAxisProps={{
                            domain: data.buyOfferChart.yAxisTick.map((tick: YAxisTick) => tick.price)
                        }}
                        valueFormatter={yAxisNumberFormatter}
                        lineChartProps={{ syncId: 'offer' }}
                    />
                </Card.Section>
            </Card>
        </>
    );
}