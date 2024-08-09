import {Card, Group, Text} from "@mantine/core";
import {DetailCreature, SellOfferFrequency, TradeEngineDetailPageData} from "~/shared/types";
import {BarChart} from "@mantine/charts";
import classes from "./CogDetailGeneralInfo.module.css";

type SellOfferFrequencyBarChartTick = {frequency: string, range: string}

interface CogDetailGeneralInfoProps {
    item: string;
    dataPoints: number;
    creatures: DetailCreature[];
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

const presentCreatures = (creatures: DetailCreature[]): string  => {
    return creatures.map((creature: DetailCreature) => creature.name).join(', ');
}

const calculateDropEstimation = (creatures: DetailCreature[]): number => {
    return Math.round(creatures.reduce((total: number, creature: DetailCreature) => {
        total += creature.killStatistic * (creature.dropRate / 100);

        return total;
    }, 0));
}

export function CogDetailGeneralInfo({ dataPoints, creatures, data }: CogDetailGeneralInfoProps) {
    return (
        <>
            <Group justify="center" mb="md">
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>All time series</Text>
                    <Text size="xl" fw={700}>{dataPoints} data points</Text>
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
                {creatures.length &&
                    <>
                        <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                            <Text size="xl" fw={700}>Est. total dropped</Text>
                            <Text size="xl" fw={700}>
                                {new Intl.NumberFormat('en-US').format(calculateDropEstimation(creatures))}
                            </Text>
                        </Card>
                        <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                            <Text size="xl" fw={700}>Dropped by</Text>
                            <Text size="xl" fw={700}>{presentCreatures(creatures)}</Text>
                        </Card>
                    </>
                }
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
                            { name: 'frequency', label: "Frequency", color: 'teal.6' },
                        ]}
                        tickLine="y"
                    />
                </Card.Section>
            </Card>
        </>
    );
}