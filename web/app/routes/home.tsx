import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Tibia Mkt" },
    { name: "description", content: "Welcome to Tibia Mkt!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
