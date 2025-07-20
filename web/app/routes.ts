import {type RouteConfig, index, route} from "@react-router/dev/routes";

export default [
    index("routes/home/home.tsx"),
    route("/:good/detail", "routes/detail/detail.tsx"),
] satisfies RouteConfig;
