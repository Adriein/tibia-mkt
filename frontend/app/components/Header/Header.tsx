import classes from "./Header.module.css";
import {Title} from "@mantine/core";
import {ColorSchemeToggle} from "~/components/ColorSchemeToggle/ColorSchemeToggle";

interface HeaderProps {
    item?: string;
}

export function Header({ item }: HeaderProps) {
    return (
        <header className={classes.header}>
            <Title className={classes.title}>{item ?? "Tibia Mkt"}</Title>
            <ColorSchemeToggle/>
        </header>
    )
}