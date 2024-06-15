import type { MetaFunction } from "@remix-run/node";
import { ColorSchemeToggle } from "~/components/ColorSchemeToggle/ColorSchemeToggle";
import { LoaderFunctionArgs } from "@remix-run/node";
import { json, useLoaderData } from "react-router";
import { CogPreview } from "~/components/CogPreview/CogPreview";
import { Grid, Container } from "@mantine/core";
import { HomePageData } from "~/shared/types";

type HomeResponse = {
    ok: boolean;
    data: HomePageData
}

const API_URL: string = `${process.env.API_PROTOCOL}://${process.env.API_URL}/home`;

export const meta: MetaFunction = () => {
  return [
    { title: "Tibia Market" },
    { name: "description", content: "Welcome to Tibia mkt!" },
  ];
};

export async function loader({ request, params }: LoaderFunctionArgs): Promise<Response> {
    const nativeRequest: Request = new Request(API_URL);
    const nativeResponse: Response = await fetch(nativeRequest);

    const response: HomeResponse = await nativeResponse.json() as HomeResponse;

    return json({
        ok: response.ok,
        data: response.data
    });
}

export default function Index() {
  const serverProps = useLoaderData<typeof loader>();
    console.log(serverProps)
  return (
      <Container fluid>
          <Grid>
              <Grid.Col span={12}>
                  <h2>Tibia Mkt</h2>
                  <ColorSchemeToggle />
              </Grid.Col>
              <Grid.Col span={12}>
                  <CogPreview data={serverProps.data}/>
              </Grid.Col>
          </Grid>
      </Container>
  );
}
