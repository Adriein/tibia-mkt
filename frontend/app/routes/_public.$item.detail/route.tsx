import {beautifyCamelCase, snakeCaseToCamelCase} from "~/shared/util";
import {Accordion, Container, Grid, rem} from "@mantine/core";
import {Header} from "~/components/Header/Header";
import {LoaderFunctionArgs} from "@remix-run/node";
import {DetailPageData, RemixMetaFunc} from "~/shared/types";
import {json, useLoaderData} from "react-router";
import {CogDetailChart} from "~/components/CogDetailChart/CogDetailChart";
import {IconChartDotsFilled, IconInfoCircle, IconRobot} from '@tabler/icons-react';

type DetailResponse = {
    ok: boolean;
    data: DetailPageData
}

export const meta = ({params}: RemixMetaFunc) => {
    return [{ title: `${beautifyCamelCase(snakeCaseToCamelCase(params.item))}` }];
};

export async function loader({ params }: LoaderFunctionArgs): Promise<Response> {
    const nativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/detail?item=${snakeCaseToCamelCase(params.item!)}`
    );

    const nativeResponse: Response = await fetch(nativeRequest);

    const response: DetailResponse = await nativeResponse.json() as DetailResponse;

    return json({
        ok: response.ok,
        data: response.data
    });
}


export default function CogDetail() {
    const serverProps: DetailResponse = useLoaderData() as DetailResponse;
    return (
        <Container fluid>
            <Grid gutter="xl">
                <Grid.Col span={12}>
                    <Header/>
                </Grid.Col>
                <Grid.Col span={12}>
                    <Accordion variant="contained" defaultValue="chart">
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
                                <CogDetailChart data={serverProps.data}/>
                            </Accordion.Panel>
                        </Accordion.Item>
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
                            <Accordion.Panel>Content</Accordion.Panel>
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