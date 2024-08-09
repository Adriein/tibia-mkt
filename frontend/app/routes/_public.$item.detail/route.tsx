import {beautifyCamelCase, snakeCaseToCamelCase} from "~/shared/util";
import {Accordion, Container, Grid, rem} from "@mantine/core";
import {Header} from "~/components/Header/Header";
import {LoaderFunctionArgs} from "@remix-run/node";
import {DetailPageData, RemixMetaFunc, TradeEngineDetailPageData} from "~/shared/types";
import {json, useLoaderData, useParams} from "react-router";
import {CogDetailChart} from "~/components/CogDetailChart/CogDetailChart";
import {IconChartDotsFilled, IconInfoCircle, IconRobot} from '@tabler/icons-react';
import {CogDetailGeneralInfo} from "~/components/CogDetailGeneralInfo/CogDetailGeneralInfo";

type DetailApiResponse = {
    ok: boolean,
    data: DetailPageData
}

type TradeEngineApiResponse = {
    ok: boolean,
    data: TradeEngineDetailPageData
}

type DetailPageResponse = {
    data: {
        detail: DetailPageData,
        tradeEngine: TradeEngineDetailPageData
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

    return (
        <Container fluid>
            <Grid gutter="xl">
                <Grid.Col span={12}>
                    <Header/>
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
                                    dataPoints={serverProps.data.detail.cog.length}
                                    creatures={serverProps.data.detail.creatures}
                                    data={serverProps.data.tradeEngine}
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
                            <Accordion.Panel>Content</Accordion.Panel>
                        </Accordion.Item>
                    </Accordion>
                </Grid.Col>
            </Grid>
        </Container>
    );
}