import {Card, Group, Text} from "@mantine/core";
import {SellOfferFrequency, TradeEngineDetailPageData} from "~/shared/types";
import {BarChart} from "@mantine/charts";

type SellOfferFrequencyBarChartTick = {frequency: number, price: string}

interface CogDetailGeneralInfoProps {
    item: string;
    data: TradeEngineDetailPageData;
}

const sellOfferFrequencyBarChart = (data: SellOfferFrequency[]): SellOfferFrequencyBarChartTick[] => {
    const result: SellOfferFrequencyBarChartTick[] = [];

    for (const item of data) {
        result.push({ frequency: item.frequency, price: new Intl.NumberFormat('en-US').format(item.price) })
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
            <Card.Section inheritPadding>
                <Text size="xl" fw={700}>Historic Average: {data.historicAveragePrice} GP</Text>
                <BarChart
                    h={300}
                    data={sellOfferFrequencyBarChart(data.sellOfferFrequency)}
                    dataKey="price"
                    series={[
                        { name: 'frequency', color: 'violet.6' },
                    ]}
                    tickLine="y"
                />
            </Card.Section>
        </Card>
    );
}