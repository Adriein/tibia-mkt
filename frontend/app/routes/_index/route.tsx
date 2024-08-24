import type { MetaFunction } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import {Grid, Container} from "@mantine/core";
import {CogChart, HomePageData} from "~/shared/types";
import {Header} from "~/components/Header/Header";
import {TIBIA_COIN} from "~/shared/constants";
import {GoodPreviewChip} from "~/components/GoodPreviewChip/GoodPreviewChip";
import classes from "./HomeRoute.module.css";

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

export async function loader(): Promise<Response> {
    const nativeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/home?item=tibiaCoin&item=honeycomb&item=swamplingWood&item=brokenShamanicStaff`
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
    const serverProps: HomeResponse = useLoaderData() as HomeResponse;
    const tibiaCoin: CogChart = serverProps.data[TIBIA_COIN];

    return (
        <Container fluid>
            <Header/>
            <Grid gutter="xl" className={classes.fullHeight}>
                <Grid.Col span={12}>
                    <GoodPreviewChip data={serverProps.data}/>
                </Grid.Col>
                <Grid.Col key={TIBIA_COIN} span={12}>
                    <CogPreview name={TIBIA_COIN} wikiLink={tibiaCoin.wiki} data={tibiaCoin}/>
                </Grid.Col>
            </Grid>
        </Container>
    );
}
