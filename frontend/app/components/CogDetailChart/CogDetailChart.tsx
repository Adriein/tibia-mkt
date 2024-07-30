import {xAxisDateFormatter, xAxisTick, yAxisNumberFormatter} from "~/shared/chart-util";
import {DetailPageData, YAxisTick} from "~/shared/types";
import {LineChart} from "@mantine/charts";
import {Card, Group, SegmentedControl, Text} from '@mantine/core';

interface CogDetailChartProps {
    data: DetailPageData;
}

const TIME_INTERVAL: string[] = ['Last Week', 'Last Month', 'Last 3 Months', 'Last 6 Months', 'Last Year', "All Series"];

export function CogDetailChart({ data }: CogDetailChartProps) {
    return (
        <>
            <SegmentedControl mb="lg" withItemsBorders={false} data={TIME_INTERVAL} />
            <Card withBorder radius="md" mb="lg">
                <Card.Section withBorder inheritPadding py="xs">
                    <Group justify="space-between">
                        <Text fw={500}>Sell Offers</Text>
                    </Group>
                </Card.Section>
                <Card.Section mt="sm" p="sm">
                    <LineChart
                        h={180}
                        data={data.cog}
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
                        h={180}
                        data={data.cog}
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