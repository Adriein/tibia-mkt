import {CogChart, HomePageData} from "~/shared/types";
import classes from "~/components/GoodPreviewChip/GoodPreviewChip.module.css";
import {ActionIcon, Anchor, Card, Grid, GridCol, Group, Image, SimpleGrid, Stack, Title, Tooltip} from "@mantine/core";
import {beautifyCamelCase, camelCaseToSnakeCase, gif} from "~/shared/util";
import TibiaWikiIcon from "~/assets/tibia-wiki.png";
import {NavigateFunction, useNavigate} from "react-router";
import {TIBIA_COIN} from "~/shared/constants";
import {IconEye} from "@tabler/icons-react";

interface GoodPreviewChipProps {
    data: HomePageData
}

export function GoodPreviewChip({ data }: GoodPreviewChipProps) {
    const navigate: NavigateFunction = useNavigate();
    return (
        <Grid gutter="xl">
            {Object.keys(data).map((cogName: string) => {
                if (cogName === TIBIA_COIN) {
                    return null;
                }

                const cog: CogChart = data[cogName];

                return (
                    <GridCol key={cogName} span={2}>
                        <Card
                            radius="md"
                            withBorder
                            onClick={() => navigate(`/${camelCaseToSnakeCase(cogName)}/detail`)}
                            className={classes.card}
                        >
                            <Card.Section inheritPadding py="xs">
                                <Grid>
                                    <Grid.Col span={8}>
                                        <Stack gap={"xs"}>
                                            <Title order={4} lineClamp={1}>{beautifyCamelCase(cogName)}</Title>
                                            <Group>
                                                <Tooltip label="Go to TibiaWiki" openDelay={50}>
                                                    <Anchor href={cog.wiki} target="_blank">
                                                        <ActionIcon variant="default" aria-label="Tibia Wiki">
                                                            <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20.8} w={20.8}/>
                                                        </ActionIcon>
                                                    </Anchor>
                                                </Tooltip>
                                                <Tooltip label="Details" openDelay={50}>
                                                    <Anchor>
                                                        <ActionIcon
                                                            variant="default"
                                                            aria-label="Details"
                                                            onClick={() => navigate(`/${camelCaseToSnakeCase(cogName)}/detail`)}
                                                        >
                                                            <IconEye className={classes.eyeIconButton} />
                                                        </ActionIcon>
                                                    </Anchor>
                                                </Tooltip>
                                            </Group>
                                        </Stack>
                                    </Grid.Col>
                                    <Grid.Col span="content">
                                        <Image src={gif(cogName)} alt={cogName} fit="contain" h={70} w={70}/>
                                    </Grid.Col>
                                </Grid>
                            </Card.Section>
                        </Card>
                    </GridCol>
                );
            })}
        </Grid>
    );
}