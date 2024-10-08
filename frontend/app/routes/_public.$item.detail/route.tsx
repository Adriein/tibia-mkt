import {beautifyCamelCase, snakeCaseToCamelCase} from "~/shared/util";
import {Accordion, Container, Grid, rem} from "@mantine/core";
import {LoaderFunctionArgs} from "@remix-run/node";
import {Cog, DetailPageData, RemixMetaFunc, SellOfferProbability} from "~/shared/types";
import {json, useLoaderData, useParams} from "react-router";
import {CogDetailChart} from "~/components/CogDetailChart/CogDetailChart";
import {IconChartDotsFilled, IconInfoCircle, IconRobot} from '@tabler/icons-react';
import {CogDetailGeneralInfo} from "~/components/CogDetailGeneralInfo/CogDetailGeneralInfo";
import {DetailHeader} from "~/components/Header/DetailHeader";

type DetailApiResponse = {
    ok: boolean,
    data: DetailPageData
}

type TradeEngineApiResponse = {
    ok: boolean,
    data: SellOfferProbability
}

type DetailPageResponse = {
    data: {
        detail: DetailPageData,
        tradeEngine: SellOfferProbability
    }
}

export const meta = ({params}: RemixMetaFunc) => {
    return [{ title: `${beautifyCamelCase(snakeCaseToCamelCase(params.item))}` }];
};

export async function loader({ params }: LoaderFunctionArgs): Promise<Response> {
    const detailNativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/detail?item=${snakeCaseToCamelCase(params.item!)}`
    );

    const tradeEngineNativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/trade-engine`,
        {
            method: "POST",
            body: JSON.stringify({
                from: "2023-11-07",
                to: "2024-07-01",
                item: snakeCaseToCamelCase(params.item!)
            })
        }
    );

    const responses: [DetailApiResponse, TradeEngineApiResponse] = await Promise.all([
        fetch(detailNativeRequest).then((response: Response) => response.json()),
        fetch(tradeEngineNativeRequest).then((response: Response) => response.json()),
    ]);


    return json({
        ok: true,
        data: {
            detail: responses[0].data,
            tradeEngine: responses[1].data
        }
    });
}


export default function CogDetail() {
    const serverProps: DetailPageResponse = useLoaderData() as DetailPageResponse;
    const params = useParams() as {item: string};
    const lastDataPoint: Cog = serverProps.data.detail.cogs[serverProps.data.detail.cogs.length - 1];

    return (
        <Container fluid>
            <Grid gutter="xl">
                <Grid.Col span={12}>
                    <DetailHeader
                        item={beautifyCamelCase(snakeCaseToCamelCase(params.item))}
                        wikiLink={serverProps.data.detail.wiki}
                        lastDataPoint={lastDataPoint}
                    />
                </Grid.Col>
                <Grid.Col span={12}>
                    <Accordion variant="contained" defaultValue="general-info">
                        <Accordion.Item value="general-info">
                            <Accordion.Control icon={
                                <IconInfoCircle
                                    style={{
                                        width: rem(20),
                                        height: rem(20)
                                    }}
                                />
                            }>
                                General Info
                            </Accordion.Control>
                            <Accordion.Panel>
                                <CogDetailGeneralInfo
                                    item={beautifyCamelCase(snakeCaseToCamelCase(params.item))}
                                    dataPoints={serverProps.data.detail.cogs.length}
                                    creatures={serverProps.data.detail.creatures}
                                    data={serverProps.data.detail.sellOfferProbability}
                                    historic={serverProps.data.detail.sellOfferHistoricData}
                                />
                            </Accordion.Panel>
                        </Accordion.Item>
                        <Accordion.Item value="chart">
                            <Accordion.Control icon={
                                <IconChartDotsFilled
                                    style={{
                                        width: rem(20),
                                        height: rem(20)
                                    }}
                                />
                            }>
                                Chart
                            </Accordion.Control>
                            <Accordion.Panel>
                                <CogDetailChart data={serverProps.data.detail}/>
                            </Accordion.Panel>
                        </Accordion.Item>
                        <Accordion.Item value="ai-trading-bot">
                            <Accordion.Control icon={
                                <IconRobot
                                    style={{
                                        width: rem(20),
                                        height: rem(20)
                                    }}
                                />
                            }>
                                AI Trading Bot
                            </Accordion.Control>
                            <Accordion.Panel>Not implemented yet</Accordion.Panel>
                        </Accordion.Item>
                    </Accordion>
                </Grid.Col>
            </Grid>
        </Container>
    );
}