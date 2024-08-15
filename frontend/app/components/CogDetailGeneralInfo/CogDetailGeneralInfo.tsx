import {ActionIcon, Card, Group, Modal, Text, Timeline, Tooltip} from "@mantine/core";
import {BarChart} from "@mantine/charts";
import {IconHistory, IconBinary, IconEqual, IconArrowUp, IconArrowDown} from '@tabler/icons-react';
import classes from "./CogDetailGeneralInfo.module.css";
import {DetailCreature, SellOfferFrequency, SellOfferHistoricData, SellOfferProbability} from "~/shared/types";
import {useDisclosure} from "@mantine/hooks";
import {useState} from "react";
import {MEAN_HISTORY_MODAL, STD_DEVIATION_MODAL, TOTAL_DROPPED_MODAL} from "~/shared/constants";
import {beautifyCamelCase} from "~/shared/util";


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

const countMajorTendency = (tendency: string[]): string => {
    let value = 0;
    let maxKey = null;

    const counter = tendency.reduce((counts: Record<string, number>, direction: string) => {
        counts[direction]++;

        return counts;
    }, { up: 0, down: 0, equal: 0 });

    for (const key in counter) {
        if (counter[key] > value) {
            maxKey = key;
            value = counter[key];
        }
    }

    return maxKey as string;
}

const renderHistoricTrendIcon = (context: string, data: SellOfferHistoricData[]) => {
    const tendency: string[] = [];
    let previousValue: number = 0;

    const last7DaysPortion: SellOfferHistoricData[] = data.slice(data.length - 8, data.length - 1);

    for (const sellOfferHistoricData of last7DaysPortion) {
        if(tendency.length === 7) {
            break;
        }

        const value: number = sellOfferHistoricData[context as keyof SellOfferHistoricData] as number;

        if (previousValue > value) {
            tendency.push("down");

            previousValue = value
        }

        if (previousValue === value) {
            tendency.push("equal");

            previousValue = value
        }

        if (previousValue < value) {
            tendency.push("up");

            previousValue = value
        }
    }

    const result: string = countMajorTendency(tendency);

    if (result === "up") {
        return <IconArrowUp />
    }

    if (result === "equal") {
        return <IconEqual />
    }

    return <IconArrowDown/>
}

export function CogDetailGeneralInfo({ dataPoints, creatures, data, historic }: CogDetailGeneralInfoProps) {
    const [opened, { open, close }] = useDisclosure(false);
    const [context, setContext] = useState(MEAN_HISTORY_MODAL);

    const assignContextOnOpenModal = (context: string): void => {
        setContext(context);
        open();
    }

    return (
        <>
            <Modal opened={opened} onClose={close} title="History" centered>
                <Timeline active={historic.length} bulletSize={24} lineWidth={2}>
                    {historic.map((value: SellOfferHistoricData) => {
                        const d: Record<string, never> = value as unknown as Record<string, never>;
                        return (
                            <Timeline.Item
                                key={value.id}
                                bullet={<IconBinary size={18} />}
                                title={`${beautifyCamelCase(context)} ingestion`}
                            >
                                <Text c="dimmed" size="sm">
                                    {
                                        new Intl.NumberFormat('en-US', {
                                            maximumFractionDigits: 2,
                                        }).format(d[context])
                                    }
                                </Text>
                                <Text size="xs" mt={4}>
                                    {
                                        new Intl.DateTimeFormat('en-US', {
                                            month: "long",
                                            weekday: "short",
                                            day: "2-digit"
                                        }).format(new Date(d.createdAt))
                                    }
                                </Text>
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
                    <Group>
                        <Text size="xl" fw={700}>Mean</Text>
                        <Tooltip label="View history" openDelay={300}>
                            <ActionIcon
                                onClick={() => assignContextOnOpenModal(MEAN_HISTORY_MODAL)}
                                variant="default"
                                aria-label="Mean history"
                            >
                                <IconHistory style={{ width: '70%', height: '70%' }} stroke={1.5} />
                            </ActionIcon>
                        </Tooltip>
                    </Group>
                    <Group>
                        <Text size="xl" fw={500}>{data.mean} gp</Text>
                        {renderHistoricTrendIcon(MEAN_HISTORY_MODAL, historic)}
                    </Group>
                </Card>
                <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                    <Group>
                        <Text size="xl" fw={700}>Standard Deviation</Text>
                        <Tooltip label="View history" openDelay={300}>
                            <ActionIcon
                                onClick={() => assignContextOnOpenModal(STD_DEVIATION_MODAL)}
                                variant="default"
                                aria-label="Standard deviation history"
                            >
                                <IconHistory style={{ width: '70%', height: '70%' }} stroke={1.5} />
                            </ActionIcon>
                        </Tooltip>
                    </Group>
                    <Group>
                        <Text size="xl" fw={500}>
                            {new Intl.NumberFormat('en-US', {
                                minimumFractionDigits: 2,
                                maximumFractionDigits: 2,
                            }).format(data.stdDeviation)} gp
                        </Text>
                        {renderHistoricTrendIcon(STD_DEVIATION_MODAL, historic)}
                    </Group>
                </Card>
                {creatures.length &&
                    <Card padding="lg" radius="md" withBorder className={classes.infoCard}>
                        <Group>
                            <Text size="xl" fw={700}>Est. total dropped</Text>
                            <Tooltip label="View history" openDelay={300}>
                                <ActionIcon
                                    onClick={() => assignContextOnOpenModal(TOTAL_DROPPED_MODAL)}
                                    variant="default"
                                    aria-label="item dropped amount history"
                                >
                                    <IconHistory style={{ width: '70%', height: '70%' }} stroke={1.5} />
                                </ActionIcon>
                            </Tooltip>
                        </Group>
                        <Group>
                            <Text size="xl" fw={500}>
                                {new Intl.NumberFormat('en-US').format(calculateDropEstimation(creatures))}
                            </Text>
                            {renderHistoricTrendIcon(TOTAL_DROPPED_MODAL, historic)}
                        </Group>
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