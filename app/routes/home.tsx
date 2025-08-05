import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Guess Card" },
    { name: "description", content: "" },
  ];
}

export default function Home() {
  //
}
