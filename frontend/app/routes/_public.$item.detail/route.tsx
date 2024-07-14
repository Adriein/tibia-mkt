import {beautifyCamelCase, snakeCaseToCamelCase} from "~/shared/util";
import {Accordion, Container, Grid} from "@mantine/core";
import {Header} from "~/components/Header/Header";

export const meta = ({params}) => {
    return [{ title: `${beautifyCamelCase(snakeCaseToCamelCase(params.item))} detail` }];
};

export default function CogDetail() {
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
                            <Accordion.Panel>Content</Accordion.Panel>
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