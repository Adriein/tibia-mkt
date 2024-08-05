import {Card, Group, Text} from "@mantine/core";
import {SellOfferFrequency, TradeEngineDetailPageData} from "~/shared/types";
import {BarChart} from "@mantine/charts";

type SellOfferFrequencyBarChartTick = {frequency: string, range: string}

interface CogDetailGeneralInfoProps {
    item: string;
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

export function CogDetailGeneralInfo({ item, data }: CogDetailGeneralInfoProps) {
    return (
        <Card withBorder shadow="sm" radius="md">
            <Card.Section withBorder inheritPadding py="xs">
                <Group justify="space-between">
                    <Text fw={500}>{item} Insights</Text>
                </Group>
            </Card.Section>
            <Card.Section withBorder inheritPadding py="xs">
                <Text size="xl" fw={700}>Mean: {data.mean} gp</Text>
                <Text size="xl" fw={700}>
                    Standard Deviation: {new Intl.NumberFormat('en-US', {
                        minimumFractionDigits: 2,
                        maximumFractionDigits: 2,
                }).format(data.stdDeviation)} gp
                </Text>

            </Card.Section>
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
                        { name: 'frequency', color: 'violet.6' },
                    ]}
                    tickLine="y"
                />
            </Card.Section>
        </Card>
    );
}