import type { MetaFunction } from "@remix-run/node";
import { LoaderFunctionArgs } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import {Grid, Container} from "@mantine/core";
import {CogChart, HomePageData} from "~/shared/types";
import {Header} from "~/components/Header/Header";
import {Outlet} from "@remix-run/react";

type HomeResponse = {
    ok: boolean;
    data: HomePageData
}

export const meta: MetaFunction = () => {
    return [
        { title: "Tibia Market" },
        { name: "description", content: "Welcome to Tibia mkt!" },
    ];
};

export async function loader(_: LoaderFunctionArgs): Promise<Response> {
    const nativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/home?item=tibiaCoin&item=honeycomb`
    );
    const nativeResponse: Response = await fetch(nativeRequest);

    const response: HomeResponse = await nativeResponse.json() as HomeResponse;

    const cogOrderMap: Map<number, string> = Object.keys(response.data)
        .reduce((result: Map<number, string>, cogName: string) => {
            const cog: CogChart = response.data[cogName];

            return result.set(cog.pagePosition, cogName);
        }, new Map<number, string>());

    let result: HomePageData = {};

    for (let i: number = 0; i < Object.keys(response.data).length; i++) {
        const cogName: string = cogOrderMap.get(i + 1)!;

        result = {...result, [cogName]: response.data[cogName]};
    }

    return json({
        ok: response.ok,
        data: result
    });
}

export default function Index() {
    const serverProps: HomeResponse = useLoaderData<typeof loader>() as HomeResponse;

    return (
        <Container fluid>
            <Grid gutter="xl">
                <Grid.Col span={12}>
                    <Header/>
                </Grid.Col>
                {Object.keys(serverProps.data).map((cogName: string) => {
                    const cog: CogChart = serverProps.data[cogName];

                    return (
                        <Grid.Col key={cogName} span={12}>
                            <CogPreview name={cogName} wikiLink={cog.wiki} data={cog}/>
                        </Grid.Col>
                    );
                })}
            </Grid>
        </Container>
    );
}
