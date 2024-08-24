import classes from "./Header.module.css";
import {Title} from "@mantine/core";
import {SelectGood} from "~/components/SelectGood/SelectGood";

export function Header() {
    return (
        <header className={classes.header}>
            <Title className={classes.title}>Tibia Mkt</Title>
            <SelectGood/>
        </header>
    )
}