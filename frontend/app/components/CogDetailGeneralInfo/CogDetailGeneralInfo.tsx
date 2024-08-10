import {ActionIcon, Card, Group, Modal, Text, Timeline, Tooltip} from "@mantine/core";
import {BarChart} from "@mantine/charts";
import {IconHistory, IconBinary} from '@tabler/icons-react';
import classes from "./CogDetailGeneralInfo.module.css";
import {DetailCreature, SellOfferFrequency, SellOfferHistoricData, SellOfferProbability} from "~/shared/types";
import {useDisclosure} from "@mantine/hooks";


type SellOfferFrequencyBarChartTick = {frequency: string, range: string}

interface CogDetailGeneralInfoProps {
    item: string;
    dataPoints: number;
    creatures: DetailCreature[];
    historic: SellOfferHistoricData[];
    data: SellOfferProbability;
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

const calculateDropEstimation = (creatures: DetailCreature[]): number => {
    return Math.round(creatures.reduce((total: number, creature: DetailCreature) => {
        total += creature.killStatistic * (creature.dropRate / 100);

        return total;
    }, 0));
}

export function CogDetailGeneralInfo({ dataPoints, creatures, data, historic }: CogDetailGeneralInfoProps) {
    const [opened, { open, close }] = useDisclosure(false);

    return (
        <>
            <Modal opened={opened} onClose={close} title="History" centered>
                <Timeline active={historic.length} bulletSize={24} lineWidth={2}>
                    {historic.map((value: SellOfferHistoricData) => {
                        return (
                            <Timeline.Item key={value.id} bullet={<IconBinary size={18} />} title="Ingested data">
                                <Text c="dimmed" size="sm">{value.stdDeviation}</Text>
                                <Text size="xs" mt={4}>{value.createdAt}</Text>
                            </Timeline.Item>
                        );
                    })}
                </Timeline>
            </Modal>
            <Group justify="center" mb="md">
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>All time series</Text>
                    <Text size="xl" fw={500}>{dataPoints} data points</Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Text size="xl" fw={700}>Mean</Text>
                    <Text size="xl" fw={500}>{data.mean} gp</Text>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Group>
                        <Text size="xl" fw={700}>Standard Deviation</Text>
                        <Tooltip label="View history" openDelay={300}>
                            <ActionIcon onClick={open} variant="default" aria-label="History">
                                <IconHistory style={{ width: '70%', height: '70%' }} stroke={1.5} />
                            </ActionIcon>
                        </Tooltip>
                    </Group>
                    <Text size="xl" fw={500}>
                        {new Intl.NumberFormat('en-US', {
                            minimumFractionDigits: 2,
                            maximumFractionDigits: 2,
                        }).format(data.stdDeviation)} gp
                    </Text>
                </Card>
                {creatures.length &&
                    <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                        <Group>
                            <Text size="xl" fw={700}>Est. total dropped</Text>
                            <Tooltip label="View history" openDelay={300}>
                                <ActionIcon onClick={open} variant="default" aria-label="History">
                                    <IconHistory style={{ width: '70%', height: '70%' }} stroke={1.5} />
                                </ActionIcon>
                            </Tooltip>
                        </Group>
                        <Text size="xl" fw={500}>
                            {new Intl.NumberFormat('en-US').format(calculateDropEstimation(creatures))}
                        </Text>
                    </Card>
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