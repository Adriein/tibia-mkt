import type { MetaFunction } from "@remix-run/node";
import { LoaderFunctionArgs } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import {Grid, Container} from "@mantine/core";
import { HomePageData } from "~/shared/types";
import {Header} from "~/components/Header/Header";

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

    return json({
        ok: response.ok,
        data: response.data
    });
}

export default function Index() {
  const serverProps = useLoaderData<typeof loader>();

  return (
      <Container fluid>
          <Grid gutter="xl">
              <Grid.Col span={12}>
                  <Header/>
              </Grid.Col>
              <Grid.Col span={12}>
                  <CogPreview data={serverProps.data}/>
              </Grid.Col>
          </Grid>
      </Container>
  );
}
