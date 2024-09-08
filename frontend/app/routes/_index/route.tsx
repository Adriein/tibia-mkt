import type { MetaFunction } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import {Grid, Container, Space, Center, Box} from "@mantine/core";
import {CogChart, HomePageData} from "~/shared/types";
import {Header} from "~/components/Header/Header";
import {TIBIA_COIN} from "~/shared/constants";
import {GoodPreviewChip} from "~/components/GoodPreviewChip/GoodPreviewChip";
import classes from "./HomeRoute.module.css";
import {HomeResponse} from "~/routes/_index/types";
import {fetchData, orderHomePageData} from "~/routes/_index/routeFunctions";

type HomeServerProps = {
    home: HomeResponse<HomePageData>,
    search: HomeResponse<string[]>
}

export const meta: MetaFunction = () => {
    return [
        { title: "Tibia Market" },
        { name: "description", content: "Welcome to Tibia mkt!" },
    ];
};

export async function loader(): Promise<Response> {
    const [ homeData, searchInputData ]: [HomeResponse<HomePageData>, HomeResponse<string[]>] = await fetchData();

    const orderedHomeData: HomePageData = orderHomePageData(homeData);

    return json({
        home: {
            ok: homeData.ok,
            data: orderedHomeData
        },
        search: {
            ok: searchInputData.ok,
            data: searchInputData.data
        }
    });
}

export default function Index() {
    const serverProps: HomeServerProps = useLoaderData() as HomeServerProps;
    const tibiaCoin: CogChart = serverProps.home.data[TIBIA_COIN];

    return (
        <Container fluid className={classes.fullHeight}>
            <Header search={serverProps.search.data}/>
            <Space h="xl"/>
            {tibiaCoin?
                <Grid gutter="xl">
                    <Grid.Col span={12}>
                        <GoodPreviewChip data={serverProps.home.data}/>
                    </Grid.Col>
                    <Grid.Col key={TIBIA_COIN} span={12}>
                        <CogPreview name={TIBIA_COIN} wikiLink={tibiaCoin.wiki} data={tibiaCoin}/>
                    </Grid.Col>
                </Grid> :
                <Grid gutter="xl">
                    <Grid.Col span={12}>
                        No data
                    </Grid.Col>
                </Grid>
            }
        </Container>
    );
}
