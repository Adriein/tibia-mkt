import type { MetaFunction } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import {Grid, Container} from "@mantine/core";
import {CogChart, HomePageData} from "~/shared/types";
import {Header} from "~/components/Header/Header";
import {TIBIA_COIN} from "~/shared/constants";
import {GoodPreviewChip} from "~/components/GoodPreviewChip/GoodPreviewChip";
import classes from "./HomeRoute.module.css";

type HomeServerProps = {
    home: HomeResponse<HomePageData>,
    search: HomeResponse<string[]>
}

type HomeResponse<T> = {
    ok: boolean;
    data: T
}

export const meta: MetaFunction = () => {
    return [
        { title: "Tibia Market" },
        { name: "description", content: "Welcome to Tibia mkt!" },
    ];
};

export async function loader(): Promise<Response> {
    const homeRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}` +
        "/home?" +
        "item=tibiaCoin&item=honeycomb&item=swamplingWood&item=brokenShamanicStaff"
    );

    const searchGoodRequest: Request = new Request(
        `${process.env.API_PROTOCOL}://${process.env.API_URL}/goods`
    );

    const responses: [HomeResponse<HomePageData>, HomeResponse<string[]>] = await Promise.all([
        fetch(homeRequest).then((response: Response) => response.json()),
        fetch(searchGoodRequest).then((response: Response) => response.json()),
    ]);

    const homeResponse: HomeResponse<HomePageData> = responses[0];
    const searchResponse: HomeResponse<string[]> = responses[1];

    const cogOrderMap: Map<number, string> = Object.keys(homeResponse.data)
        .reduce((result: Map<number, string>, cogName: string) => {
            const cog: CogChart = homeResponse.data[cogName];

            return result.set(cog.pagePosition, cogName);
        }, new Map<number, string>());

    let result: HomePageData = {};

    for (let i: number = 0; i < Object.keys(homeResponse.data).length; i++) {
        const cogName: string = cogOrderMap.get(i + 1)!;

        result = {...result, [cogName]: homeResponse.data[cogName]};
    }

    return json({
        home: {
            ok: homeResponse.ok,
            data: result
        },
        search: {
            ok: searchResponse.ok,
            data: searchResponse.data
        }
    });
}

export default function Index() {
    const serverProps: HomeServerProps = useLoaderData() as HomeServerProps;
    const tibiaCoin: CogChart = serverProps.home.data[TIBIA_COIN];

    return (
        <Container fluid>
            <Header search={serverProps.search.data}/>
            <Grid gutter="xl" className={classes.fullHeight}>
                <Grid.Col span={12}>
                    <GoodPreviewChip data={serverProps.home.data}/>
                </Grid.Col>
                <Grid.Col key={TIBIA_COIN} span={12}>
                    <CogPreview name={TIBIA_COIN} wikiLink={tibiaCoin.wiki} data={tibiaCoin}/>
                </Grid.Col>
            </Grid>
        </Container>
    );
}
