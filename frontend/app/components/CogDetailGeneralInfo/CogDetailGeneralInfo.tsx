import {Card, Group, Text} from "@mantine/core";
import {SellOfferFrequency, TradeEngineDetailPageData} from "~/shared/types";
import {BarChart} from "@mantine/charts";
import classes from "./CogDetailGeneralInfo.module.css";

type SellOfferFrequencyBarChartTick = {frequency: string, range: string}

interface CogDetailGeneralInfoProps {
    item: string;
    totalData: number;
    data: TradeEngineDetailPageData;
}

const sellOfferFrequencyBarChart = (data: SellOfferFrequency[]): SellOfferFrequencyBarChartTick[] => {
    const result: SellOfferFrequencyBarChartTick[] = [];

    for (const item of data) {
        result.push({
            frequency: new Intl.NumberFormat(
                'en-US',
                {
                    minimumFractionDigits: 2,
                    maximumFractionDigits: 2,
                }).format(item.frequency),
            range: item.range })
    }

    return result;
}

export function CogDetailGeneralInfo({ item, totalData, data }: CogDetailGeneralInfoProps) {
    return (
        <>
            <Group justify="center" mb="md">
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>All time series</Text>
                    <Text size="xl" fw={700}>{totalData} data points</Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>Mean</Text>
                    <Text size="xl" fw={700}>{data.mean} gp</Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>Standard Deviation</Text>
                    <Text size="xl" fw={700}>
                        {new Intl.NumberFormat('en-US', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2,
                        }).format(data.stdDeviation)} gp
                    </Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>Est. total dropped</Text>
                    <Text size="xl" fw={700}>
                        {new Intl.NumberFormat('en-US').format(500)}
                    </Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>Dropped by</Text>
                    <Text size="xl" fw={700}>
                        Wasp, Bear
                    </Text>
                </Card>
            </Group>
            <Card withBorder shadow="sm" radius="md">
                <Card.Section withBorder inheritPadding py="xs">
                    <Group justify="space-between">
                        <Text fw={500}>Segmented price frequency</Text>
                    </Group>
                </Card.Section>
                <Card.Section withBorder inheritPadding py="xl">
                    <BarChart
                        h={300}
                        data={sellOfferFrequencyBarChart(data.sellOfferFrequency)}
                        dataKey="range"
                        series={[
                            { name: 'frequency', color: 'teal.6' },
                        ]}
                        tickLine="y"
                    />
                </Card.Section>
            </Card>
        </>
    );
}