import type { MetaFunction } from "@remix-run/node";
import { ColorSchemeToggle } from "~/components/ColorSchemeToggle/ColorSchemeToggle";
import {LoaderFunctionArgs} from "@remix-run/node";
import {json, useLoaderData} from "react-router";
import {CogPreview} from "~/components/CogPreview/CogPreview";

type TibiaCoinCog = {buyPrice: number, sellPrice: number, date: string, world: string}

type HomeResponse = {
    ok: boolean;
    data: TibiaCoinCog[]
}

export const meta: MetaFunction = () => {
  return [
    { title: "Tibia Market" },
    { name: "description", content: "Welcome to Tibia mkt!" },
  ];
};

export async function loader({ request, params }: LoaderFunctionArgs) {
    const req: Request = new Request(`http://${process.env.API_URL}/home`);
    const response: Response = await fetch(req);

    const res: HomeResponse = await response.json() as HomeResponse;

    return json({
        ok: res.ok,
        data: res.data
    })
}

export default function Index() {
  const serverProps = useLoaderData<typeof loader>();

  return (
    <div>
      <CogPreview data={serverProps.data}/>
      <ColorSchemeToggle />
    </div>
  );
}
