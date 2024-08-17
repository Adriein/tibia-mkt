import {CogChart, HomePageData} from "~/shared/types";
import classes from "~/components/CogPreview/CogPreview.module.css";
import {ActionIcon, Anchor, Card, Grid, GridCol, Group, Image, Title, Tooltip} from "@mantine/core";
import {beautifyCamelCase, camelCaseToSnakeCase, gif} from "~/shared/util";
import TibiaWikiIcon from "~/assets/tibia-wiki.png";
import {NavigateFunction, useNavigate} from "react-router";
import {TIBIA_COIN} from "~/shared/constants";

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
                    <GridCol key={cogName} span={4}>
                        <Card
                            radius="md"
                            withBorder
                            onClick={() => navigate(`/${camelCaseToSnakeCase(cogName)}/detail`)}
                        >
                            <Group>
                                <Title order={4}>{beautifyCamelCase(cogName)}</Title>
                                <Image src={gif(cogName)} alt={cogName}/>
                                <Tooltip label="Go to TibiaWiki" openDelay={300}>
                                    <Anchor href={cog.wiki} target="_blank">
                                        <ActionIcon variant="default" aria-label="Tibia Wiki">
                                            <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20} w={20}/>
                                        </ActionIcon>
                                    </Anchor>
                                </Tooltip>
                            </Group>
                        </Card>
                    </GridCol>
                );
            })}
        </Grid>
    );
}