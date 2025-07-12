import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "TibiaMkt" },
    { name: "description", content: "Welcome to Tibia Mkt!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
