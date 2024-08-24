import classes from "./Header.module.css";
import {Title} from "@mantine/core";
import {SelectGood} from "~/components/SelectGood/SelectGood";

interface HeaderProps {
    search: string[]
}

export function Header({ search }: HeaderProps) {
    return (
        <header className={classes.header}>
            <Title className={classes.title}>Tibia Mkt</Title>
            <SelectGood search={search}/>
        </header>
    )
}