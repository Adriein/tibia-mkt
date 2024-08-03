import {Card, Group, Text} from "@mantine/core";
import {TradeEngineDetailPageData} from "~/shared/types";

interface CogDetailGeneralInfoProps {
    data: TradeEngineDetailPageData;
}

export function CogDetailGeneralInfo({ data }: CogDetailGeneralInfoProps) {
    return (
        <Card withBorder shadow="sm" radius="md">
            <Card.Section withBorder inheritPadding py="xs">
                <Group justify="space-between">
                    <Text fw={500}>Honeycomb</Text>
                </Group>
            </Card.Section>
            <Card.Section mt="sm">
                <Text size="md">Historic Average: {data.historicAveragePrice}</Text>
            </Card.Section>
        </Card>
    );
}