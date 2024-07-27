import {beautifyCamelCase, formatDate, snakeCaseToCamelCase} from "~/shared/util";
import {Accordion, Container, Grid} from "@mantine/core";
import {Header} from "~/components/Header/Header";
import {LoaderFunctionArgs} from "@remix-run/node";
import {DetailPageData, YAxisTick} from "~/shared/types";
import {json, useLoaderData} from "react-router";
import {AreaChart} from "@mantine/charts";
import {xAxisDateFormatter, xAxisTick, yAxisNumberFormatter} from "~/shared/chart-util";

type DetailResponse = {
    ok: boolean;
    data: DetailPageData
}

export const meta = ({params}) => {
    return [{ title: `${beautifyCamelCase(snakeCaseToCamelCase(params.item))}` }];
};

export async function loader({ params }: LoaderFunctionArgs): Promise<Response> {
    const nativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/detail?item=${snakeCaseToCamelCase(params.item)}`
    );

    const nativeResponse: Response = await fetch(nativeRequest);

    const response: DetailResponse = await nativeResponse.json() as DetailResponse;

    return json({
        ok: response.ok,
        data: response.data
    });
}


export default function CogDetail() {
    const serverProps: DetailResponse = useLoaderData<typeof loader>() as DetailResponse;
    return (
        <Container fluid>
            <Grid gutter="xl">
                <Grid.Col span={12}>
                    <Header/>
                </Grid.Col>
                <Grid.Col span={12}>
                    <Accordion variant="contained">
                        <Accordion.Item value="chart">
                            <Accordion.Control>
                                Chart
                            </Accordion.Control>
                            <Accordion.Panel>
                                <AreaChart
                                    h={180}
                                    data={serverProps.data.cog}
                                    dataKey="date"
                                    series={[{name: 'sellOffer', label: "Sell Offer", color: 'teal.6'}]}
                                    areaChartProps={{ syncId: 'groceries' }}
                                    xAxisProps={{
                                        interval: "preserveStartEnd",
                                        tickFormatter: xAxisDateFormatter,
                                        ticks: xAxisTick(serverProps.data.cog, serverProps.data.sellOfferChart.xAxisTick)
                                    }}
                                    yAxisProps={{
                                        domain: serverProps.data.sellOfferChart.yAxisTick.map((tick: YAxisTick) => tick.price)
                                    }}
                                    valueFormatter={yAxisNumberFormatter}
                                />
                                <AreaChart
                                    h={180}
                                    data={serverProps.data.cog}
                                    dataKey="date"
                                    areaChartProps={{ syncId: 'groceries' }}
                                    series={[{name: 'buyOffer', label: "Buy Offer", color: 'indigo.6'}]}
                                    xAxisProps={{
                                        interval: "preserveStartEnd",
                                        tickFormatter: xAxisDateFormatter,
                                        ticks: xAxisTick(serverProps.data.cog, serverProps.data.buyOfferChart.xAxisTick)
                                    }}
                                    yAxisProps={{
                                        domain: serverProps.data.buyOfferChart.yAxisTick.map((tick: YAxisTick) => tick.price)
                                    }}
                                    valueFormatter={yAxisNumberFormatter}
                                />
                            </Accordion.Panel>
                        </Accordion.Item>
                        <Accordion.Item value="general-info">
                            <Accordion.Control>
                                General Info
                            </Accordion.Control>
                            <Accordion.Panel>Content</Accordion.Panel>
                        </Accordion.Item>
                        <Accordion.Item value="ai-trading-bot">
                            <Accordion.Control>
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